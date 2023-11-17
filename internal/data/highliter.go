package data

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma"

	cHtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown/ast"
)

var (
	htmlFormatter  *cHtml.Formatter
	highlightStyle *chroma.Style
)

func initialize() {
	htmlFormatter = cHtml.New(cHtml.WithClasses(true), cHtml.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}
	styleName := "monokai"
	highlightStyle = styles.Get(styleName)
	if highlightStyle == nil {
		panic(fmt.Sprintf("didn't find style '%s'", styleName))
	}
}

func htmlHighlight(w io.Writer, source, lang, defaultLang string) error {
	if htmlFormatter == nil {
		initialize()
	}

	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	err = htmlFormatter.Format(w, highlightStyle, it)
	if err != nil {
		return err
	}

	w.Write([]byte("<style>"))
	err = htmlFormatter.WriteCSS(w, highlightStyle)
	if err != nil {
		return err
	}
	w.Write([]byte("<style>"))

	return nil
}

func renderCode(w io.Writer, codeBlock *ast.CodeBlock, entering bool) {
	defaultLang := ""
	lang := string(codeBlock.Info)
	err := htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
	if err != nil {
		fmt.Println("error while highlighting:", err)
	}
}

func highlightingRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(w, code, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

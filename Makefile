run:
	VERSION_HASH=0 npm run build:css
	go run .

dev:
	npm run watch:css & air
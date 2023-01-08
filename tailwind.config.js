/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./templates/**/*.html', './data/blogs/*.md', './data/projects/*.md'],
  theme: {
    extend: {
      colors: {
        'background-1': '#1a202c'
      },
      fontFamily: {
        'roboto': ['Roboto', 'sans-serif']
      },
      container: {
        center: true
      },
      spacing: {
        '1440': '1440px'
      }
    }
  },
  plugins: [],
}

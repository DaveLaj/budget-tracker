/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/templates/**/*.templ",
    "./internal/templates/**/*.go",
    "./input.css",
    "./assets/styles.css",
    "./assets/index.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}


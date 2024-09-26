/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './web/static/**/*.js',
    './web/templates/**/*.html',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}


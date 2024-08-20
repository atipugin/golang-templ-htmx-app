/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./../internal/view/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};

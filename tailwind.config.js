/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.html",
    "./node_modules/flowbite/**/*.js"
  ],
  plugins: [require("flowbite/plugin")],
  theme: {
    extend: {
      blur: {
        "4xl": "128px",
        "5xl": "192px",
      },
      colors: {
        "mc-yellow": "rgb(255, 214, 68)",
        "mc-purple": "#3A3185",
        "mc-light-purple": "#938CD1",
        "mc-light-green": "#B3D138",
        "mc-light-black": "rgba(0, 0, 0, 0.7)",
        "mc-dark-purple": "#231e52",
      },
      container: {
        center: true,
      },
      fontFamily: {
        pangram: ["Pangram"]
      },
      borderWidth: {
        "3": "0.2rem",
      },
      borderRadius: {
        "5xl": "2.5rem",
      },
      screens: {
        xs: "480px",
      },
    },
  },
  plugins: [],
}


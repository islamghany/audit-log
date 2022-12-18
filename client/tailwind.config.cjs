/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: "#6366F1",
        overlay: "#1f2987",
        secondary: "#D1D5DB",
        white: "#F3F4F6",
        text: "#9AC3AF",
        background: "#0F172A",
      },
    },
  },
  plugins: [],
};

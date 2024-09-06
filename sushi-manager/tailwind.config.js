/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        cabin: "Cabin Variable, sans-serif",
        lora: "Lora Variable, sans-serif",
        bitter: "Bitter Variable, sans-serif",
      },
    },
  },
  plugins: [],
};

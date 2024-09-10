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
      colors: {
        httpGet: "#62affc",
        httpPut: "#faa232",
        httpPost: "#4caf50",
        httpDelete: "#f44336",
        httpPatch: "#ff5722",
        httpOptions: "#9c27b0",
      },
    },
  },
  plugins: [],
};

const express = require("express");
const app = express();
const port = 3000;

// Mock data
const sushiData = [
  {
    id: 1,
    name: "California Roll",
    ingredients: ["Crab", "Avocado", "Cucumber"],
  },
  { id: 2, name: "Tuna Roll", ingredients: ["Tuna", "Rice", "Nori"] },
];

const restaurantData = [
  { id: 1, name: "Sushi Place", location: "123 Sushi St." },
  { id: 2, name: "Roll House", location: "456 Roll Ave." },
];

// GET /v1/sushi
app.get("/v1/sushi", (req, res) => {
  res.json(sushiData);
});

// GET /v1/sushi/restaurant
app.get("/v1/sushi/restaurant", (req, res) => {
  res.json(restaurantData);
});

app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});

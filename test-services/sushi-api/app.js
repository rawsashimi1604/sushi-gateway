const express = require("express");
const jwt = require("jsonwebtoken");

require("dotenv").config();

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
  res.json({ app_id: process.env.APP_ID, data: sushiData });
});

// GET /v1/sushi/:id
app.get("/v1/sushi/:id", (req, res) => {
  const sushiId = parseInt(req.params.id);
  const sushi = sushiData.find((sushi) => sushi.id === sushiId);
  if (!sushi) {
    res.status(404).json({ error: "Sushi not found" });
  } else {
    res.json({ app_id: process.env.APP_ID, data: sushi });
  }
});

// GET /v1/sushi/restaurant
app.get("/v1/sushi/restaurant", (req, res) => {
  res.json({ app_id: process.env.APP_ID, data: restaurantData });
});

// GET /v1/sushi/restaurant/:id
app.get("/v1/sushi/restaurant/:id", (req, res) => {
  const restaurantId = parseInt(req.params.id);
  const restaurant = restaurantData.find(
    (restaurant) => restaurant.id === restaurantId
  );
  if (!restaurant) {
    res.status(404).json({ error: "Restaurant not found" });
  } else {
    res.json({ app_id: process.env.APP_ID, data: restaurant });
  }
});

app.get("/v1/token", (req, res) => {
  const payload = {
    app_id: process.env.APP_ID,
    iss: process.env.JWT_ISSUER,
  };
  const secretKey = process.env.JWT_SECRET;
  const token = jwt.sign(payload, secretKey, {
    expiresIn: "1h",
  });
  res.json({ token });
});

app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});

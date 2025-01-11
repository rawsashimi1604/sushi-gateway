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

// GET /health
app.get("/health", (req, res) => {
  res.status(200).json({ app_id: process.env.APP_ID, data: "ok" });
});

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
  const signingMethod = req.query.alg
  const availableSigningMethods = ["HS256", "RS256"]
  
  if (!availableSigningMethods.includes(signingMethod)) {
    return res.status(400).json({ error: "Invalid signing method. Use 'HS256' or 'RS256'" });
  }

  const payload = {
    app_id: process.env.APP_ID,
    iss: process.env.JWT_ISSUER,
  };

  let token;
  try {
    if (signingMethod === 'RS256') {
      // For RS256, we need the private key
      const privateKey = process.env.JWT_PRIVATE_KEY;
      if (!privateKey) {
        return res.status(500).json({ error: "RS256 private key not configured" });
      }
      token = jwt.sign(payload, privateKey, {
        algorithm: 'RS256',
        expiresIn: "1h"
      });
    } else if (signingMethod === 'HS256') {
      // For HS256, we use the secret key
      const secretKey = process.env.JWT_SECRET;
      if (!secretKey) {
        return res.status(500).json({ error: "HS256 secret key not configured" });
      }
      token = jwt.sign(payload, secretKey, {
        algorithm: 'HS256',
        expiresIn: "1h"
      });
    } 
    res.json({ token });
  } catch (err) {
    console.error('Error generating token:', err);  
    res.status(500).json({ error: "Failed to generate token" });
  }
});

app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});

const express = require("express");

require("dotenv").config();

const app = express();
const port = 3000;

// Receive http logs
app.post("/v1/log", (req, res) => {
  console.log(req.body);
  res.json({ log_received: req.body });
});

app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});

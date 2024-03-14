const express = require("express");
const bodyParser = require("body-parser");

require("dotenv").config();

const app = express();
const port = 3000;

// create application/json parser
var jsonParser = bodyParser.json();

// Receive http logs
app.post("/v1/log", jsonParser, (req, res) => {
  console.log(req.body);
  res.json({ log_received: req.body });
});

app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});

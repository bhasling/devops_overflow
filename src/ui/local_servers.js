var express = require('express');
var app = express();
app.use("/", express.static('./out'));
app.listen(3000);

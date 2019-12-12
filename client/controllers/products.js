var Client = require("../gprc/client");
var messages = require("../pb/server_pb")

exports.getProduct = (req, res, next) => {
  res.status(200).json({
    products: [{}]
  });
};

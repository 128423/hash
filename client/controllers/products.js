const services = require("../pb/server_grpc_pb");
const messages = require("../pb/server_pb");

const grpc = require("grpc");
var db = require("../database/connection");
var Product = require("../models/product");

var client = new services.HashGrpcAPiClient(
  "localhost:50051",
  grpc.credentials.createInsecure()
);

exports.getProduct = async (req, res, next) => {
  var prod = db.model("products", Product.schema, "products");
  var result = [];
  await prod
    .find({})
    .lean()
    .exec()
    .then(docs => (result = docs))
    .catch(err => res.status(500).json({ error: err }));

  result.map(async p => {
    var msg = new messages.RequestDiscount();
    msg.setProductId(p._id);
    msg.setUserId(1);
    await client.getDiscount(msg, (err, response) => {
      if (err == null && response) {
        p.discount = response.toObject();
      }
    });
  });
  res.json({ data: result });
};

// function(e, docs) {
//   docs.map(p => {
//     var msg = new messages.RequestDiscount();
//     msg.setProductId(p._id);
//     msg.setUserId(1);
//     client.getDiscount(msg, (err, response) => {
//       var resp = undefined;
//       if (err == null && response) {
//         resp = response.toObject();
//       }
//       result.push({ ...p, discount: resp });
//     });
//   });
// }

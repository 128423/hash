const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const ProductSchema = new Schema({
  id: {
    type: Number,
    require: [true]
  },
  price_in_cents: {
    type: Number
  },
  title: {
    type: String
  },
  description: {
    type: String
  }
},{ collection: "products" });

const product = mongoose.model("product", ProductSchema);

module.exports = product;


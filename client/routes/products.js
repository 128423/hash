const express = require("express");

const productController = require("../controllers/products");

const router = express.Router();

// GET /product
router.get("/products",productController.getProduct);

module.exports = router;

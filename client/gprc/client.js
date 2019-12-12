var services = require("../pb/server_grpc_pb");

var grpc = require("grpc")

var Client = new services.HashGrpcAPiClient('localhost:50051',
grpc.credentials.createInsecure());

module.exports = Client
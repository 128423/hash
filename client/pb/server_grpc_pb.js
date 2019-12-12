// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var pb_server_pb = require('../pb/server_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_products_RequestDiscount(arg) {
  if (!(arg instanceof pb_server_pb.RequestDiscount)) {
    throw new Error('Expected argument of type products.RequestDiscount');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_products_RequestDiscount(buffer_arg) {
  return pb_server_pb.RequestDiscount.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_products_ResponseDiscount(arg) {
  if (!(arg instanceof pb_server_pb.ResponseDiscount)) {
    throw new Error('Expected argument of type products.ResponseDiscount');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_products_ResponseDiscount(buffer_arg) {
  return pb_server_pb.ResponseDiscount.deserializeBinary(new Uint8Array(buffer_arg));
}


var HashGrpcAPiService = exports.HashGrpcAPiService = {
  getDiscount: {
    path: '/products.HashGrpcAPi/getDiscount',
    requestStream: false,
    responseStream: false,
    requestType: pb_server_pb.RequestDiscount,
    responseType: pb_server_pb.ResponseDiscount,
    requestSerialize: serialize_products_RequestDiscount,
    requestDeserialize: deserialize_products_RequestDiscount,
    responseSerialize: serialize_products_ResponseDiscount,
    responseDeserialize: deserialize_products_ResponseDiscount,
  },
};

exports.HashGrpcAPiClient = grpc.makeGenericClientConstructor(HashGrpcAPiService);

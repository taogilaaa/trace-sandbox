// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var sandbox_sales_v1_pb = require('./sandbox_sales_v1_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_sandbox_sales_v1_CreateSaleOrderRequest(arg) {
  if (!(arg instanceof sandbox_sales_v1_pb.CreateSaleOrderRequest)) {
    throw new Error('Expected argument of type sandbox.sales.v1.CreateSaleOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sandbox_sales_v1_CreateSaleOrderRequest(buffer_arg) {
  return sandbox_sales_v1_pb.CreateSaleOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sandbox_sales_v1_CreateSaleOrderResponse(arg) {
  if (!(arg instanceof sandbox_sales_v1_pb.CreateSaleOrderResponse)) {
    throw new Error('Expected argument of type sandbox.sales.v1.CreateSaleOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sandbox_sales_v1_CreateSaleOrderResponse(buffer_arg) {
  return sandbox_sales_v1_pb.CreateSaleOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sandbox_sales_v1_GetSaleOrderRequest(arg) {
  if (!(arg instanceof sandbox_sales_v1_pb.GetSaleOrderRequest)) {
    throw new Error('Expected argument of type sandbox.sales.v1.GetSaleOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sandbox_sales_v1_GetSaleOrderRequest(buffer_arg) {
  return sandbox_sales_v1_pb.GetSaleOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sandbox_sales_v1_GetSaleOrderResponse(arg) {
  if (!(arg instanceof sandbox_sales_v1_pb.GetSaleOrderResponse)) {
    throw new Error('Expected argument of type sandbox.sales.v1.GetSaleOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sandbox_sales_v1_GetSaleOrderResponse(buffer_arg) {
  return sandbox_sales_v1_pb.GetSaleOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sandbox_sales_v1_GetSaleOrdersRequest(arg) {
  if (!(arg instanceof sandbox_sales_v1_pb.GetSaleOrdersRequest)) {
    throw new Error('Expected argument of type sandbox.sales.v1.GetSaleOrdersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sandbox_sales_v1_GetSaleOrdersRequest(buffer_arg) {
  return sandbox_sales_v1_pb.GetSaleOrdersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sandbox_sales_v1_GetSaleOrdersResponse(arg) {
  if (!(arg instanceof sandbox_sales_v1_pb.GetSaleOrdersResponse)) {
    throw new Error('Expected argument of type sandbox.sales.v1.GetSaleOrdersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sandbox_sales_v1_GetSaleOrdersResponse(buffer_arg) {
  return sandbox_sales_v1_pb.GetSaleOrdersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var SaleOrderServiceService = exports.SaleOrderServiceService = {
  getSaleOrder: {
    path: '/sandbox.sales.v1.SaleOrderService/GetSaleOrder',
    requestStream: false,
    responseStream: false,
    requestType: sandbox_sales_v1_pb.GetSaleOrderRequest,
    responseType: sandbox_sales_v1_pb.GetSaleOrderResponse,
    requestSerialize: serialize_sandbox_sales_v1_GetSaleOrderRequest,
    requestDeserialize: deserialize_sandbox_sales_v1_GetSaleOrderRequest,
    responseSerialize: serialize_sandbox_sales_v1_GetSaleOrderResponse,
    responseDeserialize: deserialize_sandbox_sales_v1_GetSaleOrderResponse,
  },
  getSaleOrders: {
    path: '/sandbox.sales.v1.SaleOrderService/GetSaleOrders',
    requestStream: false,
    responseStream: false,
    requestType: sandbox_sales_v1_pb.GetSaleOrdersRequest,
    responseType: sandbox_sales_v1_pb.GetSaleOrdersResponse,
    requestSerialize: serialize_sandbox_sales_v1_GetSaleOrdersRequest,
    requestDeserialize: deserialize_sandbox_sales_v1_GetSaleOrdersRequest,
    responseSerialize: serialize_sandbox_sales_v1_GetSaleOrdersResponse,
    responseDeserialize: deserialize_sandbox_sales_v1_GetSaleOrdersResponse,
  },
  createSaleOrder: {
    path: '/sandbox.sales.v1.SaleOrderService/CreateSaleOrder',
    requestStream: false,
    responseStream: false,
    requestType: sandbox_sales_v1_pb.CreateSaleOrderRequest,
    responseType: sandbox_sales_v1_pb.CreateSaleOrderResponse,
    requestSerialize: serialize_sandbox_sales_v1_CreateSaleOrderRequest,
    requestDeserialize: deserialize_sandbox_sales_v1_CreateSaleOrderRequest,
    responseSerialize: serialize_sandbox_sales_v1_CreateSaleOrderResponse,
    responseDeserialize: deserialize_sandbox_sales_v1_CreateSaleOrderResponse,
  },
};

exports.SaleOrderServiceClient = grpc.makeGenericClientConstructor(SaleOrderServiceService);

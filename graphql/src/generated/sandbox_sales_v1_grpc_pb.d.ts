// GENERATED CODE -- DO NOT EDIT!

// package: sandbox.sales.v1
// file: sandbox_sales_v1.proto

import * as sandbox_sales_v1_pb from "./sandbox_sales_v1_pb";
import * as grpc from "grpc";

interface ISaleOrderServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  getSaleOrder: grpc.MethodDefinition<sandbox_sales_v1_pb.GetSaleOrderRequest, sandbox_sales_v1_pb.GetSaleOrderResponse>;
  getSaleOrders: grpc.MethodDefinition<sandbox_sales_v1_pb.GetSaleOrdersRequest, sandbox_sales_v1_pb.GetSaleOrdersResponse>;
  createSaleOrder: grpc.MethodDefinition<sandbox_sales_v1_pb.CreateSaleOrderRequest, sandbox_sales_v1_pb.CreateSaleOrderResponse>;
}

export const SaleOrderServiceService: ISaleOrderServiceService;

export interface ISaleOrderServiceServer extends grpc.UntypedServiceImplementation {
  getSaleOrder: grpc.handleUnaryCall<sandbox_sales_v1_pb.GetSaleOrderRequest, sandbox_sales_v1_pb.GetSaleOrderResponse>;
  getSaleOrders: grpc.handleUnaryCall<sandbox_sales_v1_pb.GetSaleOrdersRequest, sandbox_sales_v1_pb.GetSaleOrdersResponse>;
  createSaleOrder: grpc.handleUnaryCall<sandbox_sales_v1_pb.CreateSaleOrderRequest, sandbox_sales_v1_pb.CreateSaleOrderResponse>;
}

export class SaleOrderServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  getSaleOrder(argument: sandbox_sales_v1_pb.GetSaleOrderRequest, callback: grpc.requestCallback<sandbox_sales_v1_pb.GetSaleOrderResponse>): grpc.ClientUnaryCall;
  getSaleOrder(argument: sandbox_sales_v1_pb.GetSaleOrderRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<sandbox_sales_v1_pb.GetSaleOrderResponse>): grpc.ClientUnaryCall;
  getSaleOrder(argument: sandbox_sales_v1_pb.GetSaleOrderRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<sandbox_sales_v1_pb.GetSaleOrderResponse>): grpc.ClientUnaryCall;
  getSaleOrders(argument: sandbox_sales_v1_pb.GetSaleOrdersRequest, callback: grpc.requestCallback<sandbox_sales_v1_pb.GetSaleOrdersResponse>): grpc.ClientUnaryCall;
  getSaleOrders(argument: sandbox_sales_v1_pb.GetSaleOrdersRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<sandbox_sales_v1_pb.GetSaleOrdersResponse>): grpc.ClientUnaryCall;
  getSaleOrders(argument: sandbox_sales_v1_pb.GetSaleOrdersRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<sandbox_sales_v1_pb.GetSaleOrdersResponse>): grpc.ClientUnaryCall;
  createSaleOrder(argument: sandbox_sales_v1_pb.CreateSaleOrderRequest, callback: grpc.requestCallback<sandbox_sales_v1_pb.CreateSaleOrderResponse>): grpc.ClientUnaryCall;
  createSaleOrder(argument: sandbox_sales_v1_pb.CreateSaleOrderRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<sandbox_sales_v1_pb.CreateSaleOrderResponse>): grpc.ClientUnaryCall;
  createSaleOrder(argument: sandbox_sales_v1_pb.CreateSaleOrderRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<sandbox_sales_v1_pb.CreateSaleOrderResponse>): grpc.ClientUnaryCall;
}

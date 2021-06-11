// package: sandbox.sales.v1
// file: sandbox_sales_v1.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class GetSaleOrderRequest extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSaleOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetSaleOrderRequest): GetSaleOrderRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetSaleOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSaleOrderRequest;
  static deserializeBinaryFromReader(message: GetSaleOrderRequest, reader: jspb.BinaryReader): GetSaleOrderRequest;
}

export namespace GetSaleOrderRequest {
  export type AsObject = {
    id: number,
  }
}

export class GetSaleOrderResponse extends jspb.Message {
  hasSaleOrder(): boolean;
  clearSaleOrder(): void;
  getSaleOrder(): SaleOrder | undefined;
  setSaleOrder(value?: SaleOrder): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSaleOrderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetSaleOrderResponse): GetSaleOrderResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetSaleOrderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSaleOrderResponse;
  static deserializeBinaryFromReader(message: GetSaleOrderResponse, reader: jspb.BinaryReader): GetSaleOrderResponse;
}

export namespace GetSaleOrderResponse {
  export type AsObject = {
    saleOrder?: SaleOrder.AsObject,
  }
}

export class GetSaleOrdersRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSaleOrdersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetSaleOrdersRequest): GetSaleOrdersRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetSaleOrdersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSaleOrdersRequest;
  static deserializeBinaryFromReader(message: GetSaleOrdersRequest, reader: jspb.BinaryReader): GetSaleOrdersRequest;
}

export namespace GetSaleOrdersRequest {
  export type AsObject = {
    email: string,
  }
}

export class GetSaleOrdersResponse extends jspb.Message {
  clearSaleOrdersList(): void;
  getSaleOrdersList(): Array<SaleOrder>;
  setSaleOrdersList(value: Array<SaleOrder>): void;
  addSaleOrders(value?: SaleOrder, index?: number): SaleOrder;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSaleOrdersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetSaleOrdersResponse): GetSaleOrdersResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetSaleOrdersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSaleOrdersResponse;
  static deserializeBinaryFromReader(message: GetSaleOrdersResponse, reader: jspb.BinaryReader): GetSaleOrdersResponse;
}

export namespace GetSaleOrdersResponse {
  export type AsObject = {
    saleOrdersList: Array<SaleOrder.AsObject>,
  }
}

export class CreateSaleOrderRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): void;

  getPaymentMethod(): string;
  setPaymentMethod(value: string): void;

  clearProductsList(): void;
  getProductsList(): Array<Product>;
  setProductsList(value: Array<Product>): void;
  addProducts(value?: Product, index?: number): Product;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSaleOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSaleOrderRequest): CreateSaleOrderRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateSaleOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSaleOrderRequest;
  static deserializeBinaryFromReader(message: CreateSaleOrderRequest, reader: jspb.BinaryReader): CreateSaleOrderRequest;
}

export namespace CreateSaleOrderRequest {
  export type AsObject = {
    email: string,
    paymentMethod: string,
    productsList: Array<Product.AsObject>,
  }
}

export class CreateSaleOrderResponse extends jspb.Message {
  hasSaleOrder(): boolean;
  clearSaleOrder(): void;
  getSaleOrder(): SaleOrder | undefined;
  setSaleOrder(value?: SaleOrder): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSaleOrderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSaleOrderResponse): CreateSaleOrderResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateSaleOrderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSaleOrderResponse;
  static deserializeBinaryFromReader(message: CreateSaleOrderResponse, reader: jspb.BinaryReader): CreateSaleOrderResponse;
}

export namespace CreateSaleOrderResponse {
  export type AsObject = {
    saleOrder?: SaleOrder.AsObject,
  }
}

export class SaleOrder extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getEmail(): string;
  setEmail(value: string): void;

  getPaymentMethod(): string;
  setPaymentMethod(value: string): void;

  hasOrderDate(): boolean;
  clearOrderDate(): void;
  getOrderDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setOrderDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  clearProductsList(): void;
  getProductsList(): Array<Product>;
  setProductsList(value: Array<Product>): void;
  addProducts(value?: Product, index?: number): Product;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SaleOrder.AsObject;
  static toObject(includeInstance: boolean, msg: SaleOrder): SaleOrder.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SaleOrder, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SaleOrder;
  static deserializeBinaryFromReader(message: SaleOrder, reader: jspb.BinaryReader): SaleOrder;
}

export namespace SaleOrder {
  export type AsObject = {
    id: number,
    email: string,
    paymentMethod: string,
    orderDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    productsList: Array<Product.AsObject>,
  }
}

export class Product extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getQuantity(): number;
  setQuantity(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Product.AsObject;
  static toObject(includeInstance: boolean, msg: Product): Product.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Product, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Product;
  static deserializeBinaryFromReader(message: Product, reader: jspb.BinaryReader): Product;
}

export namespace Product {
  export type AsObject = {
    name: string,
    quantity: number,
  }
}


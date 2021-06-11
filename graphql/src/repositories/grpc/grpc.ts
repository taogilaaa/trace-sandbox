import { SaleOrderServiceClient } from '../../generated/sandbox_sales_v1_grpc_pb';
import { GRPC_URL } from '../../generals/constants';
import grpc from 'grpc';

const saleOrderServiceClient = new SaleOrderServiceClient(
  GRPC_URL,
  grpc.credentials.createInsecure(),
);

export { saleOrderServiceClient };

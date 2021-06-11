import { sendMessage } from '../generals/nats';
import { v4 as uuid } from 'uuid';
import {
  getSaleOrder,
  getSaleOrders,
} from '../repositories/grpc/saleOrderRepository';

export type Context = ReturnType<typeof makeContext>;

function makeContext() {
  const requestId = uuid();

  return {
    date: {
      now: () => new Date(),
    },
    nats: {
      sendMessage,
    },
    grpc: {
      getSaleOrder,
      getSaleOrders,
    },
    requestId,
  };
}

export { makeContext };

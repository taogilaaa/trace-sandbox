import { SALEORDER_PLACED_CHANNEL } from '../generals/constants';
import { Context } from './context';

type CreateSaleOrderAsyncArgs = {
  email: string;
  paymentMethod: 'cash' | 'cashless';
  products: Array<{ name: string; quantity: string }>;
};

const resolvers = {
  Query: {
    hello: () => 'world',
    saleOrders: () => {
      return { nodes: [] };
    },
    saleOrder: (_root: any, _args: { id: number }) => {
      return null;
    },
  },
  Mutation: {
    createSaleOrderAsync: async (
      _root: any,
      args: CreateSaleOrderAsyncArgs,
      ctx: Context,
    ) => {
      await ctx.nats.sendMessage({
        channel: SALEORDER_PLACED_CHANNEL,
        message: {
          email: args.email,
          payment_method: args.paymentMethod,
          products: args.products,
        },
      });

      return {
        message: 'OK',
      };
    },
  },
};

export default resolvers;

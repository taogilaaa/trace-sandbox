import { UserInputError } from 'apollo-server';
import { GraphQLResolveInfo } from 'graphql';
import { SALEORDER_PLACED_CHANNEL } from '../generals/constants';
import { timestampToDate } from '../generals/timestamp';
import { Context } from './context';

type CreateSaleOrderAsyncArgs = {
  email: string;
  paymentMethod: 'cash' | 'cashless';
  products: Array<{ name: string; quantity: string }>;
};

const resolvers = {
  Query: {
    hello: () => 'world',
    saleOrders: async (
      _root: any,
      args: { email?: string },
      ctx: Context,
      info: GraphQLResolveInfo,
    ) => {
      const result = await ctx.grpc.getSaleOrders({
        span: info.span,
        email: args.email ?? '',
      });

      let saleOrders = [];
      for (let so of result.saleOrdersList) {
        let saleOrder = {
          ...so,
          orderDate: timestampToDate(so.orderDate!),
          products: so.productsList,
        };
        saleOrders.push(saleOrder);
      }

      return { nodes: saleOrders };
    },
    saleOrder: async (
      _root: any,
      args: { id: number },
      ctx: Context,
      info: GraphQLResolveInfo,
    ) => {
      const result = await ctx.grpc.getSaleOrder({
        span: info.span,
        id: args.id,
      });

      if (!result.saleOrder) {
        throw new UserInputError(`Invalid pickedDocument id ${args.id}`);
      }

      return {
        ...result.saleOrder,
        orderDate: timestampToDate(result.saleOrder.orderDate!),
        products: result.saleOrder.productsList,
      };
    },
  },
  Mutation: {
    createSaleOrderAsync: async (
      _root: any,
      args: CreateSaleOrderAsyncArgs,
      ctx: Context,
      info: GraphQLResolveInfo,
    ) => {
      await ctx.nats.sendMessage(
        {
          channel: SALEORDER_PLACED_CHANNEL,
          message: {
            email: args.email,
            payment_method: args.paymentMethod,
            products: args.products,
          },
        },
        info.span,
      );

      return {
        message: 'OK',
      };
    },
  },
};

export default resolvers;

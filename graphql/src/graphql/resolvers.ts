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
      _args: CreateSaleOrderAsyncArgs,
    ) => {
      return {
        message: 'OK',
      };
    },
  },
};

export default resolvers;

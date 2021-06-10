type CreateSaleOrderArgs = {
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
    createSaleOrder: (_root: any, args: CreateSaleOrderArgs) => {
      return {
        ...args,
        id: 1,
        orderDate: new Date().toISOString(),
      };
    },
  },
};

export default resolvers;

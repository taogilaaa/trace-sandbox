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
};

export default resolvers;

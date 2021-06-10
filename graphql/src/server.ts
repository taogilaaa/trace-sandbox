import { ApolloServer } from 'apollo-server';
import { schema } from './graphql/schema';

const server = new ApolloServer({
  schema: schema,
  playground: {
    settings: {
      // Disable auto refresh on graphql playground
      'schema.polling.enable': false,
    } as any,
  },
});

export { server };

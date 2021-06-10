import { ApolloServer } from 'apollo-server';
import { makeContext } from './graphql/context';
import { schema } from './graphql/schema';

const server = new ApolloServer({
  schema: schema,
  context: () => {
    return makeContext();
  },
  playground: {
    settings: {
      // Disable auto refresh on graphql playground
      'schema.polling.enable': false,
    } as any,
  },
});

export { server };

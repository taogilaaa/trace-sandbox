import { ApolloServer } from 'apollo-server';
import { tracerExtension } from './extensions/tracer';
import { makeContext } from './graphql/context';
import { schema } from './graphql/schema';

const server = new ApolloServer({
  schema: schema,
  context: () => {
    return makeContext();
  },
  introspection: true,
  playground: {
    settings: {
      // Disable auto refresh on graphql playground
      'schema.polling.enable': false,
    } as any,
  },
  extensions: [tracerExtension],
});

export { server };

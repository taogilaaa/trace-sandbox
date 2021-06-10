import { makeExecutableSchema } from 'apollo-server';
import resolvers from './resolvers';
import typeDefs from './schema.graphql';

export const schema = makeExecutableSchema({ typeDefs, resolvers });

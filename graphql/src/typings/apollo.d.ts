import { Span } from 'opentracing';

type Span = Span;

declare module 'graphql/type/definition' {
  interface GraphQLResolveInfo {
    span: Span;
  }
}

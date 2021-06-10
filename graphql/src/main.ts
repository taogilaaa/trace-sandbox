import { server } from './server';

void server.listen({ port: 50040, path: `/graphql` }).then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
  console.log(
    `Try your health check at: ${url}.well-known/apollo/server-health`,
  );
});

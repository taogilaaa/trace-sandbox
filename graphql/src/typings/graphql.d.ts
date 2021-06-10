// Fixes: https://github.com/prisma/graphql-import-loader/issues/11#issuecomment-384432736
declare module '*.graphql' {
  const value: any;
  export = value;
}

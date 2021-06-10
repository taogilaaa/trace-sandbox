module.exports = {
  roots: ['<rootDir>/src'],
  testRegex: '(.*(-|.)test\\.(tsx?|jsx?))$',
  transform: {
    '^.+\\.tsx?$': 'babel-jest',
    '^.+\\.jsx?$': 'babel-jest',
    '\\.graphql$': 'jest-raw-loader',
  },
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx', 'json', 'node'],
};

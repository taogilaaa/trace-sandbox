const base = require('tscripts/config/babel.config');

module.exports = function (api) {
  const baseConfig = base(api);
  const plugins = [...baseConfig.plugins, 'babel-plugin-inline-import'];

  return {
    ...baseConfig,
    plugins,
  };
};

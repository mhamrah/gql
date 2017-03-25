module.exports = {
  entry: "./app.js",
  output: {
    filename: "public/bundle.js"
  },
  module: {
    rules: [
      {
        test: /\.(js)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: [['env', { targets: {"browsers": ["last 2 versions", "safari >= 7"]}}]],
            plugins: ['babel-plugin-transform-react-jsx']
          }
        },
      },
    ],
  },
}
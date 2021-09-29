const path = require('path')

module.exports = {
  entry: './src/app.js',
  mode: process.env.NODE_ENV || 'production',
  output: {
    path: path.resolve(__dirname, 'cmd', 'public', 'scripts'),
    filename: 'bundle.js'
  },
  resolve: {
    alias: { vue$: 'vue/dist/vue.esm.js' }
  }
}

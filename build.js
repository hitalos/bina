const esBuild = require('esbuild')
const vuePlugin = require('esbuild-vue')

esBuild.build({
  entryPoints: ['src/app.js'],
  bundle: true,
  minify: true,
  outfile: 'cmd/public/scripts/bundle.js',
  plugins: [vuePlugin({ production: false })],
  define: {
    'process.env.NODE_ENV': JSON.stringify('production'),
  },
}).then((result) => {
  console.log(result)
})

const esBuild = require('esbuild')
const vuePlugin = require('esbuild-vue')

const devMode = process.env.NODE_ENV === 'dev'
const options = {
  bundle: true,
  minify: !devMode,
}

esBuild.build({
  ...options,
  entryPoints: ['src/app.js'],
  outfile: 'cmd/public/scripts/bundle.js',
  plugins: [vuePlugin({ production: !devMode })],
}).catch(console.error)

esBuild.build({
  ...options,
  entryPoints: ['src/app.css'],
  outfile: 'cmd/public/styles/bundle.css',
}).catch(console.error)

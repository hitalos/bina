const fs = require('fs')
const gulp = require('gulp')
const clean = require('gulp-clean')
const concat = require('gulp-concat')
const concatCss = require('gulp-concat-css')
const cssnano = require('gulp-cssnano')

const env = process.env.NODE_ENV

gulp.task('scripts', ['clean-js'], () => {
  gulp.src(['src/resources/scripts/app.js'])
  .pipe(gulp.dest('public/scripts'))
  let min = ''
  if (env === 'production' || env === 'prod') {
    min = '.min'
  }
  return gulp.src([
    `node_modules/vue/dist/vue${min}.js`,
    'node_modules/vue-material/dist/vue-material.js',
    `node_modules/vuex/dist/vuex${min}.js`,
    `node_modules/axios/dist/axios${min}.js`,
  ])
  .pipe(concat('vendors.js'))
  .pipe(gulp.dest('public/scripts'))
})

gulp.task('styles', ['clean-css'], () =>
  gulp.src([
    'node_modules/vue-material/dist/vue-material.css',
    'src/resources/styles/app.css'
  ])
  .pipe(concatCss('app.css'))
  .pipe(cssnano())
  .pipe(gulp.dest('public/styles'))
)

gulp.task('copy-env', () => {
  fs.stat('.env', (err) => {
    if (err) {
      console.log('Copying ".env" file…')
      fs.createReadStream('.env-example').pipe(fs.createWriteStream('.env'))
    }
  })
})

gulp.task('clean-js', () => gulp.src('public/scripts').pipe(clean()))
gulp.task('clean-css', () => gulp.src('public/styles').pipe(clean()))

gulp.task('watch', () => {
  gulp.watch('src/resources/scripts/*.js', ['scripts'])
  gulp.watch('src/resources/styles/*.css', ['styles'])
})

gulp.task('default', ['scripts', 'styles'])

const compression = require('compression')
const express = require('express')
const helmet = require('helmet')
const path = require('path')

const routes = require('./routes')

process.on('uncaughtException', console.error)

const app = express()
app.use(helmet())
app.use(compression())

app.set('view engine', 'pug')
app.set('views', './src/resources/views')

app.get('/', (req, res) => {
  res.render('index')
})

app.use(express.static(path.join(__dirname, '/../public')))
app.use(routes)

app.use((req, res, next) => {
  const error = new Error('Not Found')
  error.status = 404
  next(error)
})

/* eslint no-unused-vars: 0 */
app.use((error, req, res, next) => {
  res.status(error.status || 500)
  if (app.get('env') === 'development') {
    res.render('error', { error })
  } else {
    res.render('error', { error: { message: error.message } })
  }
})

module.exports = app

const http = require('http')
const router = require('express').Router()

const contacts = require('./contacts')

router.use('/contacts', contacts)
router.get('/images/logo.png', (req, res) => {
  http.get(process.env.LOGO_URL, (response) => {
    response.pipe(res)
  })
})

module.exports = router

const router = require('express').Router()

const contacts = require('./contacts')

router.use('/contacts', contacts)
router.get('/images/logo.png', (req, res) => {
  res.redirect(process.env.LOGO_URL)
})

module.exports = router

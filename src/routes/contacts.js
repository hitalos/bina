const router = require('express').Router()

const ldapService = require('../ldapService')

router.get('/all.json', (req, res) => {
  ldapService((err, result) => {
    if (err) throw err
    res.send(result)
  })
})

router.get('/:contact/photo.jpg', (req, res) => {
  res.redirect(`${process.env.PHOTOS_URL}/${req.params.contact}.jpg`)
})

module.exports = router

const crypto = require('crypto')
const debug = require('debug')('Bina:Photos')
const http = require('http')

const ldapService = require('../ldapService')

module.exports = (req, res, next) => {
  if (process.env.ENABLE_GRAVATAR === 'true') {
    ldapService((err, result) => {
      if (err) throw err
      const contact = result.filter(item => item.id === req.params.contact)[0]
      if (!contact.emails.mail) {
        debug('Contact mail not found in ldap')
        res.redirect(process.env.LOGO_URL)
        return
      }
      const email = contact.emails.mail
      const md5Hash = crypto.createHash('md5').update(email).digest('hex')

      debug('Redirecting to Gravatar image')
      res.redirect(`http://www.gravatar.com/avatar/${md5Hash}`)
    })
  } else {
    const url = `${process.env.PHOTOS_URL}${req.params.contact}.jpg`
    debug(`Reading image from "${url}"`)
    http.get(url, (response) => {
      if (response.statusCode === 200) {
        res.set('Content-Type', response.headers['content-type'])
        const rawData = []
        response.on('data', (chunk) => { rawData.push(chunk) })
        response.on('end', () => {
          res.send(Buffer.concat(rawData))
          res.end()
        })
      } else {
        debug('Redirecting image to default icon')
        res.redirect('/images/default.png')
      }
    })
  }
}

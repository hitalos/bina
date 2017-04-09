const ldapjs = require('ldapjs')

const idField = process.env.IDENT_FIELD
const fullNameField = process.env.FULL_NAME_FIELD
const phonesFields = process.env.PHONE_FIELDS.split(',')
const otherFields = process.env.OTHER_FIELDS.split(',')
const emailFields = process.env.EMAIL_FIELDS.split(',')

const credentials = {
  user: process.env.LDAP_USER,
  pass: process.env.LDAP_PASS,
}
const ldapClient = ldapjs.createClient({
  url: process.env.LDAP_HOST,
  tlsOptions: {
    rejectUnauthorized: false,
  },
  connectTimeout: 10000,
  timeout: 10000,
})

ldapClient.on('error', (err) => {
  console.log(err)
  console.log(`Error connecting with ${process.env.LDAP_HOST}`)
  process.exit(1)
})

ldapClient.on('timeout', () => {
  console.log('Timeout consulting LDAP SERVER')
  process.exit(1)
})

ldapClient.on('connectTimeout', () => {
  console.log(`Timeout connecting with ${process.env.LDAP_HOST}`)
  process.exit(1)
})

const resultCache = {
  time: new Date(),
  duration: process.env.CACHE_DURATION || 5 * 60 * 1000, // default 5min
  data: null,
  expired() {
    return this.data ? (this.time.valueOf() + this.duration) < Date.now() : true
  },
  setData(data) {
    this.time = new Date()
    this.data = data
  },
}

module.exports = (cb) => {
  if (!resultCache.expired()) {
    // Using cache result
    cb(null, resultCache.data)
    return
  }
  ldapClient.bind(credentials.user, credentials.pass, (bindError) => {
    if (bindError) {
      cb(bindError, null)
      return
    }
    const base = process.env.LDAP_BASE
    /* eslint prefer-template: 0 */
    const filter = '(&' +
      '(|' +
      phonesFields.map(item => `(${item}=*)`).join('') +
      ')' +
      '(objectCategory=person)' +
      '(!(UserAccountControl:1.2.840.113556.1.4.803:=2))' + // User active
      '(|' +
        '(objectClass=user)' + // User object
        '(objectClass=contact)' + // Contact object
      ')' +
    ')'
    const attributes = ['objectClass'].concat(
      fullNameField,
      phonesFields,
      otherFields,
      idField,
      emailFields
    )
    const options = {
      scope: 'sub',
      paged: true,
      sizeLimit: 100,
      filter,
      attributes,
    }
    ldapClient.search(base, options, (errSearch, result) => {
      if (errSearch) {
        cb(errSearch, null)
        return
      }
      const list = []
      result.on('searchEntry', (entry) => {
        const contact = {}
        contact.id = entry.object[idField]
        contact.fullName = entry.object[fullNameField]
        contact.phones = {}
        phonesFields.forEach(
          (phone) => {
            if (entry.object[phone]) {
              contact.phones[phone] = entry.object[phone]
            }
          }
        )
        contact.emails = {}
        emailFields.forEach(
          (email) => {
            if (entry.object[email]) {
              contact.emails[email] = entry.object[email]
            }
          }
        )
        otherFields.forEach(
          (field) => {
            if (entry.object[field]) {
              contact[field] = entry.object[field]
            }
          }
        )
        if (entry.object.objectClass.some(oc => oc === 'user')) {
          contact.objectClass = 'user'
        } else {
          contact.objectClass = 'contact'
        }
        list.push(contact)
      })
      result.on('end', () => {
        // Order by displayName
        list.sort((a, b) => a.fullName.localeCompare(b.fullName))

        // set cache
        resultCache.setData(list)
        cb(null, list)
      })
    })
  })
}

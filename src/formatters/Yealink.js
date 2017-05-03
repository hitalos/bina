const libxml = require('libxmljs')

module.exports = (result) => {
  const xml = new libxml.Document('1.0', 'UTF-8')
  const addressBook = xml.node('customIPPhoneDirectory')
  result.forEach((contact) => {
    if (contact.phones.ipPhone) {
      const el = addressBook.node('DirectoryEntry')
      el.node('Name', contact.fullName)
      Object.keys(contact.phones).forEach((phone) => {
        el.node('Telephone', phone)
      })
    }
  })
  return xml
}

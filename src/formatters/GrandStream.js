const libxml = require('libxmljs')

module.exports = (result) => {
  const xml = new libxml.Document('1.0', 'UTF-8')
  const addressBook = xml.node('AddressBook')
  result.forEach((contact) => {
    if (contact.phones.ipPhone) {
      const el = addressBook.node('Contact')
      el.node('FirstName', contact.fullName.split(' ')[0])
      el.node('LastName', contact.fullName.split(' ').slice(-1).pop())
      el.node('Phone').node('phonenumber', contact.phones.ipPhone)
    }
  })
  addressBook.attr('count', addressBook.childNodes().length)
  return xml
}

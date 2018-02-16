/* eslint no-unused-vars: 0 no-undef: 0 */
const show = (contact, searchTerms) => {
  const terms = removeAccents(searchTerms).toLowerCase()
    .trim()
    .replace(/\s\s+/g, ' ')
    .split(' ')
  const fullName = removeAccents(contact.fullName).toLowerCase()
  const phones = Object.keys(contact.phones).map(key => contact.phones[key])

  if (terms.every(str => fullName.indexOf(str) >= 0)) return true
  if (contact.department) {
    const pdon = contact.physicalDeliveryOfficeName
    const department = removeAccents(`${contact.department} ${pdon || ''}`)
      .toLowerCase()
    if (terms.every(str => department.indexOf(str) >= 0)) return true
  }
  if (contact.title) {
    const title = removeAccents(contact.title).toLowerCase()
    if (terms.every(str => title.indexOf(str) >= 0)) return true
  }

  return terms.some(str => phones.some(phone => phone.indexOf(str) >= 0))
}

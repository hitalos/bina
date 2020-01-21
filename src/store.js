import axios from 'axios'

const removeAccents = (str) => [
  { chr: 'a', regex: /[ÀÁÂÃÄȂàáâãäª]/ },
  { chr: 'e', regex: /[ÉÊËéêë]/ },
  { chr: 'i', regex: /[ÍÎÏíîï]/ },
  { chr: 'o', regex: /[ÓÔÕÖóôõöº]/ },
  { chr: 'u', regex: /[ÚÛÜúûü]/ },
  { chr: 'c', regex: /[ÇḈç]/ },
  { chr: 'n', regex: /[Ññ]/ }
].reduce((acum, accent) => acum.replace(accent.regex, accent.chr), str)

const show = (contact, searchTerms) => {
  const terms = removeAccents(searchTerms).toLowerCase()
    .trim()
    .replace(/\s\s+/g, ' ')
    .split(' ')
  const fullName = removeAccents(contact.fullName).toLowerCase()
  const phones = Object.keys(contact.phones).map((key) => contact.phones[key])

  if (terms.every((str) => fullName.indexOf(str) >= 0)) return true
  if (contact.others) {
    if (contact.others.department) {
      const pdon = contact.others.physicalDeliveryOfficeName
      const department = removeAccents(`${contact.others.department} ${pdon || ''}`)
        .toLowerCase()
      if (terms.every((str) => department.indexOf(str) >= 0)) return true
    }
    if (contact.others.title) {
      const title = removeAccents(contact.others.title).toLowerCase()
      if (terms.every((str) => title.indexOf(str) >= 0)) return true
    }
  }

  return terms.some((str) => phones.some((phone) => phone.indexOf(str) >= 0))
}
/* eslint no-new: 0 */
export default (Vuex) => new Vuex.Store({
  state: {
    contacts: [],
    searchTerms: ''
  },
  getters: {
    total(state) {
      return state.contacts.length
    },
    count(state) {
      return state.contacts.filter((contact) => contact.show).length
    },
    limitedList(state) {
      return state.contacts
        .filter((contact) => contact.show)
        .slice(0, 30)
    },
  },
  mutations: {
    /* eslint no-param-reassign: 0, no-alert: 0 */
    populate(state) {
      axios.get('/contacts/all.json').then((response) => {
        state.contacts = response.data.map((contact) => Object.assign(contact, { show: false }))
      }).catch((err) => {
        if (err.response.data.name === 'InvalidCredentialsError') {
          alert('Erro!\nCredenciais do usuário de sistema inválidas.\nContate um administrador do domínio!')
        }
        console.log(err.response.data)
      })
    },
    filterChanged(state, searchTerms) {
      state.searchTerms = searchTerms
      state.contacts = state.contacts.map((contact) => Object.assign(contact, { show: show(contact, searchTerms) }))
    },
  },
})

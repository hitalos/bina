/* eslint no-undef: 0, no-new: 0, no-param-reassign: 0 */
const store = new Vuex.Store({
  state: {
    contacts: [],
    searchTerms: ''
  },
  getters: {
    total(state) {
      return state.contacts.length
    },
    count(state) {
      return state.contacts.filter(contact => contact.show).length
    },
    limitedList(state) {
      return state.contacts
        .filter(contact => contact.show)
        .slice(0, 30)
    },
  },
  mutations: {
    populate(state) {
      axios.get('/contacts/all.json').then((response) => {
        state.contacts = response.data.map(contact =>
          Object.assign(contact, { show: false })
        )
      })
    },
    filterChanged(state, searchTerms) {
      state.searchTerms = searchTerms
      state.contacts = state.contacts.map(contact =>
        Object.assign(contact, { show: show(contact, searchTerms) })
      )
    },
  },
})

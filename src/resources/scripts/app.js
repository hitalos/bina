/* eslint no-undef: 0, no-new: 0, no-param-reassign: 0 */
const removeAccents = str =>
  [
    { chr: 'a', regex: /[ÀÁÂÃÄȂàáâãäª]/ },
    { chr: 'e', regex: /[ÉÊËéêë]/ },
    { chr: 'i', regex: /[ÍÎÏíîï]/ },
    { chr: 'o', regex: /[ÓÔÕÖóôõöº]/ },
    { chr: 'u', regex: /[ÚÛÜúûü]/ },
    { chr: 'c', regex: /[ÇḈç]/ },
    { chr: 'n', regex: /[Ññ]/ }
  ].reduce((acum, accent) =>
    acum.replace(accent.regex, accent.chr),
    str
  )

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

const SearchField = Vue.component('search-field', {
  computed: {
    count() {
      return this.$store.getters.count
    },
    total() {
      return this.$store.getters.total
    },
  },
  template: '#search-field-template',
  methods: {
    filterChanged(e) {
      const val = e.target.value.trim()
      if (val.length >= 3) {
        this.$store.commit('filterChanged', val)
      }
    },
  },
})

const CardList = Vue.component('card-list', {
  computed: {
    limitedList() {
      return this.$store.getters.limitedList
    },
  },
  template: '#card-list-template',
})

const Counters = Vue.component('counters', {
  computed: {
    limitedList() {
      return this.$store.getters.limitedList
    },
    count() {
      return this.$store.getters.count
    },
  },
  template: '#counters-template'
})

const ContactCard = Vue.component('contact-card', {
  props: ['contact'],
  data() {
    return {
      labels: {
        telephoneNumber: 'Principal',
        mobile: 'Celular',
        ipPhone: 'VoIP',
        facsimileTelephoneNumber: 'Fax',
        homePhone: 'Casa',
        otherTelephone: 'Outro'
      },
    }
  },
  computed: {
    defaultPhone() {
      return this.contact.phones[Object.keys(this.contact.phones)[0]]
    },
  },
  template: '#contact-card-template',
  methods: {
    invert() {
      const el = this.$refs['flip-container']
      console.log('invert')
      if (el.classList) el.classList.add('reverse')
      else el.className += ' reverse'
    },
    revert() {
      const el = this.$refs['flip-container']
      console.log('revert')
      if (el.classList) el.classList.remove('reverse')
      else {
        el.className = el.className.replace(
          new RegExp(`(^|\\b)${className.split(' ').join('|')}(\\b|$)`, 'gi'),
          ' '
        )
      }
    },
  },
})

Vue.use(VueMaterial)
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
new Vue({
  el: '#app',
  name: 'App',
  store,
  template: '#app-template',
  components: {
    SearchField,
    CardList,
    ContactCard,
    Counters,
  },
  beforeCreate() {
    this.$store.commit('populate')
  },
})

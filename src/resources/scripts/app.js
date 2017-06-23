/* eslint no-undef: 0, no-new: 0 */
function removeAccents(str) {
  let strWithoutAccents = str
  const replaces = [
    { chr: 'a', regex: /[ÀÁÂÃÄȂàáâãäª]/ },
    { chr: 'e', regex: /[ÉÊËéêë]/ },
    { chr: 'i', regex: /[ÍÎÏíîï]/ },
    { chr: 'o', regex: /[ÓÔÕÖóôõöº]/ },
    { chr: 'u', regex: /[ÚÛÜúûü]/ },
    { chr: 'c', regex: /[ÇḈç]/ },
    { chr: 'n', regex: /[Ññ]/ }
  ]
  replaces.forEach((accent) => {
    strWithoutAccents = strWithoutAccents.replace(accent.regex, accent.chr)
  })
  return strWithoutAccents
}

function show(contact, searchTerms) {
  const terms = removeAccents(searchTerms).toLowerCase().trim()
    .replace(/\s\s+/g, ' ')
    .split(' ')
  const fullName = removeAccents(contact.fullName).toLowerCase()
  const phones = Object.keys(contact.phones).map(key => contact.phones[key])

  if (terms.every(str => fullName.indexOf(str) >= 0)) return true
  if (contact.department) {
    const department = removeAccents(`${contact.department} ${contact.physicalDeliveryOfficeName || ''}`).toLowerCase()
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
  template:
    `<md-layout class="md-flex-20 md-flex-small-33 md-flex-xsmall-100">
      <md-input-container>
        <label>Busca</label>
        <md-input tabindex="1" @keyup.native="filterChanged"/>
      </md-input-container>
    </md-layout>`,
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
  template:
    `<md-layout md-flex class="card-list">
      <contact-card
        class="md-flex-xlarge-20 md-flex-large-33 md-flex-medium-50 md-flex-small-50 md-flex-xsmall-100"
        v-for="contact in limitedList"
        :key="contact.fullName"
        :contact="contact"
      />
    </md-layout>`,
  },
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
  template:
    `<div class="flip-container" ref="flip-container">
      <div class="flipper">
        <div class="front">
          <md-layout :class="contact.objectClass">
            <md-card class="md-flex-100 md-with-hover">
              <md-card-header :title='"Vínculo: " + contact.title'>
                <md-card-header-text>
                  <div class="md-headline">{{ contact.fullName }}</div>
                </md-card-header-text>
                <md-card-media>
                  <img
                    v-if="contact.id"
                    :src="'/contacts/' + contact.id + '.jpg'"
                    alt="Foto"
                  />
                  <img
                    v-if="!contact.id"
                    src="/images/logo.png"
                  />
                </md-card-media>
              </md-card-header>
              <md-card-content>
                <p v-if="contact.department">Lotação:
                  <strong>{{ contact.department }}
                    <span v-if="contact.physicalDeliveryOfficeName">
                    - {{ contact.physicalDeliveryOfficeName }}
                    </span>
                  </strong>
                </p>
              </md-card-content>
              <md-card-actions>
                <md-button :title='defaultPhone'>
                  <a class="md-display-2" :href="'tel:' + defaultPhone">
                    {{ defaultPhone }}
                  </a>
                </md-button>
                <div class="md-flex"/>
                <md-button @click.native='invert'>Ver mais detalhes</md-button>
              </md-card-actions>
            </md-card>
          </md-layout>
        </div>
        <div class="back">
          <md-layout :class="contact.objectClass">
            <md-card class="md-flex-100 md-with-hover">
              <md-card-content>
                <p v-if="contact.title">
                  Vínculo: <strong>{{ contact.title }}</strong>
                </p>
                <p v-if="contact.emails.mail">
                  Mail: <strong>{{ contact.emails.mail }}</strong>
                </p>
                <p v-for="(phone, key) in contact.phones">
                  {{ labels[key] }}: <strong>{{ phone }}</strong>
                </p>
              </md-card-content>
              <md-card-actions>
                <md-button v-if="contact.objectClass=='user'">
                  <a :href="'/contacts/' + contact.id + '.vcf'">Baixar vCard</a>
                </md-button>
                <div class="md-flex"/>
                <md-button @click.native="revert">voltar</md-button>
              </md-card-actions>
            </md-card>
          </md-layout>
        </div>
      </div>
    </div>`,
  methods: {
    invert() {
      const el = this.$refs['flip-container']
      if (el.classList) el.classList.add('reverse')
      else el.className += ' reverse'
    },
    revert() {
      const el = this.$refs['flip-container']
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
  template:
    `<div id="app" class="phone-viewport">
      <md-toolbar>
        <md-layout>
          <h1 class="md-title">Bina</h1>
        </md-layout>
        <search-field/>
      </md-toolbar>
      <card-list/>
    </div>`,
  components: {
    SearchField,
    CardList,
    ContactCard,
  },
  beforeCreate() {
    this.$store.commit('populate')
  },
})

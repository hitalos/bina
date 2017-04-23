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
  const fullName = removeAccents(contact.fullName).toLowerCase()
  const str = removeAccents(searchTerms).toLowerCase()
  if (fullName.indexOf(str) >= 0) return true
  if (contact.department) {
    const department = removeAccents(contact.department).toLowerCase()
    if (department.indexOf(str) >= 0) return true
  }
  if (contact.title) {
    const title = removeAccents(contact.title).toLowerCase()
    if (title.indexOf(str) >= 0) return true
  }
  return false
}

const bus = new Vue()

const SearchField = Vue.component('search-field', {
  data() {
    return { searchTerms: '' }
  },
  template:
    `<md-layout class="md-flex-20 md-flex-small-33 md-flex-xsmall-100">
      <md-input-container>
        <label>Busca</label>
        <md-input tabindex="1" v-model="searchTerms" @change="filterChanged"/>
      </md-input-container>
    </md-layout>`,
  methods: {
    filterChanged(searchTerms) {
      if (searchTerms.trim().length >= 3) {
        bus.$emit('filter-changed', searchTerms)
      }
    }
  },
})

const CardList = Vue.component('card-list', {
  template:
    `<md-layout md-flex>
      <contact-card
        class="md-flex-xlarge-20 md-flex-large-33 md-flex-medium-50 md-flex-small-50 md-flex-xsmall-100"
        v-if="contact.show"
        v-for="contact in contacts"
        :key="contact.fullName"
        :contact="contact"
      />
    </md-layout>`,
  data() {
    return { contacts: [] }
  },
  created() {
    this.$http.get('/contacts/all.json').then((response) => {
      this.$data.contacts = response.body.map((contact) => {
        contact.show = false
        return contact
      })
    })
    bus.$on('filter-changed', (searchTerms) => {
      this.contacts.forEach((contact) => {
        contact.show = show(contact, searchTerms)
      })
    })
  },
})

const ContactCard = Vue.component('contact-card', {
  props: ['contact'],
  template:
    `<div class="flip-container">
      <div class="flipper">
        <div class="front">
          <md-layout :class="contact.objectClass">
            <md-card class="md-flex-100 md-with-hover">
              <md-card-header>
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
                <p v-if="contact.department">Lotação: <strong>{{ contact.department }}
                  <span v-if="contact.physicalDeliveryOfficeName"> - {{ contact.physicalDeliveryOfficeName }}</span></strong>
                </p>
              </md-card-content>
              <md-card-actions>
                <md-button>
                  <a class="md-display-2" :href="'tel:' + contact.phones.ipPhone">
                    {{ contact.phones.ipPhone }}
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
                <p v-if="contact.title">Vínculo: <strong>{{ contact.title }}</strong></p>
                <p v-if="contact.emails.mail">Mail: <strong>{{ contact.emails.mail }}</strong></p>
                <p v-if="contact.phones.mobile">Celular: <strong>{{ contact.phones.mobile }}</strong></p>
                <p v-if="contact.phones.facsimileTelephoneNumber">Fax: <strong>{{ contact.phones.facsimileTelephoneNumber }}</strong></p>
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
    invert(event) {
      const el = event.target.parentNode.parentNode.parentNode.parentNode.parentNode.parentNode
      if (el.classList) el.classList.add('reverse')
      else el.className += ' reverse'
    },
    revert(event) {
      const el = event.target.parentNode.parentNode.parentNode.parentNode.parentNode.parentNode
      if (el.classList) el.classList.remove('reverse')
      else el.className = el.className.replace(new RegExp(`(^|\\b)${className.split(' ').join('|')}(\\b|$)`, 'gi'), ' ')
    },
  },
})

Vue.use(VueMaterial)
new Vue({
  el: '#app',
  name: 'App',
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
})

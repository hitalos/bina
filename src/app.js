import Vue from 'vue'
import Vuex from 'vuex'
import VueMaterial from 'vue-material'

import CardList from './components/cardList'
import ContactCard from './components/contactCard'
import Counters from './components/counters'
import SearchField from './components/searchField'
import store from './store'

Vue.use(Vuex)
Vue.use(VueMaterial)

/* eslint no-new: 0 */
new Vue({
  el: '#app',
  name: 'App',
  store: store(Vuex),
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
  methods: {
    searchFocus() {
      document.querySelectorAll('.md-input')[0].focus()
    },
  },
})

if (window.location.protocol === 'https:' && !navigator.serviceWorker.controller) {
  navigator.serviceWorker.register('pwabuilder-sw.js', {
    scope: './'
  })
}

/* eslint no-undef: 0, no-new: 0 */
Vue.use(VueMaterial)

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

/* eslint no-unused-vars: 0 no-undef: 0 */
const CardList = Vue.component('card-list', {
  computed: {
    limitedList() {
      return this.$store.getters.limitedList
    },
  },
  template: '#card-list-template',
})

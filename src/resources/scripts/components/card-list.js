const CardList = Vue.component('card-list', {
  computed: {
    limitedList() {
      return this.$store.getters.limitedList
    },
  },
  template: '#card-list-template',
})

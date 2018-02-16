/* eslint no-unused-vars: 0 no-undef: 0 */
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

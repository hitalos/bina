import Vue from 'vue'

export default Vue.component('counters', {
  computed: {
    limitedList() {
      return this.$store.getters.limitedList
    },
    count() {
      return this.$store.getters.count
    },
  },
  template: '#counters-template',
})

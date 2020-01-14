import Vue from 'vue'

export default Vue.component('card-list', {
  computed: {
    limitedList() {
      return this.$store.getters.limitedList
    },
  },
  template: '#card-list-template',
})

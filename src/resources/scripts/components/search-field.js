/* eslint no-unused-vars: 0 no-undef: 0 */
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

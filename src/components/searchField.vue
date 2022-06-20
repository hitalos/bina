<template>
	<md-layout class="md-flex-20 md-flex-small-33 md-flex-xsmall-100">
		<md-input-container>
			<label>Busca<md-input tabindex="1" @keyup.native="filterChanged"></md-input></label>
		</md-input-container>
	</md-layout>
</template>


<script>
import Vue from 'vue'

export default Vue.component('search-field', {
  computed: {
    count() {
      return this.$store.getters.count
    },
    total() {
      return this.$store.getters.total
    },
  },
  methods: {
    filterChanged(e) {
      const val = e.target.value.trim()
      if (val.length >= 3) {
        this.$store.commit('filterChanged', val)
      }
    },
  },
})
</script>

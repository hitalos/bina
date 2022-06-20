<template>
	<div class="phone-viewport">
		<md-toolbar @click.native="searchFocus">
			<md-layout>
				<h1 class="md-title">Bina</h1>
			</md-layout>
			<search-field></search-field>
		</md-toolbar>
		<card-list></card-list>
	</div>
</template>

<script>
import Vue from 'vue'

import CardList from './cardList.vue'
import ContactCard from './contactCard.vue'
import Counters from './counters.vue'
import SearchField from './searchField.vue'

export default Vue.component('app', {
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
</script>

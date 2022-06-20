<template>
	<md-layout class="md-flex-100" v-if="count">
		<p v-if="count &gt;= limitedList.length">Encontrado(s) {{ count }} contato(s).</p>
		<p v-if="count &gt; limitedList.length">Mostrando apenas os primeiros {{ limitedList.length }}.</p>
	</md-layout>
</template>

<script>
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
})
</script>

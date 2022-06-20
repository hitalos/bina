<template>
<md-layout class="card-list" md-flex>
  <counters></counters>
  <contact-card class="md-flex-xlarge-20 md-flex-large-33 md-flex-medium-50 md-flex-small-50 md-flex-xsmall-100" v-for="contact in limitedList" :key="contact.fullName" :contact="contact"></contact-card>
  </md-layout>
</template>

<script>
import Vue from 'vue'

export default Vue.component('card-list', {
  computed: {
    limitedList() {
      return this.$store.getters.limitedList
    },
  },
})
</script>

import Vue from 'vue'
import Vuex from 'vuex'
import VueMaterial from 'vue-material'

import App from './components/app.vue'
import store from './store'

Vue.use(Vuex)
Vue.use(VueMaterial)

/* eslint no-new: 0 */
new Vue({
  el: '#app',
  store: store(Vuex),
  render: (h) => h(App),
})

if (window.location.protocol === 'https:' && !navigator.serviceWorker.controller) {
  navigator.serviceWorker.register('pwabuilder-sw.js', {
    scope: './',
  })
}

import Vue from 'vue'
import App from './App.vue'
import '@babel/polyfill'
import vuetify from './plugins/vuetify'
import i18n from '@/plugins/i18n'
import router from './routes'

Vue.config.productionTip = false;

new Vue({
  router,
  i18n,
  vuetify,
  render: h => h(App),
}).$mount('#app');

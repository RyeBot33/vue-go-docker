//import { createApp } from 'vue'
import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';

//createApp(App).mount('#app')
new Vue({
  vuetify,
  render: h => h(App)
}).$mount("#app");

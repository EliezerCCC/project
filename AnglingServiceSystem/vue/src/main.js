import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import HeaderTop from './components/HeaderTop.vue'
import 'element-theme-chalk';

Vue.use(ElementUI)
Vue.use(VueAxios, axios)
Vue.config.productionTip = false
Vue.component('HeaderTop',HeaderTop);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

import axios from "axios";
import * as echarts from "echarts";
import "element-theme-chalk";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import Vue from "vue";
import VueAxios from "vue-axios";
import App from "./App.vue";
import HeaderTop from "./components/HeaderTop.vue";
import LeftManager from "./components/Left.vue";
import QuillEditor from "./components/QuillEditor.vue";
import global from "./global/global.vue";
import router from "./router";

Vue.prototype.global = global;
Vue.prototype.$echarts = echarts;
Vue.use(ElementUI);
Vue.use(VueAxios, axios);
Vue.config.productionTip = false;
Vue.component("HeaderTop", HeaderTop);
Vue.component("LeftManager", LeftManager);
Vue.component("QuillEditor", QuillEditor);

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");

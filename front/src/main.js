// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store/store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';
import Header from "./components/Header";
import Vue2TouchEvents from "vue2-touch-events";
import connector from "./connector/connector";

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(Vue2TouchEvents)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  connector,
  components: { App, Header},
  template: '<App/>'
})

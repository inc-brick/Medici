// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueAnalytics from 'vue-analytics'
import store from './store/store'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';
import Header from "./components/Header";
import connector from "./connector";
import Vue2TouchEvents from "vue2-touch-events";
import VueYoutube from "vue-youtube";
import VueVimeoPlayer from "vue-vimeo-player";

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(Vue2TouchEvents)
Vue.use(VueAnalytics, {
  id: 'UA-152687870-1',
  router
})
Vue.use(VueYoutube)
Vue.use(VueVimeoPlayer)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  connector,
  components: { App, Header},
  template: '<App/>'
})

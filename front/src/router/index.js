import Vue from 'vue'
import Router from 'vue-router'
import Artist from '@/pages/Artist'
import Contact from "../pages/Contact";
import Result from "../pages/Result";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Artist',
      component: Artist
    },
    {
      path: '/contact',
      name: 'Contact',
      component: Contact
    },
    {
      path: '/result',
      name: 'Result',
      component: Result
    }
  ]
})

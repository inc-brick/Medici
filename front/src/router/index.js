import Vue from 'vue'
import Router from 'vue-router'
import Artist from '../presentations/pages/Artist'
import Contact from "../presentations/pages/Contact";
import Result from "../presentations/pages/Result";

Vue.use(Router)

export default new Router({

  routes: [
    {
      path: '/artist/:id',
      name: 'Artist',
      component: Artist
    },
    {
      path: '/contact/:id',
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

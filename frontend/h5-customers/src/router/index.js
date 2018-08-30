import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import AjaxPage from '@/components/AjaxPage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/t-ajax',
      name: 'AjaxPage',
      component: AjaxPage
    }
  ]
})



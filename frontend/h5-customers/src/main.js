// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import http from './utils/http.js'

import ElementUI from 'element-ui';
import MintUI from 'mint-ui'

import 'element-ui/lib/theme-chalk/index.css';
import 'mint-ui/lib/style.css'

Vue.use(ElementUI);
Vue.use(MintUI)

Vue.config.productionTip = false

Vue.prototype.$ajax = http

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})



import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import locale from 'element-ui/lib/locale/lang/ja'
import App from './App.vue'

Vue.use(ElementUI, { locale })

new Vue({
  el: '#app',
  render: h => h(App)
})

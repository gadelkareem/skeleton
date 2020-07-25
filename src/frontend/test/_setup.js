import Vue from 'vue'
import Vuetify from 'vuetify'
import Vuex from 'vuex'
import VueRouter from 'vue-router'

Vue.use(Vuetify)
Vue.use(Vuex)
Vue.use(VueRouter)

require('../src/plugins/mixin')
Vue.use(require('vue-chartist'))
Vue.config.silent = true

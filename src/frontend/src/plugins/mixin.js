import Vue from 'vue'
import initData from '../globals/init-data'
import helpers from '../globals/helpers'
import validator from '../globals/validator'
import paginator from '../globals/paginator'
import permissions from '../globals/permissions'

Vue.prototype.$validator = validator
Vue.prototype.$paginator = paginator

Vue.mixin({
  methods: {
    ...helpers,
    ...permissions,
    initData
  }
})

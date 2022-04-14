import { mount, createLocalVue } from '@vue/test-utils'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import initData from '@@/globals/init-data'
import Vuetify from 'vuetify'
import Page from '../contact'

const localVue = createLocalVue()

describe('invoices.vue', () => {
  let store, vuetify

  beforeEach(() => {
    const actions = {
      'page/title': jest.fn(),
      'loading/start': jest.fn(),
      'loading/finish': jest.fn()
    }
    const getters = {
      'user/user': () => initData().user
    }
    store = new Vuex.Store({
      actions,
      getters,
      state: {
        loading: { status: false }
      }
    })
    vuetify = new Vuetify()
  })

  it('should match snapshot', () => {
    const router = new VueRouter({})
    const w = mount(Page, {
      store,
      router,
      localVue,
      vuetify,
      stubs: ['router-link', 'router-view'],
      created () {
        this.$vuetify.lang = {
          t: () => {}
        }
        this.$vuetify.breakpoint = {
          width: 0,
          height: 0
        }
        this.$vuetify.theme = { dark: false }
      }
    })

    expect(w.html()).toMatchSnapshot()
    w.destroy()
  })
})

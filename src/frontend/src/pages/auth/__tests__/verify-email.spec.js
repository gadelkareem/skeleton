import { mount, createLocalVue } from '@vue/test-utils'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import Page from '../verify-email'
import Vuetify from 'vuetify'

const localVue = createLocalVue()

describe('verify-email.vue', () => {
  let actions, store, vuetify

  beforeEach(() => {
    actions = {
      'page/title': jest.fn(),
      'loading/start': jest.fn(),
      'loading/finish': jest.fn()
    }
    store = new Vuex.Store({
      actions,
      state: {
        loading: { status: false }
      }
    })
    vuetify= new Vuetify()
  })

  it('should match snapshot', () => {
    const router = new VueRouter({})
    const w = mount(Page, {
      store,
      router,
      localVue,
      vuetify,
      stubs: ['router-link', 'router-view']
    })

    expect(w.html()).toMatchSnapshot()
    w.destroy()
  })
})

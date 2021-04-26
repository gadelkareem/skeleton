import { mount } from '@vue/test-utils'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import Vuetify from 'vuetify'
import Page from '../social-callback'

const vuetify = new Vuetify()

describe('social-callback.vue', () => {
  let actions
  let store

  beforeEach(() => {
    actions = {
      'page/title': jest.fn(),
      'loading/start': jest.fn(),
      'loading/finish': jest.fn(),
      'auth/socialCallback': jest.fn()
    }
    store = new Vuex.Store({
      actions,
      state: {
        loading: { status: false }
      }
    })
  })

  it('should match snapshot', () => {
    const router = new VueRouter({})
    const w = mount(Page, {
      store,
      router,
      stubs: ['router-link', 'router-view'],
      vuetify
    })

    expect(w.html()).toMatchSnapshot()

    w.destroy()
  })
})

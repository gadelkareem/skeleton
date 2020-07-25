import { mount } from '@vue/test-utils'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import Page from '../forgot-password'

describe('forgot-password.vue', () => {
  let actions
  let store

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
  })

  it('should match snapshot', () => {
    const router = new VueRouter({})
    const w = mount(Page, {
      store,
      router,
      stubs: ['router-link', 'router-view']
    })

    expect(w.html()).toMatchSnapshot()
    w.destroy()
  })
})

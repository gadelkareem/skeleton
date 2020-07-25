import { mount } from '@vue/test-utils'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import initData from '@@/globals/init-data'
import Page from '../delete-account'

describe('delete-account.vue', () => {
  let store

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

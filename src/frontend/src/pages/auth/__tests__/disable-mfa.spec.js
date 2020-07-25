import { mount } from '@vue/test-utils'
import Vuex from 'vuex'
import Page from '../disable-mfa'

describe('disable-mfa.vue', () => {
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
    const w = mount(Page, {
      store,
      stubs: ['router-link', 'router-view']
    })

    expect(w.html()).toMatchSnapshot()
    w.destroy()
  })
})

import { mount, createLocalVue } from '@vue/test-utils'
import Vuex from 'vuex'
import Page from '../logout'
import Vuetify from 'vuetify'

const localVue = createLocalVue()

describe('logout.vue', () => {
  let actions,  store, vuetify

  beforeEach(() => {
    actions = {
      'page/title': jest.fn(),
      'loading/start': jest.fn(),
      'loading/finish': jest.fn(),
      'auth/logout': jest.fn(),
      'auth/removeSession': jest.fn(),
      'reset': jest.fn()
    }
    store = new Vuex.Store({
      actions,
      state: {
        loading: { status: false }
      }
    })
    vuetify = new Vuetify()
  })

  it('should match snapshot', () => {
    const w = mount(Page, {
      store,
      localVue,
      vuetify,
      stubs: ['router-link', 'router-view']
    })

    expect(w.html()).toMatchSnapshot()
    w.destroy()
  })
})

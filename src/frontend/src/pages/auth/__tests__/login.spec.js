import { mount } from '@vue/test-utils'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import Vuetify from 'vuetify'
import Login from '../login'

const vuetify = new Vuetify()

describe('login.vue', () => {
  let actions
  let store

  beforeEach(() => {
    actions = {
      'page/title': jest.fn(),
      'loading/start': jest.fn(),
      'loading/finish': jest.fn(),
      'auth/login': jest.fn()
    }
    store = new Vuex.Store({
      actions,
      state: {
        loading: { status: false }
      }
    })
  })

  it('should match snapshot', () => {
    const w = mount(Login, {
      store,
      stubs: ['router-link', 'router-view'],
      vuetify
    })

    expect(w.html()).toMatchSnapshot()
    w.destroy()
  })

  it('can process login', async () => {
    const router = new VueRouter({})
    const w = mount(Login, {
      store,
      router,
      stubs: ['router-link', 'router-view'],
      vuetify
    })

    w.find('[data-username]').setValue('user1')
    w.find('[data-password]').setValue('pass1')
    w.find('form').trigger('submit.prevent')
    await w.vm.$nextTick()
    await w.vm.$nextTick()

    expect(w.vm.$route.path).toEqual('/dashboard/home/')
    expect(w.find('.v-alert__content > div').text())
      .toBe('Successfully logged in!')
    w.destroy()
  })
})

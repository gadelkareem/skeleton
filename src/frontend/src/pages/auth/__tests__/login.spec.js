import { mount, createLocalVue } from '@vue/test-utils'
import Vuex from 'vuex'
import VueRouter from 'vue-router'
import Login from '../login'
import Vuetify from 'vuetify'

const localVue = createLocalVue()

describe('login.vue', () => {
  let actions, store, vuetify

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
    vuetify= new Vuetify()
  })

  it('should match snapshot', () => {
    const w = mount(Login, {
      store,
      localVue,
      vuetify,
      stubs: ['router-link', 'router-view']
    })

    expect(w.html()).toMatchSnapshot()
    w.destroy()
  })

  it('can process login', async () => {
    const router = new VueRouter({})
    const w = mount(Login, {
      store,
      router,
      localVue,
      vuetify,
      stubs: ['router-link', 'router-view']
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

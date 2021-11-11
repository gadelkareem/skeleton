import { mount } from '@vue/test-utils'
import Vuetify from 'vuetify'
import Alert from '../Alert'

const vuetify = new Vuetify()

describe('Alert.vue', () => {
  it('should have a custom error and match snapshot', () => {
    const w = mount(Alert, {
      propsData: {
        errors: [{ title: 'test error' }]
      },
      vuetify
    })

    expect(w.html()).toMatchSnapshot()

    const title = w.find('.v-alert__content > div')

    expect(title.text()).toBe('test error')
  })

  it('should have a custom success message', () => {
    const w = mount(Alert, {
      propsData: {
        successTxt: 'test success',
        success: true
      },
      vuetify
    })

    expect(w.find('.v-alert__content > div').text()).toBe('test success')
  })
})

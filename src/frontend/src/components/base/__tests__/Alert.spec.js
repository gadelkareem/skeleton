import { mount } from '@vue/test-utils'
import Alert from '../Alert'

describe('Alert.vue', () => {
  it('should have a custom error and match snapshot', () => {
    const w = mount(Alert, {
      propsData: {
        errors: [{ title: 'test error' }]
      }
    })

    expect(w.html()).toMatchSnapshot()

    const title = w.find('.v-alert__content > div')

    expect(title.text()).toBe('test error')
  })

  it('should have a custom success message', () => {
    const w = mount(Alert, {
      propsData: {
        successText: 'test success',
        success: true
      }
    })

    expect(w.find('.v-alert__content > div').text()).toBe('test success')
  })
})

import api from './api'

export default {
  contact (body) {
    return api.request({
      url: '/common/contact',
      method: 'POST',
      body,
      type: 'contact-request'
    })
  }
}

import api from './api'

export default {
  setupIntent (id) {
    return api.request({
      url: '/customers/' + id + '/setup-intent',
      method: 'GET',
      type: 'setup-intents'
    })
  },
  update (customer) {
    return api.patch('customer', customer)
  },
  listPaymentMethods (id, resetCache) {
    return api.get('customers/' + id + '/payment-methods', { resetCache })
  },
  customerSubscription (id) {
    return api.get('customers/' + id + '/subscription')
  }
}

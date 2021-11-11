import api from './api'

export default {
  deletePaymentMethod (id) {
    return api.delete('payment-method', id)
  },
  addPaymentMethod (paymentMethod) {
    return api.post('payment-method', paymentMethod)
  }
}

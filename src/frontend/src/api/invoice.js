import api from './api'

export default {
  upcomingInvoice (subscription) {
    return api.get('invoices/upcoming', subscription)
  }
}

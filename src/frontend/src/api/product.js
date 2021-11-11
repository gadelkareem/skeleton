import api from './api'

export default {
  list (params) {
    return api.get('product', params)
  }
}

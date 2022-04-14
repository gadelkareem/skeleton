import api from './api'

export default {
  list (params) {
    return api.get('audit-log', params)
  }
}

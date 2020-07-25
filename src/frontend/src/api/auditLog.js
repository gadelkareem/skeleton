import api from './api'

export default {
  list (params) {
    return api.fetch('audit-log', params)
  }
}

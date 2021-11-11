import api from './api'

export default {
  create (sub) {
    return api.post('subscription', sub)
  },
  update (sub) {
    return api.patch('subscription', sub)
  },
  createOrUpdate (sub) {
    return sub.id ? this.update(sub) : this.create(sub)
  },
  cancel (id) {
    return api.delete('subscription', id)
  }
}

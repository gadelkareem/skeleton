import md5 from 'md5'
export default {
  parseError (err) {
    const errors = err && err.response && err.response.data ? err.response.data.errors : []
    if (errors && errors[0] && errors[0].title) {
      return Object.values(errors)
    }
    console.log(err)
    return [{ title: 'Unknown Error, Reload the page and try again.' }]
  },
  initUser () {
    const user = this.$store.getters['user/user']
    if (user && user.id) {
      return Promise.resolve()
    }
    return this.$store.dispatch('user/fetchUser', this.$store.getters['auth/userId'])
      .catch(err => Promise.reject(err))
  },
  firstLetterUpper (s) {
    return !s || !s.length || s[0].toUpperCase()
  },
  hash  (s) {
    return md5(s)
  },
  isEqualObjects (o1, o2) {
    return JSON.stringify(o1) === JSON.stringify(o2)
  }
}

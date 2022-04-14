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
  },
  setCookie (name, value, days) {
    let expires = ''
    if (days) {
      const date = new Date()
      date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000))
      expires = '; expires=' + date.toUTCString()
    }
    document.cookie = name + '=' + (value || '') + expires + '; path=/'
  },
  getCookie (name) {
    const nameEQ = name + '='
    const ca = document.cookie.split(';')
    for (let i = 0; i < ca.length; i++) {
      let c = ca[i]
      while (c.charAt(0) === ' ') { c = c.substring(1, c.length) }
      if (c.indexOf(nameEQ) === 0) { return c.substring(nameEQ.length, c.length) }
    }
    return null
  },
  eraseCookie (name) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;'
  },
  formatTiers (products) {
    const tiers = []
    for (const p of products) {
      const price = p.prices.data[0].recurring.interval === 'month' || !p.prices.data[1] ? p.prices.data[0] : p.prices.data[1]
      const t = {
        title: p.name,
        subheader: p.subheader,
        priceID: price.id,
        price: price.unit_amount / 100,
        prices: p.prices.data,
        description: p.description.split(','),
        buttonText: 'Signup',
        buttonVariant: 'outlined',
        disabled: p.subheader !== 'open'
      }
      tiers.push(t)
    }
    return tiers.sort((a, b) => (a.price > b.price) ? 1 : ((b.price > a.price) ? -1 : 0))
  },
  formatMoney (m) {
    const formatter = new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'EUR',
      minimumFractionDigits: 0

    })

    return formatter.format(m)
  }
}

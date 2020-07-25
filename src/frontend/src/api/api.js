import Kitsu from 'kitsu'
import helpers from '../globals/helpers'
import AuthAPI from './auth'

const api = new Kitsu({
  baseURL: process.env.APIURL,
  headers: {
    'X-Requested-With': 'XMLHttpRequest'
  },
  camelCaseTypes: false,
  axiosOptions: {
    withCredentials: true
  }
})

// Add a response interceptor
const createInterceptor = () => {
  const interceptor = api.interceptors.response.use((response) => {
    return response
  }, (err) => {
    const errors = helpers.parseError(err)
    if (errors[0] && errors[0].status === '401' && errors[0].code === 'INVALID_TOKEN') {
      api.interceptors.response.eject(interceptor)
      return AuthAPI.refreshToken()
        .then((r) => {
          AuthAPI.setToken(r.data.token)
          err.config.headers.Authorization = api.headers.Authorization
          return api.axios.request(err.config)
        })
        .catch((e) => {
          console.log('unexpected refresh token error:', e)
          AuthAPI.removeToken()
          return Promise.reject(err)
        })
        .finally(createInterceptor)
    }
    return Promise.reject(err)
  })
}
createInterceptor()

export default api

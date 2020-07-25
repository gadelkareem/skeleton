import AuthAPI from '../api/auth'

export const state = () => ({
  userId: 0
})

export const getters = {
  isAuthenticated (state) {
    if (process.browser) {
      state.userId = Number(localStorage.getItem('userId'))
    }
    return state.userId !== 0 && !isNaN(state.userId)
  },
  userId (state) {
    return state.userId
  }
}

export const mutations = {
  AUTHENTICATING_SUCCESS (state, payload) {
    AuthAPI.setToken(payload.token)
    state.userId = payload.user_id
    if (process.browser) {
      localStorage.setItem('userId', payload.user_id)
    }
  },
  AUTHENTICATING_ERROR (state) {
    state.userId = 0
    AuthAPI.removeToken()
    if (process.browser) {
      localStorage.clear()
    }
  }
}

export const actions = {
  login ({ commit }, payload) {
    return AuthAPI.login(payload)
      .then(res => commit('AUTHENTICATING_SUCCESS', res.data))
      .catch((err) => {
        commit('AUTHENTICATING_ERROR')
        return Promise.reject(err)
      })
  },
  logout ({ commit }) {
    return AuthAPI.logout()
      .then(() => commit('AUTHENTICATING_ERROR'))
      .catch((err) => {
        return Promise.reject(err)
      })
  },
  socialCallback ({ commit }, payload) {
    return AuthAPI.socialCallback(payload)
      .then(res => commit('AUTHENTICATING_SUCCESS', res.data))
      .catch((err) => {
        commit('AUTHENTICATING_ERROR')
        return Promise.reject(err)
      })
  },
  removeSession ({ commit }) {
    commit('AUTHENTICATING_ERROR')
  },
  refreshToken ({ commit }) {
    return AuthAPI.refreshToken()
      .then(res => commit('AUTHENTICATING_SUCCESS', res.data))
      .catch((err) => {
        commit('AUTHENTICATING_ERROR')
        return Promise.reject(err)
      })
  }
}

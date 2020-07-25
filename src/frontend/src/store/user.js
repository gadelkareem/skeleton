import UserAPI from '../api/user'
import initData from '../globals/init-data'

export const state = () => ({
  user: initData().user
})

export const getters = {
  user (state) {
    return state.user
  }
}

export const mutations = {
  SET_USER (state, user) {
    state.user = user
  }
}

export const actions = {
  fetchUser ({ commit }, id) {
    return UserAPI.get(id)
      .then(res => commit('SET_USER', res.data))
      .catch((err) => {
        commit('SET_USER', null)
        return Promise.reject(err)
      })
  },
  updateUser ({ commit }, user) {
    return UserAPI.update(user)
      .then(res => commit('SET_USER', res.data))
      .catch((err) => {
        return Promise.reject(err)
      })
  },
  removeUser ({ commit }, user) {
    commit('SET_USER', null)
  }
}

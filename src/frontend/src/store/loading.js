
export const state = () => ({
  status: false
})

export const mutations = {
  SET_STATUS (state, status) {
    state.status = status
  }
}

export const actions = {
  start ({ commit }) {
    if (process.browser) {
      window.$nuxt.$root.$loading.start()
    }
    commit('SET_STATUS', true)
  },
  finish ({ commit }) {
    if (process.browser) {
      window.$nuxt.$root.$loading.finish()
    }
    commit('SET_STATUS', false)
  }
}

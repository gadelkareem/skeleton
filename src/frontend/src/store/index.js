export const strict = false

export const mutations = {
  RESET (state) {
    Object.keys(state).forEach((key) => {
      Object.assign(state[key], null)
    })
  }
}

export const actions = {
  reset ({ commit }) {
    commit('RESET')
  }
}

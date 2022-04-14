export const state = () => ({
  title: '',
  hideFooter: false
})

export const getters = {
  title (state) {
    return state.title
  },
  hideFooter (state) {
    return state.hideFooter
  }
}

export const mutations = {
  SET_TITLE (state, title) {
    state.title = title
  },
  SET_HIDE_FOOTER (state, b) {
    state.hideFooter = b
  }
}

export const actions = {
  title ({ commit }, title) {
    commit('SET_TITLE', title)
  },
  hideFooter ({ commit }, b) {
    commit('SET_HIDE_FOOTER', b)
  }
}

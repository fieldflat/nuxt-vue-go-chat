import { SET_USER } from './mutation-types.js'
import { SIGN_UP } from './action-types.js'

export const state = () => ({
  isLoggedIn: false,
  user: null
})

export const getters = {
  isLoggedIn: state => state.isLoggedIn,
  user: state => state.user
}

export const mutations = {
  [SET_USER](state, { user, isLoggedIn }) {
    state.user = user
    state.isLoggedIn = isLoggedIn
  }
}

export const actions = {
  async [SIGN_UP]({ commit, state }, { name, password }) {
    const payload = {
      name: name,
      password: password
    }
    const response = await this.$axios.$post('/signUp', payload)
    commit(SET_USER, { user: response, isLoggedIn: true })
  }
}

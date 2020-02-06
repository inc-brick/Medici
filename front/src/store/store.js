import Vue from 'vue';
import Vuex from 'vuex'

import * as fetchArtistInfo from './modules/fetchArtistInfo'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    fetchArtistInfo
  },
  state: {
    currentSelectedWorkIndex: 0
  },
  mutations: {
    setCurrentSelectedWork(state, newIndex) {
      state.currentSelectedWorkIndex = newIndex;
    }
  },
  actions: {
    setCurrentSelectedWork(context, index) {
      context.commit('setCurrentSelectedWork', index)
    }
  }
})

import Vue from 'vue';
import Vuex from 'vuex'
import Router from "vue-router";
import {ArtistId, ContactInfo} from "../type/requestTypes";
import {ArtistInfo} from "../type/responseTypes";

Vue.use(Vuex)

export default new Router({
  state: {
    artistId: ArtistId,
    contactInfo: ContactInfo,
    artistInfo: ArtistInfo
  },
  mutations: {
    setArtistData (state, data) {
      state.artistInfo = data
    }
  },
  actions: {

  }
})

import Vue from 'vue';
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    fetchArtistInfo: {
      artistId: 1,
      artistNameJPN: "山本　捷平",
      artistNameENG: "Shohei Yamamoto",
      description: "2019年、京都造形芸術大学大学院芸術専攻ペインティング領域卒業。\n" +
        "特殊なローラーを使った技法を用いて描かれる作品群は、図像の反復を以って描かれます。\n" +
        "reiterate (繰り返し、反復)というタイトルに表わされるように、規則的な運動を繰り返す時間を通して変化する主題の意味性やアナログな身体性を感じさせます。",
      artistPriceRange: "10万円~50万円",
      works: [
        {
          id: 1,
          name: "reiterate -master- ",
          createdYear: 2019,
          height: 1167,
          width: 910,
          material: "パネル、麻布、アクリル絵具",
          price: 10,
          url: "./../static/reiterate_master.png"
        },
        {
          id: 2,
          name: "reiterate -tree lines- ",
          createdYear: 2019,
          height: 1300,
          width: 3240,
          material: "パネル、麻布、アクリル絵具",
          price: 15,
          url: "./../static/reiterate_flowers.png"
        },
        {
          id: 3,
          name: "reiterate -flowers- ",
          createdYear: 2019,
          height: 530,
          width: 450,
          material: "パネル、麻布、アクリル絵具",
          price: 18,
          url: "./../static/reiterate_tree_lines.png"
        }
      ],
      events: [],
      medias: [
        // {
        //   id: 1,
        //   isYoutube: true,
        //   videoId: 'nZADYDelP8M'
        // },
        // {
        //   id: 2,
        //   isYoutube: true,
        //   videoId: 'nxz-UjcoJ5k'
        // },
        // {
        //   id: 3,
        //   isYoutube: false,
        //   videoId: '114290321'
        // }
      ]
    },
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

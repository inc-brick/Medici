import Vue from 'vue';
import Vuex from 'vuex'
import Router from "vue-router";

Vue.use(Vuex)

export default new Router({
  state: {
    fetchArtistInfo: {
      artistId: 1,
      artistNameJPN: "山本　捷平",
      artistNameENG: "Shohei Yamamoto",
      description: "2019年、京都造形芸術大学大学院芸術専攻ペインティング領域卒業。\n" +
        "特殊なローラーを使った技法を用いて描かれる作品群は、図像の反復を以って描かれます。\n" +
        "reiterate (繰り返し、反復)というタイトルに表わされるように、規則的な運動を繰り返す時間を通して変化する主題の意味性やアナログな身体性を感じさせます。2019年、京都造形芸術大学大学院芸術専攻ペインティング領域卒業。\n" +
        "特殊なローラーを使った技法を用いて描かれる作品群は、図像の反復を以って描かれます。\n" +
        "reiterate (繰り返し、反復)というタイトルに表わされるように、規則的な運動を繰り返す時間を通して変化する主題の意味性やアナログな身体性を感じさせます。",
      works: [
        {
          id: 1,
          name: "reiterate -master- ",
          createdYear: 2019,
          height: 1167,
          width: 910,
          material: "パネル、麻布、アクリル絵具",
          url: "https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87bh-s.jpg"
        },
        {
          id: 2,
          name: "reiterate -tree lines- ",
          createdYear: 2019,
          height: 1300,
          width: 3240,
          material: "パネル、麻布、アクリル絵具",
          url: "https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87-cgimage-s.jpg"
        },
        {
          id: 3,
          name: "reiterate -flowers- ",
          createdYear: 2019,
          height: 530,
          width: 450,
          material: "パネル、麻布、アクリル絵具",
          url: "https://www.nao.ac.jp/news/sp/20190410-eht/images/m87-eso-s.jpg"
        }
      ],
      events: [
        {
          id: 1,
          url: ""
        }
      ]
    }
  },
  mutations: {

  },
  actions: {

  }
})

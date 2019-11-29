import Vue from 'vue';
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    fetchArtistInfo: {
      artistId: 1,
      artistNameJPN: "山本　捷平",
      artistNameENG: "YAMAMOTO Shohei",
      description: "芸術におけるオリジナルと複製、虚と実の関係は時代を超えて論じられてきたテーマだが、デジタルデータの複製が日常的になった今日、その議論をいかに発展させられるか。急速な情報化によりアイデンティティが曖昧になっている現代において“実体とはなにか“について考える。\n" +
        "現在取り組んでいる『Reiterateシリーズ』は自作のローラーにより絵具をアナログに反復させる絵画作品である。その中でも“Reiterate-line-“は戦後アメリカにおける連邦美術計画の中で生まれた抽象表現主義の作家、ジャクソン・ポロックを元に制作されたシリーズで、絵画を平置きにし絵具を垂らし、自作ローラーによって身体性を伴った反復を施した作品である。ピカソらが生んだキュビズムの極致に挑んだポロックに倣って、抽象の抽象化(解体)による解釈に挑んだ。\n",
      works: [
        {
          id: 1,
          name: "reiterate - Laocoonte - ",
          createdYear: 2019,
          height: 2273,
          width: 9090,
          material: "パネルに麻布、アクリル絵の具、油絵具",
          price: "15 ~ 50",
          url: "./../static/reiterate-laocoonte-.jpg"
        },
        {
          id: 2,
          name: "reiterate - tree lines - ",
          createdYear: 2019,
          height: 1300,
          width: 3240,
          material: "パネルに麻布、アクリル絵の具",
          price: "15 ~ 50",
          url: "./../static/tree-line.jpg"
        },
        {
          id: 3,
          name: "reiterate - flowers - ",
          createdYear: 2019,
          height: 530,
          width: 455,
          material: "パネルに麻布、アクリル絵の具",
          price: "15 ~ 50",
          url: "./../static/flowers2.jpg"
        },
        {
          id: 4,
          name: "reiterate - lines - ",
          createdYear: 2019,
          height: 2273,
          width: 1818,
          material: "パネルに麻布、アクリル絵の具",
          price: "15 ~ 50",
          url: "./../static/reiterate_lines.jpg"
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

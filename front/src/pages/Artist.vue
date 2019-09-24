<template>
  <el-main>
    <h3>山本　捷平</h3>
    <h3>Shohei Yamamoto</h3>
    <WorkView :urls="paths.work" @changing-work="changeWork"></WorkView>
    <h2>About</h2>
    <p>ここのそのアーティストについての内容</p>
    <h2>Upcoming Events</h2>
    <EventView :urls="paths.event"></EventView>
    <h2>Contacts</h2>
    <router-link :to="{ name : 'Contact', params : { id: artistId }}"><i class="el-icon-shopping-cart-2"></i></router-link>
    <h2>Media</h2>
    <MediaView :urls="paths.media"></MediaView>
  </el-main>
</template>

<script>
import WorkView from "../components/WorkView";
import EventView from "../components/EventView";
import MediaView from "../components/MediaView";
import axios from "axios"

export default {
  name: "artist",
  components: {WorkView, EventView, MediaView},
  data () {
      return {
          msg: 'Welcome to Your Vue.js App',
          artistId: 0,
          paths: {
              work: ['https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87bh-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87-cgimage-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/m87-eso-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-bhshadow-s.jpg'],
              event: ['https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87bh-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87-cgimage-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/m87-eso-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-bhshadow-s.jpg'],
              media: ['https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87bh-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-m87-cgimage-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/m87-eso-s.jpg',
                  'https://www.nao.ac.jp/news/sp/20190410-eht/images/20190410-eht-bhshadow-s.jpg']
          },
          currentSelectedWork: '',
          currentSelectedEvent: ''
      }
  },
  methods: {
      changeWork: function (newIndex) {
          this.currentSelectedWork = this.paths.work[newIndex]
          console.log(this.currentSelectedWork)
      }
  },
  created() {
      axios.get("localhost:8080").then(res => console.log("connected to the server")).catch(() => console.log("failed to connect"))
      this.artistId = this.$route.params['id']
  }
}
</script>

<style scoped>
.el-icon-shopping-cart-2 {
  font-size: 4rem;
}
.el-main {
  padding: 60px 20px 20px;
}
</style>

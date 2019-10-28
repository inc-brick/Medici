<template>
  <el-main>
    <h3>山本　捷平</h3>
    <h3>Shohei Yamamoto</h3>
    <WorkView :works="fetchArtistInfo['works']" @changing-work="changeWork"></WorkView>
    <h2>About</h2>
    <p class="style">{{this.fetchArtistInfo['description']}}</p>
    <div v-if="fetchArtistInfo['events'].length !== 0">
      <h2>Upcoming Events</h2>
      <EventView :events="fetchArtistInfo['events']"></EventView>
    </div>
    <h2>Contacts</h2>
    <router-link :to="{ name : 'Contact', params : { id: artistId }}"><i class="el-icon-shopping-cart-2"></i></router-link>
    <div v-if="fetchArtistInfo['medias'].length !== 0">
      <h2>Media</h2>
      <MediaView :medias="fetchArtistInfo['medias']"></MediaView>
    </div>
  </el-main>
</template>

<script>
import WorkView from "../components/WorkView";
import EventView from "../components/EventView";
import MediaView from "../components/MediaView";
import connector from "../connector";

export default {
  name: "artist",
  components: {WorkView, EventView, MediaView},
  data () {
      return {
          msg: 'Welcome to Your Vue.js App',
          artistId: 0,
          fetchArtistInfo: {},
          currentSelectedWork: '',
          currentSelectedEvent: ''
      }
  },
  methods: {
      changeWork: function (newIndex) {
          this.currentSelectedWork = this.fetchArtistInfo.works[newIndex]['url']
          console.log(this.currentSelectedWork)
      }
  },
  created() {
      connector.getArtistData()
      console.log("connector.getArtistData is called")
      console.log(this.fetchArtistInfo)
      this.fetchArtistInfo = this.$store.state.fetchArtistInfo
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
.style {
  font-size: 11px;
}
</style>

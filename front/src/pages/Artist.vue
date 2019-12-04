<template>
  <el-main>
    <h3 class="artist_name">山本　捷平<br>YAMAMOTO Shohei</h3>
    <WorkView :works="fetchArtistInfo['works']" :artist-price-range="fetchArtistInfo['artistPriceRange']" @changing-work="changeWork"></WorkView>
    <h4>About</h4>
    <p class="style">{{this.fetchArtistInfo['description']}}</p>
    <div v-if="fetchArtistInfo['events'].length !== 0">
      <h4>Upcoming Events</h4>
      <EventView :events="fetchArtistInfo['events']"></EventView>
    </div>
    <h4>Contact</h4>
    <router-link :to="{ name : 'Contact', params : { id: fetchArtistInfo.artistId }}">
      <el-image :src="'./../static/shopping_cart.png'"></el-image>
    </router-link>
    <div v-if="fetchArtistInfo['medias'].length !== 0">
      <h4>Media</h4>
      <MediaView :medias="fetchArtistInfo['medias']"></MediaView>
    </div>
    <p>&nbsp;</p>
    <div style="line-height: 0.2em">
      <p class="style">Powered by brick Inc.</p>
      <p class="style"><a href= "https://inc-brick.com/" >https://inc-brick.com/</a></p>
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
          fetchArtistInfo: {},
          currentSelectedWork: '',
      }
  },
  methods: {
      changeWork: function (newIndex) {
          this.$store.dispatch('setCurrentSelectedWork', newIndex)
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
h4 {
  margin-bottom: 5px;
}
.artist_name {
  margin-top: 10px;
  margin-bottom: 30px;
}
.el-icon-shopping-cart-2 {
  font-size: 4rem;
}
.el-main {
  padding: 60px 20px 20px;
}
.style {
  font-size: 11px;
}
.el-image {
  height: 40px;
  width: 40px;
}
</style>

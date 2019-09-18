<template>
  <div class="artist">
    <h3>山本　捷平</h3>
    <h3>Shohei Yamamoto</h3>
    <WorkView :urls="urls"></WorkView>
    <h2>About</h2>
    <p>ここのそのアーティストについての内容</p>
    <h2>Upcoming Events</h2>
    <EventView :urls="urls"></EventView>
    <h2>Contacts</h2>
    <i class="el-icon-shopping-cart-2"></i>
    <h2>Media</h2>
    <MediaView :urls="urls"></MediaView>
  </div>
</template>

<script>
import WorkView from "../components/WorkView";
import EventView from "../components/EventView";
import MediaView from "../components/MediaView";

export default {
  name: "artist",
  components: {WorkView, EventView, MediaView},
  data () {
      return {
          msg: 'Welcome to Your Vue.js App',
          urls: ['https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg']
      }
  }
}
</script>

<style scoped>

</style>

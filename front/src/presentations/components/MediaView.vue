<template>
  <div class="block">
    <el-carousel :trigger="carouselConfig.trigger"
                 :indicator-position="carouselConfig.indicatorPosition"
                 :arrow="carouselConfig.arrow"
                 :autoplay="carouselConfig.autoplay"
                 ref="carousel"
                 v-touch:swipe.left="swipeLeft"
                 v-touch:swipe.right="swipeRight"
                 @change="changeWork">
      <el-carousel-item v-for="(media,index) in medias" :key="index">
        <div v-if="media.isYoutube" class="padding">
          <youtube :video-id="media.videoId" ref="youtube" @playing="playing"></youtube>
        </div>
        <div v-else>
          <vimeo-player ref="player" :video-id="media.videoId" @ready="onReady"></vimeo-player>
        </div>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
    name: 'media-view',
    props: {
        medias: Array
    },
    data() {
        return {
            carouselConfig: {
                trigger: 'click',
                indicatorPosition: 'outside',
                arrow: 'always',
                autoplay: false,
            }
        }
    },
    methods: {
        playing() {
            console.log('\o/ we are watching!!!')
        },
        onReady() {
            this.playerReady = true
        },
        changeWork: function (newIndex, oldIndex) {
            console.log("changing-work: to: " + newIndex + " from: " + oldIndex)
            this.$emit("changing-work", newIndex)
        },
        swipeRight: function () {
            console.log("swipe to right")
            this.$refs.carousel.prev()
        },
        swipeLeft: function () {
            console.log("swipe to left")
            this.$refs.carousel.next()
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .padding {
    padding: 0 auto 0;
  }
</style>

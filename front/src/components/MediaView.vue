<template>
  <div class="block">
    <el-carousel trigger="click" height="400px" indicator-position="outside" arrow="always">
      <el-carousel-item v-for="videoId in videoIds" :key="videoId.id">
        <div v-if="videoId.isYoutube" class="padding">
          <youtube :video-id="videoId.id" ref="youtube" @playing="playing"></youtube>
        </div>
        <div v-else>
          <vimeo-player ref="player" :video-id="videoId.id" @ready="onReady" :player-height="height"></vimeo-player>
        </div>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
    name: 'media-view',
    props: {
        urls: Array
    },
    data() {
        return {
            fit: 'fill',
            videoIds: [
                {
                    isYoutube: true,
                    id: 'nZADYDelP8M'
                },
                {
                    isYoutube: true,
                    id: 'nxz-UjcoJ5k'
                },
                {
                    isYoutube: false,
                    id: '114290321'
                }
            ],
        }
    },
    methods: {
        playing() {
            console.log('\o/ we are watching!!!')
        },
        onReady() {
            this.playerReady = true
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

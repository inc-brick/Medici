<template>
  <div class="block">
    <el-carousel :trigger="carouselConfig.trigger"
                 :indicator-position="carouselConfig.indicatorPosition"
                 :arrow="carouselConfig.arrow"
                 :autoplay="carouselConfig.autoplay"
                 ref="carousel"
                 v-touch:swipe.left="swipeLeft"
                 v-touch:swipe.right="swipeRight"
                 @change="changeWork"
                 :height="carouselConfig.height"
                 :initial-index="initialIndex"
                 >
        <el-carousel-item v-for="(event,index) in events" :key="index">
        　<a :href="event.eventUrl" target="_blank">
            <el-image :src="event.photoUrl"
                      :fit="carouselConfig.fit"
                      class="style"
            >
              <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </a>
          <div style="margin-top: 5px">
            <div class="image_description" style="text-align: center;font-weight: bold;font-size: 14px">
              {{event.name}}
            </div>
            <div class="image_description" style="text-align: center;font-size: 7px">
              {{event.schedule}}
            </div>
            <div class="image_description" style="text-align: center;font-size: 7px">
              {{event.place}}
            </div>
        </div>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
    name: 'event-view',
    props: {
        initialIndex: Number,
        events: Array
    },
    data() {
        return {
            carouselConfig: {
                trigger: 'click',
                indicatorPosition: 'outside',
                arrow: 'always',
                fit: 'contain',
                autoplay: false,
                height: '200px'
            }
        }
    },
    methods: {
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
.style {
  height: 150px;
  width: 100%;
}
.el-carousel__arrow {
  width: 24px;
  height: 24px;
}
</style>

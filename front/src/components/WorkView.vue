<template>
  <div class="block">
    <el-carousel :trigger="carouselConfig.trigger" :indicator-position="carouselConfig.indicatorPosition" :arrow="carouselConfig.arrow" :autoplay="carouselConfig.autoplay" @change="changeWork">
      <el-carousel-item v-for="url in urls" :key="url">
        <el-image :src="url" :fit="carouselConfig.fit" class="style">
          <div slot="error" class="image-slot">
            <i class="el-icon-picture-outline"></i>
          </div>
        </el-image>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
    name: 'work-view',
    props: {
        urls: Array
    },
    data() {
        return {
            carouselConfig: {
                trigger: 'click',
                indicatorPosition: 'outside',
                arrow: 'always',
                fit: 'contain',
                autoplay: false,
            }
        }
    },
    methods: {
        changeWork: function (newIndex, oldIndex) {
            console.log("changing-work: to: " + newIndex + " from: " + oldIndex)
            this.$emit("changing-work", newIndex)
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.style {
  height: 300px;
  width: 100%;
}
</style>

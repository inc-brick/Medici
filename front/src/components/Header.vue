<template>
  <el-header>
    <div class="header_img_block">
      <img src="../../static/brick-logo.png" class="header_img">
    </div>
  </el-header>
</template>
<script>
export default {
  name: 'Header',
  props: {
      artistId: Number
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-header {
  background-color: #ffffff;
  position: fixed;
  z-index:9999;
  width: 100%;
  height: 60px;
}
.header_img_block {
  line-height: 60px;
}
.header_img {
  height: 30px;
  vertical-align: middle;
}
</style>

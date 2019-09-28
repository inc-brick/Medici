<template>
  <el-header>
    <img src="./../assets/brick-logo.png" style="height: 30px; width: 60px">
  </el-header>
</template>
<script>
export default {
  name: 'Header',
  data() {
      return {
          url: "/Users/kentatakimur/dev/Medici/front/src/assets/brick-logo.png",
          fit: "contain",
          height: "30px"
      }
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
  line-height: 60px;
}
</style>

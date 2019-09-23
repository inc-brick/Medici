<template>
  <header>
  　<div class="header">
      <img src="./../assets/brick-logo.png" style="height: 30px; width: 60px">
<!--  <el-header :height=height class="header_title">-->
<!--    <el-image-->
<!--      style="width: 100px; height: 100px"-->
<!--      :src="url"-->
<!--      :fit="fit"></el-image>-->
<!--  </el-header>-->
    </div>
  </header>
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
.header {
  display: flex;
}

.header_title {
  font-size: 18px;
  font-weight: bold;
}
</style>

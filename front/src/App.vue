<template>
  <el-container class="is-vertical">
    <Header></Header>
    <el-container>
      <router-view/>
    </el-container>
  </el-container>
</template>

<script>
import Header from "./components/Header";

export default {
    name: 'App',
    components: {Header}
}
</script>

<style>
#app {
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
}
body {
  margin: 0;
}
</style>

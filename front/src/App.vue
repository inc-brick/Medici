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
body {
  margin: 0;
  font-family: "Times New Roman",serif;
}
</style>

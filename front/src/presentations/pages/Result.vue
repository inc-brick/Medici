<template>
  <el-main>
    <p>担当者からのご連絡お待ちください。</p>
    <div class="back-to-artist-button">
      <router-link to="/artist/1"><el-button type="Info">Back to Artist</el-button></router-link>
    </div>
  </el-main>
</template>

<script>
export default {
  name: "result"
}
</script>

<style scoped>
.el-main {
  padding: 60px 20px 20px;
}
.result{
  margin-top: 12rem;
  font-weight: bold;
}
.back-to-artist-button {
  text-align: center;
  margin-top: 32px;
}
</style>

<template>
  <el-main>
    <p>お問い合わせありがとうございます！</p>
    <p>担当者からのご連絡お待ちください。</p>
    <div class="back-to-artist-button">
      <router-link to="/"><el-button type="Info">Back to Artist</el-button></router-link>
    </div>
  </el-main>
</template>

<script>
export default {
  name: "result"
}
</script>

<style scoped>
.result{
  margin-top: 12rem;
  font-weight: bold;
}
.back-to-artist-button {
  text-align: center;
  margin-top: 32px;
}
</style>

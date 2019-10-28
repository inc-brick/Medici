<template>
  <el-main>
    <h3>Contact</h3>
    <p>3 step で完了します</p>
    <h4>1. ご興味のある作品を選択してください</h4>
    <WorkView :works="works" :initial-index="initialImageIndex"></WorkView>
    <h4>2. ご希望の連絡方法を選択して下さい。後日、担当者からご連絡差し上げます。</h4>
    <el-row style="height: 100px">
      <el-col :span="6" :offset="6"><i class="el-icon-message box-shadow" :class="{messageSelected: messageSelected}" @click="selectMessage"></i></el-col>
      <el-col :span="6"><i class="el-icon-phone-outline box-shadow" :class="{phoneSelected: phoneSelected}" @click="selectPhone"></i></el-col>
    </el-row>
    <h4>3. ご連絡先情報を入力してください。</h4>
    <el-form>
      <el-form-item><el-input placeholder="Name" class="medium" v-model="name"></el-input></el-form-item>
      <el-form-item><el-input placeholder="Email" class="medium" v-model="email"></el-input></el-form-item>
      <el-form-item><el-input placeholder="Phone" class="medium" v-model="phone"></el-input></el-form-item>
      <el-form-item class="center">
        <el-button type="primary" @click="submit">Submit</el-button>
      </el-form-item>
    </el-form>
  </el-main>
</template>

<script>
import WorkView from "../components/WorkView";
import EventView from "../components/EventView";
import MediaView from "../components/MediaView";

export default {
  name: "contact",
  components: {WorkView, EventView, MediaView},
  data () {
      return {
          works: [],
          initialImageIndex: '',
          msg: 'Welcome to Your Vue.js App',
          urls: ['https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'],
          name: '',
          email: '',
          phone: '',
          messageSelected: false,
          phoneSelected: false
      }
  },
  methods: {
      selectMessage: function () {
          if (!this.messageSelected) {
              this.messageSelected = true
              this.phoneSelected = false
          } else {
              this.messageSelected = false
          }
      },
      selectPhone: function () {
          if (!this.phoneSelected) {
              this.phoneSelected = true
              this.messageSelected = false
          } else {
              this.phoneSelected = false
          }
      },
      submit: function () {
          this.$router.push('/result')
      }
  },
  created() {
      this.works = this.$store.state.fetchArtistInfo.works
      this.initialImageIndex = this.$store.state.currentSelectedWorkIndex
  }
}
</script>

<style scoped>
.box-shadow {
  box-shadow: 0 10px 10px rgba(0, 0, 0, .12), 0 0 12px rgba(0, 0, 0, .04)
}
.el-icon-message {
  font-size: 4rem;
}
.el-icon-phone-outline {
  font-size: 4rem;
}
.center {
  text-align: center;
}
.i {
  border-radius: 4px;
}
.phoneSelected {
  background-color: #8c939d;
}
.messageSelected {
  background-color: #8c939d;
}
</style>

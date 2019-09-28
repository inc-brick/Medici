<template>
  <el-main>
    <h3>Contact</h3>
    <p>3 step で完了します</p>
    <h4>1. ご興味のある作品を選択してください</h4>
    <WorkView :urls="urls"></WorkView>
    <h4>2. ご希望の連絡方法を選択して下さい。後日、担当者からご連絡差し上げます。</h4>
    <el-row>
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
import connector from "../connector/connector";

export default {
  name: "contact",
  components: {WorkView, EventView, MediaView},
  data () {
      return {
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
          }
      },
      selectPhone: function () {
          if (!this.phoneSelected) {
              this.phoneSelected = true
              this.messageSelected = false
          }
      },
      submit: function () {
          let ok = connector.postContactsInfo()
          if(ok) {
              this.$router.push('/result')
          } else {
              console.log("想定外のエラーが発生しました")
          }
      }
  }
}
</script>

<style scoped>
.box-shadow {
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)
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
.phoneSelected {
  background-color: #99a9bf;
}
.messageSelected {
  background-color: #99a9bf;
}
</style>

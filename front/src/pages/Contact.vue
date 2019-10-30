<template>
  <el-main>
    <h3>Contact</h3>
    <p>3 step で完了します</p>
    <h4>1. ご興味のある作品を選択してください</h4>
    <WorkView :works="works" :initial-index="initialImageIndex" :is-artist-page="false"></WorkView>
    <h4>2. ご希望の連絡方法を選択して下さい。後日、担当者からご連絡差し上げます。</h4>
    <el-row style="height: 100px">
      <el-col :span="6" :offset="6"><i class="el-icon-message box-shadow" :class="{messageSelected: messageSelected}" @click="selectMessage"></i></el-col>
      <el-col :span="6"><i class="el-icon-phone-outline box-shadow" :class="{phoneSelected: phoneSelected}" @click="selectPhone"></i></el-col>
    </el-row>
    <h4>3. ご連絡先情報を入力してください。</h4>
    <el-form v-if="!isSend"
             :model="formVal"
             ref="form"
             :action="'https://docs.google.com/forms/u/1/d/e/1FAIpQLSc10-M1uZi5jD2jmyK_ICom4KipjEWXv6O6xHqTQq6vyvO_hg/formResponse'"
             target="hidden_iframe"
             @submit.prevent="submitting('form')"
    >
      <el-form-item
        prop="name"
        :rules="{required: true, message: 'please input your name', trigger: ['blur', 'change']}"
      >
        <el-input placeholder="Name" class="medium" :name="formName.name" v-model="formVal.name"></el-input>
      </el-form-item>
      <el-form-item
        v-if="messageSelected || (!messageSelected && !phoneSelected)"
        prop="email"
        :rules="[
        {required: true, message: 'please input your email address', trigger: ['blur', 'change']},
        {type: 'email', message: 'Please input correct email address', trigger: ['blur', 'change']}
        ]"
      >
        <el-input placeholder="Email" class="medium" :name="formName.email" v-model="formVal.email"></el-input>
      </el-form-item>
      <el-form-item
        v-if="phoneSelected || (!phoneSelected && !messageSelected)"
        prop="phone"
        :rules="[
        {required: true, message: 'please input your phone number', trigger: ['blur']}
        ]"
      >
        <el-input placeholder="Phone" class="medium" :name="formName.phone" v-model="formVal.phone"></el-input>
      </el-form-item>
      <el-form-item class="center">
        <el-input :type="'submit'">Submit</el-input>
      </el-form-item>
    </el-form>
  </el-main>
</template>

<script>
import WorkView from "../components/WorkView";
import EventView from "../components/EventView";
import MediaView from "../components/MediaView";
import connector from "../connector";

export default {
  name: "contact",
  components: {WorkView, EventView, MediaView},
  data () {
      return {
          works: [],
          initialImageIndex: 0,
          msg: 'Welcome to Your Vue.js App',
          urls: ['https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'],
          formName: {
              name: "entry.1239827993",
              email: "entry.1682303839",
              phone: "entry.1390168152"
          },
          formVal: {
              name: '',
              email: '',
              phone: ''
          },
          messageSelected: false,
          phoneSelected: false,
          inputType: {
              email: "email",
              phone: "tel"
          },
          isSend: false
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
      // submit: function (formName) {
      //     this.$refs[formName].validate((valid) => {
      //         if (valid) {
      //             connector.postGform(this.formVal)
      //                 .then(() => {
      //                     this.$router.push("/result")
      //                 })
      //         }
      //     });
      // },
      submitting: function (formName) {
          this.$refs[formName].validate((valid) => {
              if (valid) {
                  document.form.submit()
                  this.isSend = true
              }
          });
      }
  },
  created() {
      this.works = this.$store.state.fetchArtistInfo.works
      this.initialImageIndex = this.$store.state.currentSelectedWorkIndex
  },
  mounted: function() {
      let iframe = document.createElement("iframe");
      iframe.setAttribute('name','hidden_iframe');
      iframe.setAttribute('style','display: none');
      document.body.appendChild(iframe);
      if (this.isSend) {
          this.$router.push("/result")
      }
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

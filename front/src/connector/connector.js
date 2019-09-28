import axios from "axios";

const apiConfig = {
  headers: {
    'Content-Type': 'application/json'
  }
}

const API_BASE_URL = "http://localhost:8000"
axios.create({
    baseURL: API_BASE_URL
  }
)

const connector = {
  // TODO(takimura): このままだとCORSのエラーになる。サーバー側の修正が必要
  async fetchArtistData(id) {
    await axios.get("/artist/" + id,
      apiConfig
    ).then(res => {
      if (res.status === 200) {
        this.$store.commit("setArtistInfo", res.data)
        return true
      }
      return false
    }).catch(err => {
      console.log(err)
    })
  },
  async postContactsInfo() {
    await axios.post("/contact",
      {
        contactInfo: this.$store.state.contactInfo
      },
      apiConfig
    ).then(res => {
      if (res.status === 200)  {
        return true
      }
      return false
    }).catch(err => {
      console.log(err)
    })
  }
}

export default connector

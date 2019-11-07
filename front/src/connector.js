import axios from "axios";

const ENV = {
  MOCK: "mock",
  DEVELOPMENT: "development",
  PRODUCTION: "production"
}

if (process.env.NODE_ENV === ENV.MOCK) {

}


const connector = {
  // TODO(takimura): このままだとCORSのエラーになる。サーバー側の修正が必要
  async getArtistData() {
    // await axios.get("http://localhost:8080", {})
  },
  async postArtistWork() {
    await axios.post("http://localhost:8080", {})
  },
  async postContactsInfo() {
    await axios.post("http://localhost:8080", {})
  },
  async postGform(data) {
    let res = await axios.post(
      "http://app-brick.com:8000/post/contact",
      data,
      {
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Headers": "*",
          "Access-Control-Allow-Methods": "*"
        }
      })
    console.log(res)
    return res
  }
}

export default connector

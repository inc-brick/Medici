import axios from "axios";

const ENV = {
  MOCK: "mock",
  DEVELOPMENT: "development",
  PRODUCTION: "production"
}

const BASE_URL = process.env.NODE_ENV === ENV.PRODUCTION ? "https://app-brick.com" : "http://localhost:8000"

export const connector = {
  // TODO(takimura): このままだとCORSのエラーになる。サーバー側の修正が必要
  async getArtistData() {
    if (process.env.NODE_ENV === ENV.DEVELOPMENT) {
      console.log("This is displayed only when you connected to the dev server")
    }
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
      "https://app-brick.com/post/contact",
      data,
      {
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*"
        }
      })
    console.log(res)
    return res
  }
}

export default axios.create({
  baseURL: BASE_URL
})

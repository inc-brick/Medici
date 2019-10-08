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
  }
}

export default connector

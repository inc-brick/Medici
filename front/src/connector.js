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
    await axios.post(
      "https://docs.google.com/forms/u/1/d/e/1FAIpQLSc10-M1uZi5jD2jmyK_ICom4KipjEWXv6O6xHqTQq6vyvO_hg/formResponse",
      {
        "entry.1239827993": data['name'],
        "entry.1682303839": data['email'],
        "entry.1390168152": data['phone']
    },
      {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          "Access-Control-Allow-Origin": ""
        }
      })
  }
}

export default connector

import connector from "../connector";

const resource = "/post"
export default {
  postGform(data) {
    return connector.post(
      resource + "/contact",
      data,
      {
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*"
        }
      })
  }
}

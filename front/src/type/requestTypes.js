// リクエストに用いるパラメータは全てここに定義する
export const ArtistId = function (artistId) {
  this.artistId = artistId
}

export const ContactInfo = function (artistId, workId, method, name, email, phone) {
  this.artistId = artistId
  this.workId = workId
  this.method = method
  this.name = name
  this.email = email
  this.phone = phone
}

// レスポンスに用いるパラメータは全てここに定義する
export const ArtistInfo = function (artistId, artistNameJPN, artistNameENG, description, works, events) {
  this.artistId = artistId
  this.artistNameJPN = artistNameJPN
  this.artistNameENG = artistNameENG
  this.description = description
  this.works = works
  this.events = events
}

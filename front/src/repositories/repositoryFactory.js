import ArtistRepository from './artist/artistRepository'
import ContactRepository from './contact/contactRepository'

const repositories = {
  artist: ArtistRepository,
  contact: ContactRepository
}

export default {
  get: repositoryName => repositories[repositoryName]
}

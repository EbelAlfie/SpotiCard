import { getSpotifyPlayerCard } from "../card/SpotifyCard.js"
import { TokenUseCase } from "../domain/TokenUseCase.js"
import { TrackUseCase } from "../domain/TrackUseCase.js"
import { TokenRepository } from "../service/TokenRepository.js"
import { TrackRepository } from "../service/TrackRepository.js"

const tokenHandler = async (response) => {
    const tokenRepository = new TokenRepository()
    const tokenUseCase = new TokenUseCase(tokenRepository)

    const tokenObj = await tokenUseCase.fetchAccessToken()

    if (tokenObj instanceof Error) {
        response.status(500)
        response.send(tokenObj.message)
        return 
    }

    const clientToken = await tokenUseCase.fetchClientToken({
        clientId: tokenObj?.clientId
    })
    if (clientToken instanceof Error) {
        response.status(500)
        response.send(clientToken.message)
        return 
    }
    
    return {
        ...clientToken,
        ...tokenObj
    }
}

export const getSpotifyCard = async (request, response) => {
    const tokens = await tokenHandler(response)
    if (!tokens) return 

    const trackRepository = new TrackRepository(tokens)
    const trackUseCase = new TrackUseCase(trackRepository)

    const trackResult = await trackUseCase.getTrackById({
        trackId: "0MJ5wsGpqu0gTJkx53ewxc"
    })

    if (trackResult instanceof Error) {
        response.status(500)
        response.send(trackResult.message)
        return
    }

    const image = trackResult.images?.length > 0 ? trackResult?.images[0]?.url : ""

    const spotifyCard = getSpotifyPlayerCard(
        {
            imageUrl: image, 
            songTitle: trackResult.name, 
            artists: trackResult.artists?.map(item => item.name).join(", "),
            audioUrl: trackResult.previewUrl
        }
    )

    response.status(200)
    response.send(spotifyCard)
}
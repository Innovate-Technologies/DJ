import { getSeparationRules } from "../itframe/api.js"

export default async (list) => {
    const rules = await getSeparationRules()
    if (!rules || !rules.enabled) {
        return list
    }
    const listOfArtist = []
    const listOfSongIDs = []
    const listOfSummedTime = []
    if (rules.separationType === "artist") {
        for (let song of list) {
            listOfArtist.push(song.artist)
        }
    }

    if (rules.separationType === "song") {
        for (let song of list) {
            listOfSongIDs.push(song._id)
        }
    }

    for (let song of list) {
        listOfSummedTime.push((listOfArtist[listOfArtist.length - 1] || 0) + song.duration)
    }

    for (let song of list) {
        let indexes
        if (rules.separationType === "artist") {
            indexes = getAllIndexFor(listOfArtist, song.artist)
        }
        if (rules.separationType === "song") {
            indexes = getAllIndexFor(listOfSongIDs, song._id)
        }
        if (indexes.length > 1) {
            for (let id in indexes) {
                if (indexes.hasOwnProperty(id) && id !== 0) {
                    let timecodePrevious = listOfSummedTime[indexes[id - 1]]
                    let timecodeThis = listOfSummedTime[indexes[id]]
                    if (timecodeThis - timecodePrevious < rules.interval) {
                        // do SOMETHING!!
                    }
                }
            }
        }
    }

}

const getAllIndexFor = (array, value) => {
    const returnIndex = []
    for (let id in array) {
        if (array.hasOwnProperty(id) && array[id] === value) {
            returnIndex.push(id)
        }
    }
    return returnIndex
}

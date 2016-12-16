import _ from "underscore";
import { getSeparationRules } from "../itframe/api.js"

export default async (list) => {
    const rules = await getSeparationRules()
    if (!rules || !rules.enabled) {
        return list
    }
    const listOfArtist = []
    const listOfSongIDs = []
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
        let indexes
        if (rules.separationType === "artist") {
            indexes = getAllIndexFor(listOfArtist, song.artist)
        }
        if (rules.separationType === "song") {
            indexes = getAllIndexFor(listOfSongIDs, song._id)
        }
        if (indexes.length > 1) {
            let previousIndex;
            for (let index of indexes) {
                if (!previousIndex) {
                    const duration = getTimeBetween(list, previousIndex, index)
                    if (duration >= rules.interval) {
                        song = null
                    }
                }
                previousIndex = index
            }
        }
    }
    const list2 = _.without(list, null)
    if (list2.length > 5) {
        return list2
    }
    return list // too many songs deleted, no separation possible!
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

const getTimeBetween = (list, index1, index2) => {
    let duration = 0
    for (let i = index1 + 1; i < index2; i++) {
        duration += list[i].duration || 0
    }
    return duration
}

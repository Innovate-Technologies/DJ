import { getIntervals } from "../itframe/api.js"
import _ from "underscore"

export default async (songsList) => {
    const intervals = await getIntervals()
    const currentIntervals = selectCurrentIntervals(intervals)

    if (currentIntervals.length > 0) {
        for (let interval of currentIntervals) {
            songsList = insertInterval(songsList, interval)
        }
    }

    return songsList
}
const selectCurrentIntervals = (intervals) => {
    const now = new Date()
    const currentIntervals = []
    for (let interval of intervals) {
        if (new Date(interval.start) < now && (new Date(interval.end) > now || interval.forever)) {
            debug("Use interval " + interval.name)
            currentIntervals.push(interval)
        }
    }
    return currentIntervals
}


const insertInterval = function (playlist, {songs, intervalType, songsAtOnce, every, intervalMode}) {
    let count = 0
    let orderCount = 0
    const newPlaylist = []

    if (songs.length === 0) {
        return playlist
    }

    for (let song of playlist) {
        newPlaylist.push(song)
        if (intervalMode === "songs") {
            count++
        } else if (intervalMode === "seconds") {
            count += song.duration
        }

        if (count >= every) {
            if (intervalType === "all") {
                for (let songInInterval of songs) {
                    songInInterval.ignoreSeperation = true;
                    newPlaylist.push(songInInterval)
                }
            } else {
                for (let i = 0; i < songsAtOnce; i++) {
                    let songInInterval
                    if (intervalType === "random") {
                        songInInterval = _.shuffle(songs)[0]
                        songInInterval.ignoreSeperation = true;
                    }
                    if (intervalType === "order") {
                        songInInterval = songs[orderCount]
                        orderCount++
                        if (orderCount >= songs.length) {
                            orderCount = 0
                        }
                    }
                    newPlaylist.push(songInInterval)
                }
            }
            count = 0
        }
    }
    return newPlaylist
}

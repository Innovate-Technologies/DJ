import { getIntervals, getSongInfo } from "../itframe/api.js"
import _ from "underscore"

export default async (songsList) => {
    const intervals = await getIntervals()
    const currentIntervals = selectCurrentIntervals(intervals)
    debug(currentIntervals)
    if (currentIntervals.length > 0) {
        for (let interval of currentIntervals) {
            insertInterval(songsList, interval)
        }
    }
    return songsList
}
var selectCurrentIntervals = (intervals) => {
    var currentIntervals = []
    for (var interval of intervals) {
        if (interval.forever) {
            currentIntervals.push(interval)
        } else {
            var now = new Date()
            if (interval.start < now && (interval.end > now || interval.forever)) {
                debug("Use interval " + interval.name)
                currentIntervals.push(interval)
            }
        }
    }
    return currentIntervals
}


const insertInterval = function (playlist, {songs, intervalType, songsAtOnce, every, intervalMode}) {
    let count = 0
    let orderCount = 0
    let newPlaylist = []

    if (songs.length === 0) {
        return playlist
    }

    for (var song of playlist) {
        newPlaylist.push(song)
        if (intervalMode === "songs") {
            count++
        } else if (intervalMode === "seconds") {
            count += song.duration
        }

        if (count >= every) {
            for (var i = 0; i < songsAtOnce; i++) {
                var intervalSong
                if (intervalType === "random") {
                    intervalSong = _.shuffle(songs)[0]
                    intervalSong.ignoreSeperation = true;
                    newPlaylist.push(intervalSong)
                } else if (intervalType === "order") {
                    intervalSong = songs[orderCount]
                    intervalSong.ignoreSeperation = true;
                    newPlaylist.push(intervalSong)
                    orderCount++
                    if (orderCount >= songs.length) {
                        orderCount = 0
                    }
                } else if (intervalType === "all") {
                    for (var songInterval of songs) {
                        songInterval.ignoreSeperation = true;
                        newPlaylist.push(songInterval)
                    }
                }
            }
            count = 0
        }
    }
    return newPlaylist
}

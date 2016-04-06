const itframe = requireFromRoot("components/itframe/api.js")
const wait = require("wait.for")
const _ = require("underscore")
const queue = require("queue-async")

var songsForID = {}

module.exports = (songsList, callback) => {
    wait.launchFiber(() => {
        var intervals = wait.for(itframe.getIntervals)
        var currentIntervals = selectCurrentIntervals(intervals)
        debug(currentIntervals)
        if (currentIntervals.length > 0) {
            for (var interval of currentIntervals) {
                var q = queue(10);
                for (var songID of interval.songs) {
                    q.defer(addSongInfo, songID)
                }
                wait.for(q.awaitAll)
                songsList = insertInterval(songsList, interval)
            }
            return callback(null, songsList);
        }
        callback(null, songsList);
    })
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


var insertInterval = function (playlist, {songs, intervalType, songsAtOnce, every, intervalMode}) {
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
                    intervalSong = songsForID[_.shuffle(songs)[0]]
                    intervalSong.ignoreSeperation = true;
                    newPlaylist.push(intervalSong)
                } else if (intervalType === "order") {
                    intervalSong = songsForID[songs[orderCount]]
                    intervalSong.ignoreSeperation = true;
                    newPlaylist.push(intervalSong)
                    orderCount++
                    if (orderCount >= songs.length) {
                        orderCount = 0
                    }
                } else if (intervalType === "all") {
                    for (var songInterval of songs) {
                        songsForID[songInterval].ignoreSeperation = true;
                        newPlaylist.push(songsForID[songInterval])
                    }
                }
            }
            count = 0
        }
    }
    return newPlaylist
}

var addSongInfo = function (id, callback) {
    itframe.getSongInfo(id, (err, res) => {
        if (err) {
            return callback(err)
        }
        songsForID[id] = res
        callback(null)
    })
}

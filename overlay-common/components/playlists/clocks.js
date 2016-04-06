const itframe = requireFromRoot("components/itframe/api.js")
const wait = require("wait.for")
const _ = require("underscore")

module.exports = (callback) => {
    wait.launchFiber(() => {
        var clocks = wait.for(itframe.getClocks)
        var currentClock = selectCurrentClock(clocks)
        debug(currentClock)
        var songs = {}
        for (var tag of currentClock.tags) {
            songs[tag.tag] = wait.for(itframe.getMusicForTag, tag.tag)
        }
        var next100Songs = generatePlaylist(currentClock.tags, songs)
        callback(null, next100Songs)
    })
}

var selectCurrentClock = (clocks) => {
    var dayOfWeek = new Date().getDay()
    var hour = new Date().getHours()
    var minute = new Date().getMinutes()

    // sort clocks on time

    clocks.sort((a, b) => { return a.end.dayOfWeek - b.end.dayOfWeek })
    clocks.sort((a, b) => {
        if (a.end.dayOfWeek === b.end.dayOfWeek) {
            return a.end.hour - b.end.hour
        }
        return 0
        } )
    clocks.sort((a, b) => {
        if (a.end.dayOfWeek === b.end.dayOfWeek && a.end.hour === b.end.hour) {
            return a.end.minute - b.end.minute
        }
        return 0
        } )
    let reload = () => {
        global.connection.reloadClocks()
    }

    for (var id in clocks) {
        if (clocks.hasOwnProperty(id)) {
            new global.Cron(clocks[id].start.minute + " " + clocks[id].start.hour + " * * " + clocks[id].start.dayOfWeek, reload, null, true);
            if (clocks[id].end.dayOfWeek > dayOfWeek) {
                if ((id - 1) < 0) {
                    return clocks[id]
                }
                return clocks[id - 1]

            } else if (clocks[id].end.dayOfWeek === dayOfWeek) {
                if (clocks[id].end.hour > hour) {
                    return clocks[id]
                } else if (clocks[id].end.hour === hour) {
                    if (clocks[id].end.minute > minute) {
                        return clocks[id]
                    }
                }
            }
        }
    }

    return null
}

var generatePlaylist = function (tags, songsForTag) {
    var playlist = [] // max 100 songs

    for (var tag of tags) {
        debug("Selecting " + tag.percent + " songs from " + tag.tag)
        songsForTag[tag.tag] = _.shuffle(songsForTag[tag.tag])
        for (var i = 0; i < tag.percent; i++) { // select num procent songs
                if (songsForTag[tag.tag].hasOwnProperty(i)) {
                    debug("Adding " + songsForTag[tag.tag][i]._id)
                    playlist.push(songsForTag[tag.tag][i])
                } else {
                    var song = songsForTag[tag.tag][Math.floor((Math.random() * songsForTag[tag.tag].length))]
                    debug("Adding random song " + song._id)
                    playlist.push(song)
                }
         }
    }
    return _.shuffle(playlist)
}

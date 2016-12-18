import EventEmitter from "events"
import _ from "underscore"
import clocks from "../playlists/clocks.js"
import intervals from "../playlists/intervals.js"

export default () => {
    const exp = {}
    let clocksCheckInterval = null

    exp.queueEvents = new EventEmitter()
    exp.queue = []
    exp.currentSong = {}
    exp.isWriting = false

    exp.loadClocks = async () => {
        exp.isWriting = true
        exp.queue = exp.queue.concat(await intervals(await clocks()))
        exp.isWriting = false
        exp.queueEvents.emit("queueUpdate")
        setClocksInterval()
    }

    exp.reloadClocks = async () => {
        exp.isWriting = true
        exp.queue = await intervals(await clocks())
        exp.isWriting = false
        exp.queueEvents.emit("queueReset")
        setClocksInterval()
    }

    exp.loadEvent = async () => {
        exp.queueEvents.emit("queueReset")
        // do some more things
    }

    exp.queueEvents.on("playsSong", (song) => {
        if (exp.queue[0] && exp.queue[0]._id === song._id) {
            exp.isWriting = true
            exp.currentSong = _.clone(global.queueManager.queue[0])
            exp.queue.splice(0, 1)
            exp.isWriting = false
        }
    })

    const setClocksInterval = () => {
        clearInterval(clocksCheckInterval)
        clocksCheckInterval = setInterval(() => {
            if (!global.queueManager.isWriting) {
                console.log(`oooh still ${global.queueManager.queue.length} left`)
                if (global.queueManager.queue.length <= 5) {
                    debug("Loading in more songs")
                    global.queueManager.loadClocks()
                }
            }
        }, 1000)
    }

    return exp
}

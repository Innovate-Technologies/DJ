const wait = require("wait.for")
const clocks = requireFromRoot("components/playlists/clocks.js")
const intervals = requireFromRoot("components/playlists/intervals.js")

let connections = [];

let exportFunctions = {}

module.exports = () => {
    for (var stream of global.config.streams) {
        debug("Starting " + stream.stream)
        var engine = null // TO DO: ADD LS CLASS
        engine.init(stream)
        connections.push(engine)
    }
    return exportFunctions
}

exportFunctions.start = () => {
    wait.launchFiber(() => {
        global.isWritingQueue = true
        var list = wait.for(clocks)
        list = wait.for(intervals, list)

        for (let song of list) {
            for (let engine of connections) {
                engine.addToQueue(song)
            }
        }

        for (let engine of connections) {
            debug("Call engine")
            engine.startQueue()
        }
        global.isWritingQueue = false
    })
}

exportFunctions.loadClocks = () => {
    wait.launchFiber(() => {
        global.isWritingQueue = true
        var list = wait.for(clocks)
        exportFunctions.add(wait.for(intervals, list))
        global.isWritingQueue = false
    })
}

exportFunctions.reloadClocks = () => {
    exportFunctions.clearQueue()
    exportFunctions.loadClock()
}

exportFunctions.add = (list) => {
    for (let song of list) {
        for (let engine of connections) {
            engine.addToQueue(song)
        }
    }
}

exportFunctions.getQueue = () => {
    return connections[0].getQueue()
}

exportFunctions.clearQueue = () => {
    for (let engine of connections) {
        engine.clearQueue()
    }
}

exportFunctions.skip = () => {
    for (let engine of connections) {
        engine.skipSong()
    }
}


/* Live Input */


exportFunctions.startLive = (stream) => {
    for (let engine of connections) {
        engine.startLive(stream)
    }
}

exportFunctions.endLive = () => {

}

exportFunctions.sendMetadata = (song, dj) => {
    for (let engine of connections) {
        engine.sendMetadata(song, dj)
    }
}

exportFunctions.checkLiveToken = (token) => {
    return true
}

exportFunctions.isLive = () => {
    return false
}

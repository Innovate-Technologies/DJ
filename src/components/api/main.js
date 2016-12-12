export default ({app, io}) => {
    app.use((req, res, next) => {
        res.setHeader("X-Powered-By", "DJ/" + global.config.version.DJ)
        res.setHeader("server", "DJ/" + global.config.version.DJ)
        next();
    });
    app.use("/api/:key/*", (req, res, next) => {
        if (req.params.key !== global.config.apikey) {
            return res.status(401).send({ status: "error", error: "Invalid key" })
        }
        next();
    });

    app.get("/api/:key/songs/queue", (req, res) => {
        res.json(getQueue())
    })

    app.get("/api/:key/songs/current", (req, res) => {
        res.json(getCurrentSong())
    })

    app.post("/api/:key/songs/skip", (req, res) => {
        global.connection.skipSong()
        res.json({ status: "ok" })
    })

    app.post("/api/:key/clocks/reload", (req, res) => {
        global.queueManager.reloadClocks()
        res.json({ status: "ok" })
    })

    io.of("/queueEvents").on("connection", (socket) => {
        socket.on("key", (key) => {
            if (key === global.config.apikey) {
                socket.emit("currentSong", getCurrentSong())
                socket.emit("queue", getQueue())
                global.queueManager.queueEvents.on("queueUpdate", () => {
                    socket.emit("queue", getQueue())
                })
                global.queueManager.queueEvents.on("queueReset", () => {
                    socket.emit("queue", getQueue())
                })
                global.queueManager.queueEvents.on("playsSong", () => {
                    socket.emit("currentSong", getCurrentSong())
                    socket.emit("queue", getQueue())
                })
            } else {
                socket.emit("error", "invalid key")
            }
        });
    });
}

const getCurrentSong = () => {
    const currentSong = JSON.parse(JSON.stringify(global.queueManager.currentSong)) // dirty clone
    currentSong.internalURL = null
    currentSong.processedURLS = null
    return currentSong
}

const getQueue = () => {
    const queue = JSON.parse(JSON.stringify(global.queueManager.queue)) // dirty clone
    for (let id in queue) {
        if (queue.hasOwnProperty(id)) {
            queue[id].internalURL = null
            queue[id].processedURLS = null
        }
    }
    return queue
}

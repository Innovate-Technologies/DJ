export default ({app, io}) => {
    app.use((req, res, next) => {
        res.setHeader("X-Powered-By", "DJ/" + global.config.version.DJ)
        res.setHeader("server", "DJ/" + global.config.version.DJ)
        next();
    });
    app.use("/api/:key/*", (req, res, next) => {
        if (req.params.key !== global.config.apikey) {
            return res.status(401).send({status: "error", error: "Invalid key"})
        }
        next();
    });

    app.get("/api/:key/queue", (req, res) => {
        res.json(connection.getQueue())
    })

    app.post("/api/:key/song/skip", (req, res) => {
        global.connection.skipSong()
        res.json({status: "ok"})
    })

    app.post("/api/:key/clocks/reload", (req, res) => {
        global.queueManager.reloadClocks()
        res.json({status: "ok"})
    })
}

export default ({app, io}) => {
    app.use((req, res, next) => {
        res.setHeader("X-Powered-By", "DJ/" + global.config.version.DJ)
        res.setHeader("server", "DJ/" + global.config.version.DJ)
        next();
    });
    app.use("/private/:key/", (req, res, next) => {
        if (req.params.key !== global.config.internal.dj.key) {
            return res.status(401).send({status: "error", error: "Invalid key"})
        }
        next();
    })

    app.get("/private/:key/queue", (req, res) => {
        res.json(connection.getQueue())
    })

    app.post("/private/:key/song/skip", (req, res) => {
        global.connection.skip()
        res.json({status: "ok"})
    })

    app.post("/private/:key/clocks/reload", (reqq, res) => {
        global.connection.reloadClocks()
        res.json({status: "ok"})
    })
}

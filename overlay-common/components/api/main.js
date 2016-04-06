module.exports = ({app}) => {
    app.use(function (req, res, next) {
        res.setHeader("X-Powered-By", "DJ/" + global.config.version.DJ)
        next();
    });
    app.use("/private/:key/", (req, res, next) => {
        if (req.params.key !== global.config.internal.dj.key) {
            return res.status(401).send({status: "error", error: "Invalid key"})
        }
        next();
    })
    app.post("/private/:key/song/skip", (req, res) => {
        global.connection.skip()
        res.send({status: "ok"})
    })
    app.post("/private/:key/clocks/reload", (reqq, res) => {
        global.connection.reloadClocks()
        res.send({status: "ok"})
    })
}

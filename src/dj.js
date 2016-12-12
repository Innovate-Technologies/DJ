import express from "express"
import { getConfig } from "./components/itframe/api"

if (process.env.username) {
    global.djconfig.username = process.env.username
}

if (!global.djconfig.username) {
    console.log("No username passed")
    process.exit(1)
}

const engine = global.djconfig.DJEngine ? global.djconfig.DJEngine : "liquidsoap"

const app = express();
require("http").createServer(app).listen(80);


(async () => {
    try {
        console.log("     _____        ___    \n    /  /::\\      /  /\\   \n   /  /:/\\:\\    /  /:/   \n  /  /:/  \\:\\  /__/::\\   \n /__/:/ \\__\\:| \\__\\/\\:\\  \n \\  \\:\\ /  /:/    \\  \\:\\ \n  \\  \\:\\  /:/      \\__\\:\\\n   \\  \\:\\/:/       /  /:/\n    \\  \\::/       /__/:/ \n     \\__\\/        \\__\\/  \n                         \n")
        console.log("Copyright 2015-2016 Innovate Technologies")
        console.log("------------------------------------")
        global.config = await getConfig(global.djconfig.username)
        console.log(global.config)

        if (global.config.timezone) {
            process.env.TZ = global.config.timezone
        }
        // making sure Date() isn't called till now

        global.at = require("node-at")
        global.cron = require("cron").CronJob
        global.queueManager = require("./components/songs/queue.js")()

        await global.queueManager.loadClocks()

        global.connection = requireFromRoot("components/" + engine + "/connect.js")()
        connection.start()
    } catch (err) {
        debug(err)
    }
})()

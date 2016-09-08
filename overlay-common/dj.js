import { getConfig } from "./components/itframe/api"
import express from "express"

global.isWritingQueue = false;

if (process.env.username) {
    global.djconfig.username = process.env.username
}

if (!global.djconfig.username) {
    console.log("No username passed")
    process.exit(1)
}

const engine = global.djconfig.DJEngine ? global.djconfig.DJEngine : "liquidsoap"

global.at = require("node-at")
global.cron = require("cron").CronJob

const app = express();
require("http").createServer(app).listen(8080);


(async () => {
    try {
        global.isWritingQueue = true
        console.log("     _____        ___    \n    /  /::\\      /  /\\   \n   /  /:/\\:\\    /  /:/   \n  /  /:/  \\:\\  /__/::\\   \n /__/:/ \\__\\:| \\__\\/\\:\\  \n \\  \\:\\ /  /:/    \\  \\:\\ \n  \\  \\:\\  /:/      \\__\\:\\\n   \\  \\:\\/:/       /  /:/\n    \\  \\::/       /__/:/ \n     \\__\\/        \\__\\/  \n                         \n")
        console.log("Copyright 2015-2016 Innovate Technologies")
        console.log("------------------------------------")
        global.config = getConfig(global.djconfig.username)
        console.log(global.config)
        connection = requireFromRoot("components/" + engine + "/connect.js")()
        connection.loadClocks()
        connection.start()
        global.isWritingQueue = false
    } catch (err) {
        debug(err)
    }
})()


// Check if the Clock has enough songs left to play
setInterval(() => {
    if (!global.isWritingQueue) {
        if (global.connection.getQueue().length <= 2) {
            debug("Loading in more songs")
            global.connection.loadClocks()
        }
    }
}, 1000)

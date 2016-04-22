import { getConfig } from "./components/itframe/api"
import express from "express"

global.isWritingQueue = false;

if (!process.env.username) {
    console.log("No username passed")
    process.exit(1)
}

const engine = process.env.DJEngine ? process.env.DJEngine : "liquidsoap"

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
        global.config = getConfig(process.env.username)
        console.log(global.config)
        global.connection = requireFromRoot("components/" + engine + "/connect.js")()
        global.connection.loadClocks()
        global.connection.start()
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

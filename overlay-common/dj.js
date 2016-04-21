if (!process.env.compiled) {
    require("babel-polyfill");
    require("babel-register");
}

global.debug = function (log) {
    if (process.env.DEBUG) {
        console.log("[" + new Date().toTimeString() + "] " + log);
    }
}

global.appRoot = __dirname;
global.requireFromRoot = function (path) {
    debug("Requiring " + arguments[0]);
    return require(global.appRoot + "/" + path);
};

global.isWritingQueue = false;

if (!process.env.username) {
    console.log("No username passed")
    process.exit(1)
}

const engine = process.env.DJEngine ? process.env.DJEngine : "liquidsoap"

global.at = require("node-at")
global.Cron = require("cron").CronJob

const express = require("express")
var app = express();
require("http").createServer(app).listen(80);

const itframe = requireFromRoot("components/itframe/api.js")
const wait = require("wait.for")

wait.launchFiber(() => {
    global.isWritingQueue = true
    console.log("     _____        ___    \n    /  /::\\      /  /\\   \n   /  /:/\\:\\    /  /:/   \n  /  /:/  \\:\\  /__/::\\   \n /__/:/ \\__\\:| \\__\\/\\:\\  \n \\  \\:\\ /  /:/    \\  \\:\\ \n  \\  \\:\\  /:/      \\__\\:\\\n   \\  \\:\\/:/       /  /:/\n    \\  \\::/       /__/:/ \n     \\__\\/        \\__\\/  \n                         \n")
    console.log("Copyright 2015-2016 Innovate Technologies")
    console.log("------------------------------------")
    global.config = wait.for(itframe.getConfig, process.env.username)
    global.connection = requireFromRoot("components/" + engine + "/connect.js")()
    global.connection.loadClocks()
    global.connection.start()
    global.isWritingQueue = false
})


// Check if the Clock has enough songs left to play
setInterval(() => {
    if (!global.isWritingQueue) {
        if (global.connection.getQueue().length <= 2) {
            debug("Loading in more songs")
            global.connection.loadClocks()
        }
    }
}, 1000)

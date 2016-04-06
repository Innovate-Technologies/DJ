var net = require("net");

var runCommand = function (command, callback) {
    var calledBack = false
    var client = new net.Socket();
    client.connect(1234, "127.0.0.1", function () {
        client.write(command + "\n");
    });
    client.on("data", function () {
        if (!calledBack) {
            calledBack = true
            client.end()
            return callback(null, true)
        }
    });
}

var runCommandWithOutput = function (command, callback) {
    var callbackData = []
    var client = new net.Socket();
    client.connect(1234, "127.0.0.1", function () {
        client.write(command + "\n");
    });
    client.on("data", function (data) {
        console.log(JSON.stringify(data.toString()))
        if (data.toString() === "END") {
            client.end()
            return callback(null, callbackData)
        }
        callbackData.push(data.toString())
    });
}

var pushToQueue = function (url, callback) {
    runCommand("request.equeue_4773.push " + url, function (err, res) {
        if (err) {
            callback(err)
            return
        }
        callback(null, JSON.stringify(res))
    })
}

var nextSong = function (callback) {
    runCommand("output(dot)shoutcast.skip", function (err, res) {
        if (err) {
            callback(err)
            return
        }
        callback(null, JSON.stringify(res))
    })
}

module.exports.cleanQueue = function (callback) {
    runCommandWithOutput("request.equeue_4773.queue", function (err, res) {
        if (err) {
            return callback(err)
        }
        console.log(res)
    })
}


module.exports.pushToQueue = pushToQueue
module.exports.nextSong = nextSong

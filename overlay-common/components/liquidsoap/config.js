/* global global */
var fs = require("fs")
module.exports.writeStreams = (config) => {
	var configFile = ""
	for (var stream of config.streams) {
		configFile += "output.shoutcast(%mp3(bitrate=" + stream.stream.replace("kbps", "") + "), host=\"opencast.radioca.st\",port=" + config.input.SHOUTcast + ", password=\"" + stream.password + "\",fall,fallible=true,name=\"Swift\")\n"
	}
	fs.writeFileSync(global.appRoot + "/output.liq", configFile)
}

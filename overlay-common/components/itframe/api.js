var rest = require("restler")

module.exports.getConfig = (username, callback) => {
	rest.post("https://itframe.innovatete.ch/cast/config", {
		timeout: 100000,
		data: {
			username: username,
			token: process.env.ITFrameToken,
		},
	}).on("complete", (body) => {
		callback(null, body);
	}).on("timeout", () => {
		callback(new Error("Timeout"))
	})
}

module.exports.getAllMusic = (callback) => {
	rest.get(djBaseUrl() + "all-songs", {
		timeout: 100000,
	}).on("complete", (body) => {
		callback(null, body);
	}).on("timeout", () => {
		callback(new Error("Timeout"))
	})
}

module.exports.getSongInfo = (id, callback) => {
	rest.get(djBaseUrl() + "song/" + id, {
		timeout: 100000,
	}).on("complete", (body) => {
		callback(null, body);
	}).on("timeout", () => {
		callback(new Error("Timeout"))
	})
}

module.exports.getMusicForTag = (tag, callback) => {
	rest.get(djBaseUrl() + "songs-with-tag/" + tag, {
		timeout: 100000,
	}).on("complete", (body) => {
		callback(null, body);
	}).on("timeout", () => {
		callback(new Error("Timeout"))
	})
}

module.exports.getClocks = (callback) => {
	rest.get(djBaseUrl() + "all-clocks", {
		timeout: 100000,
	}).on("complete", (body) => {
		callback(null, body);
	}).on("timeout", () => {
		callback(new Error("Timeout"))
	})
}

module.exports.getIntervals = (callback) => {
	rest.get(djBaseUrl() + "all-intervals", {
		timeout: 100000,
	}).on("complete", (body) => {
		callback(null, body);
	}).on("timeout", () => {
		callback(new Error("Timeout"))
	})
}

var djBaseUrl = () => {
	return "https://itframe.innovatete.ch/dj/" + global.config.username + "/" + global.config.internal.dj.key + "/"
}

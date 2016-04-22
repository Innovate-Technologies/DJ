/**
 * This file sets essential global variables and then bootstraps the app
 * after enabling ES6-syntax using Babel.
 *
 * Only ES5-compatible syntax should be used in this file,
 * as Babel hasn't been loaded yet. Keep this file slim as its sole role
 * is to set up essential globals and bootstrap the app.
 */

global.appRoot = __dirname;

global.requireFromRoot = function (path) {
    debug("Requiring " + arguments[0]);
    return require(global.appRoot + "/" + path);
};

global.debug = function (log) {
    if (process.env.DEBUG) {
        console.log("[" + new Date().toLocaleTimeString() + "] " + log);
    }
}

if (!process.env.compiled) {
    require("babel-polyfill");
    require("babel-register");
}
require("./dj.js");

import { getClocks, getMusicForTag } from "../itframe/api"
import _ from "underscore"

export default async () => {
    debug("Loading clocks")
    const clocks = await getClocks()
    const currentClock = getClockForDayHourMinute(clocks, new Date().getDay(), new Date().getHours(), new Date().getMinutes())
    debug(currentClock)
    const songs = {}
    for (let tag of currentClock.tags) {
        songs[tag.tag] = await getMusicForTag(tag.tag)
    }
    const next100Songs = generatePlaylist(currentClock.tags, songs)
    debug("Clocks ready")
    return next100Songs
}


const generatePlaylist = function (tags, songsForTag) {
    const playlist = [] // max 100 songs

    for (var tag of tags) {
        debug("Selecting " + tag.percent + " songs from " + tag.tag)
        songsForTag[tag.tag] = _.shuffle(songsForTag[tag.tag])
        for (let i = 0; i < tag.percent; i++) { // select num procent songs
            if (songsForTag[tag.tag].hasOwnProperty(i)) {
                debug("Adding " + songsForTag[tag.tag][i]._id)
                playlist.push(songsForTag[tag.tag][i])
            } else {
                const song = songsForTag[tag.tag][Math.floor((Math.random() * (songsForTag[tag.tag].length - 1)))]
                if (song) {
                    debug("Adding random song " + song._id)
                    playlist.push(song)
                }
            }
        }
    }
    return _.shuffle(playlist)
}

const getClockForDayHourMinute = (clocks, day, hour, minute) => {
    for (let id in this.clocks) {
        if (this.clocks.hasOwnProperty(id)) {
            if (this.clocks[id].start.dayOfWeek < day && this.clocks[id].end.dayOfWeek > day) {
                // ] day [
                return this.clocks[id];
            } else if ((this.clocks[id].start.dayOfWeek === day || this.clocks[id].end.dayOfWeek === day)) {
                // [ day ]
                if ((this.clocks[id].start.dayOfWeek === day && this.clocks[id].start.hour <= hour) || (this.clocks[id].end.dayOfWeek === day && this.clocks[id].end.hour >= hour)) {
                    // check end minutes
                    // [ day ] [ hour ]
                    if (this.clocks[id].start.hour < hour && this.clocks[id].end.hour > hour) {
                        // ] hour [
                        return this.clocks[id];
                    } else if ((this.clocks[id].start.hour === hour && this.clocks[id].start.minute >= minute) || (this.clocks[id].end.hour === hour && this.clocks[id].end.minute >= minute)) {
                        // [ day ] [ hour ] [ minute ]
                        return this.clocks[id];
                    }
                }
            }
        }
    }
    return null;
}
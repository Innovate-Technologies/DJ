import rest from "restler"

export const getConfig = (username) => new Promise((resolve, reject) => {
    rest.post("https://itframe.innovatete.ch/cast/config", {
        timeout: 100000,
        data: {
            username: username,
            token: global.djconfig.ITFrameToken,
        },
    }).on("complete", body => resolve(body)).on("error", err => reject(err)).on("timeout", err => reject(err))
})

/*export const getConfig = (username) => {
    console.log("https://itframe.innovatete.ch/cast/config")
    rest.post("https://itframe.innovatete.ch/cast/config", {
        timeout: 100000,
        data: {
            username: username,
            token: process.env.ITFrameToken,
        },
    }).on("complete", (body) => {
        console.log(body)
    })
}*/


const getFromITFrame = (query) => new Promise((resolve, reject) => {
    rest.post(`https://itframe.innovatete.ch/dj/${global.config.username}/${global.config.internal.dj.key}/${query}`, {
        timeout: 100000,
    }).on("complete", body => resolve(body)).on("timeout", err => reject(err))
})

export const getAllMusic = async () => {
    return await getFromITFrame("all-songs")
}

export const getSongInfo = async (id) => {
    return await getFromITFrame(`song/${id}`)
}

export const getMusicForTag = async (tag) => {
    return await getFromITFrame(`songs-with-tag/${tag}`)
}

export const getClocks = async () => {
    return await getFromITFrame("all-clocks")
}

export const getIntervals = async () => {
    return await getFromITFrame("all-intervals")
}

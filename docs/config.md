# DJ Configuration

DJ shares the configuration file of Cast. In a production environment these files are stored and fetched from ITFrame. DJ still has a 2nd configuration file to bootstrap the setup to the local machine/container

## Extended Cast configuration
*Note: DJ also reads some properties of Cast's configuration file. The extra options below are not enough.*

```json
{
  "name": "Swift",
  "genre": "Misc",  
  "version": {
    "DJ": "0"
  },
  "DJ": {
    "enabled": false,
    "fadeLength": 10
  },
  "internal": {
    "dj": {
      "key": "keepmesecret"
    }
  },
}
```
`name` and `genre` are sent to the source on connecting as a normal encoder should. They are sent to all streams.

`version` this option was added internally for tracking updates. Currently DJ also sends this tag with it's API. It map be replaced with a semver version once it is out of testing (as Cast did).

`DJ` is the general option for the DJ settings. These settings are meant to also be sent to Control as the client may edit these. `enabled` is not checked within DJ but is used for ITFrame to determine wether to start the DJ container or not. `fadeLength` is as the name states the length in seconds for crossfades.

`Internal` is meant never to be sent to the client. It is designed to contain objects like the `key` which is the API key used so communicate with DJ.

## DJ (bootstrap) config.json
The config.json file gets loaded on the start of the application. It contains all info it needs to connect to ITFrame and loads it's components
```json
{
  "username": "opencast",
  "ITFrameToken": "secret",
  "engine": "liquidsoap"
}
```
`username` contains the username it has to fetch the info for from ITFrame. This can still be set overwritten via an environment variable as it is used in containers.

`ITFrameToken` is a secret token to be sent to ITFrame to fetch the configuration

`engine` the audio engine loaded in from `components/${engine}/connect.js`

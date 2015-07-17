# gofip

A simple FIP Radio player.
Displays the current/previous/next music that are/will be played on the FIP webradio.
It will also send a notification when the music changes using the system's default notification system.

![gofip](https://cloud.githubusercontent.com/assets/694365/8745908/7750a1c2-2c84-11e5-97fe-da9d154a4350.png)

## Requirements
You need to have GStreamer on your computer, otherwise the streaming capability won't work.

## Build
Here are the go dependencies you need to install in order to build it yourself.

```
go get "github.com/0xAX/notificator"
go get "github.com/andlabs/ui"
go get "github.com/ziutek/gst"
```

## TODO

 - Parse the startTime and endTime in the JSON response to request the server only when the music actually changes. (Currently requests the server every minute.)

## Known issues

 - When the streams stops (due to instable connection for example) it won't sync and it will keep reading the stream. I should add an option to resync the stream. 

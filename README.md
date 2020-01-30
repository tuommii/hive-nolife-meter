# Hive-Nolife-Meter

In Hive you can't see other people progress (levels) otherwise than going to
their profile page or via API. So i made app for learning purposes, that crawl's data and then creates a static
webpage. This app doesn't utilize goroutines because of the API limits.

## Screenshot

UI is heavily inspired by UNIX =)
![Screenshot](https://github.com/tuommii/hive-nolife-meter/blob/master/screenshot.png "Screenshot")

## Setup

* You might have to install `go get golang.org/x/net/html`
* Configure variables and add your cookie in `run.sh`
* You can get your cookie from `Chrome DevTools -> Application Tab -> Cookies`
* Add users to `users.json`
* Edit `Makefile`
* run `run.sh`

## Test
* In project root: `make test`

# Hive-Nolife-Meter

In Hive you can't see other people progress (levels) otherwise than going to
their profilepage. So i made app that crawls needed data and then create's a static
webpage. This app doesn't utilize goroutines because i don't know API limits. It only uses *[go's](https://golang.org/)* **awesome** standard library. No other dependencies needed.

## Screenshot

UI is heavily inspired by UNIX =)
![Test .PNG](https://github.com/tuommii/hive-nolife-meter/blob/master/screenshot..png)

## Setup

* Configure variables and add your cookie in `run.sh`
* You can get your cookie from `Chrome DevTools -> Application Tab -> Cookies`
* Add users to `users.json`
* Edit `Makefile`
* run `run.sh`

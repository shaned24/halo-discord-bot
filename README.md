# Halo Discord Bot

This bot can respond to a user command `!serviceRecord GamerTag` which
will go fetch multiplayer stats from [autocode](https://autocode.com/lib/halo/infinite/).


## To test

Clone this repo

```shell
export DISCORD_TOKEN=<your bot token>
export AUTOCODE_TOKEN=<your autocode token>
go run main.go
```

Then in discord type
```shell
!serviceRecord GamerTag
```
package lib

import (
	"uilliam.com/sip/types"
)

// GenerateLoggingCode creates two code snippets:
//
//  1. The logger worker (/logger/logger.go) that processes logs from a channel.
//
//  2. The logger(level, message) function in main, which sends logs into the channel.
//
//     "logging": {
//     "enabled": true,
//     "type": "File", // or "connection" or "Console"
//     "connection": "MyLog.log", // can be a file path, connection name, etc.
//     "level": "info" // or "debug", "error", "warn"
//     }
func GenerateLoggingCode(logging types.LoggingDefinition, path string) (string, []string) {

}

// i could not get this working tonight so I will try again tomorrow, I think the best way to do this is to:

// the way I want this to work is the codegen will return a function called logger(level, message) and
// everything in the app will call logger because we may generate other parts of the app after the logger
// part so we need a func that we can all nomater what.,I think the best think to do is to have a new part
// of the lib called logger and live in / logger this will live on a separate thread, the logger func that
// is in the main file will enabled is true put the level and message into a pool then return and if false
// it will just return right away

// so to recap this file I will make will be 2 parted, part one we will construct the logger lib and
// build the file under /logger/logger.go (this func will be its own thread in the end so it will need to accept
// a pool as its input and will be doing the heavy lifting as far as logic for where and what should be
// logged based on what gets put into the pool)

// then second we will build a logger function and return that code into the main

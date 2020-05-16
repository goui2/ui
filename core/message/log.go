package message

import "log"

const logPackageName = "goui.core.message"

var logIt func(fmt string, v ...interface{}) = func(fmt string, v ...interface{}) {
	log.Printf(logPackageName+">"+fmt, v...)
}

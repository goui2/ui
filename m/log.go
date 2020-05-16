package m

const logPackageName = "goui.m"

var logIt func(fmt string, v ...interface{}) = func(fmt string, v ...interface{}) {
	//log.Printf(logPackageName+">"+fmt, v...)
}

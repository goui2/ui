module github.com/goui2/ui/samples/s02

go 1.14


require (
	github.com/goui2/ui/base v0.0.0
	github.com/goui2/ui/com v0.0.0
	github.com/goui2/ui/m v0.0.0
	github.com/goui2/ui/model v0.0.0
	github.com/goui2/ui/core v0.0.0
	github.com/goui2/ui/core/mvc v0.0.0
)

replace github.com/goui2/ui/m => ../../m
replace github.com/goui2/ui/com => ../../com
replace github.com/goui2/ui/model => ../../model
replace github.com/goui2/ui/base => ../../base

replace github.com/goui2/ui/core => ../../core

replace github.com/goui2/ui/core/mvc => ../../core/mvc
replace github.com/goui2/ui/core/message => ../../core/message

replace github.com/goui2/ui/core/mvc/brick => ../../core/mvc/brick
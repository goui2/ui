module github.com/goui2/ui/core/mvc

go 1.14

replace github.com/goui2/ui/core => ../

replace github.com/goui2/ui/core/mvc/brick => ./brick

replace github.com/goui2/ui/base => ../../base

replace github.com/goui2/ui/com => ../../com

require (
	github.com/goui2/ui/base v0.0.0-00010101000000-000000000000
	github.com/goui2/ui/com v0.0.0-00010101000000-000000000000
	github.com/goui2/ui/model v0.0.0-00010101000000-000000000000 // indirect
	github.com/goui2/ui/core v0.0.0-00010101000000-000000000000
	github.com/goui2/ui/core/mvc/brick v0.0.0-00010101000000-000000000000
)

replace github.com/goui2/ui/model => ../../model

replace github.com/goui2/ui/core/message => ../message

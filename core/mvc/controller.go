package mvc

import "github.com/goui2/ui/com"

type Controller interface {
	Handler(name string) (com.EventHandler, bool)
	Init()
	Destroy()
}

type ControllerFactory func(view View) Controller

var (
	controllerFactories = make(map[string]ControllerFactory)
)

func RegisterController(name string, cf ControllerFactory) {
	controllerFactories[name] = cf
}

func LoadController(name string) (ControllerFactory, bool) {
	cf, ok := controllerFactories[name]
	return cf, ok
}

package base

import (
	"sync"
	"time"

	"github.com/goui2/ui/com"
)

var (
	queuedEvents chan func()
)

func init() {
	queuedEvents = make(chan func(), 10000)
	go func() {
		for {
			for fn := range queuedEvents {
				//				logIt("eventloop take event")
				func() {
					defer func() {
						r := recover()
						if r != nil {
						}
					}()
					fn()
				}()
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()
}

type EventProvider interface {
	Object
	AttachEvent(eventId string, data com.EventData, fn com.EventHandler)
	AttachEventOnce(eventId string, data com.EventData, fn com.EventHandler)
	FireEvent(eventId string, param com.EventParam)
	DetachEvent(eventId string, fn com.EventHandler)
	HasListener(eventId string) bool
	Destroy()
	GetSource() EventProvider
}

func createEventProvider(id string, epm EventProviderMetadata, parent Constructor, sets ...InstanceSetting) EventProvider {
	ep := &eventProvider{
		events: make(map[string]map[com.EventHandler]eventHandler),
		mutex:  &sync.Mutex{},
	}
	for _, e := range epm.AllEvents() {
		ep.events[e.Name] = make(map[com.EventHandler]eventHandler)
	}
	parentSet := AdjustSelfSetting(ep, epm, sets)
	ep.Object = parent.New(id, parentSet...)

	return ep
}

type eventProvider struct {
	Object
	events map[string]map[com.EventHandler]eventHandler
	mutex  *sync.Mutex
}

type eventHandler struct {
	data com.EventData
	fn   com.EventHandler
	once bool
}

func (ep *eventProvider) AttachEvent(eventId string, data com.EventData, fn com.EventHandler) {
	ep.mutex.Lock()
	defer ep.mutex.Unlock()
	if handlers, ok := ep.events[eventId]; ok {
		handlers[fn] = eventHandler{
			data: data,
			fn:   fn,
		}
	} else {
		panic("unknown event: " + eventId)
	}

}

func (ep *eventProvider) DetachEvent(eventId string, fn com.EventHandler) {
	ep.mutex.Lock()
	defer ep.mutex.Unlock()
	if handlers, ok := ep.events[eventId]; ok {
		delete(handlers, fn)
	} else {
		panic("unknown event: " + eventId)
	}

}

func (ep *eventProvider) AttachEventOnce(eventId string, data com.EventData, fn com.EventHandler) {
	ep.mutex.Lock()
	defer ep.mutex.Unlock()
	if handlers, ok := ep.events[eventId]; ok {
		handlers[fn] = eventHandler{
			data: data,
			fn:   fn,
			once: true,
		}
	} else {
		panic("unknown event: " + eventId)
	}

}
func (ep *eventProvider) doFireEventItem(event com.Event, h com.EventHandler) {
	func() {
		defer func() {
			r := recover()
			if r != nil {
				//				logIt("EventProvider(%v).FireEventRecover(%v)", ep, r)
			}
		}()
		h.Handle(event)
	}()

}

func (ep *eventProvider) doFireEvent(eventId string, parameters com.EventParam) {
	if handlers, ok := ep.events[eventId]; ok {
		//		logIt("EventProvider.doFireEvent(%s,%v,handlers(%d))", eventId, parameters, len(handlers))

		toDelete := make([]com.EventHandler, 0)
		for k := range handlers {
			handler := handlers[k]
			event := &event{
				eventId:    "1",
				eventName:  eventId,
				data:       handler.data,
				source:     ep,
				parameters: parameters,
			}
			//			logIt("EventProvider.FireEvent(%s,%v,%v,%v) -- %p", eventId, parameters, k, handler, ep.GetSource())
			ep.doFireEventItem(event, handler.fn)
			if handler.once {
				toDelete = append(toDelete, k)
			}
		}
		for _, k := range toDelete {
			delete(handlers, k)
		}
	} else {
		panic("unknown event: " + eventId)
	}
}

func (ep *eventProvider) FireEvent(eventId string, parameters com.EventParam) {
	//	logIt("EventProvider.FireEvent(%s,%v) -- %p", eventId, parameters, ep.GetSource())
	queuedEvents <- func() {
		ep.doFireEvent(eventId, parameters)
	}
	//	logIt("EventProvider.FireEvent leave(%s,%v) -- %p", eventId, parameters, ep.GetSource())
}

func (ep *eventProvider) HasListener(eventId string) bool {
	ep.mutex.Lock()
	defer ep.mutex.Unlock()
	if handlers, ok := ep.events[eventId]; ok {
		return len(handlers) > 0
	} else {
		panic("unknown event: " + eventId)
	}
}

func (ep *eventProvider) Destroy()                 {}
func (ep *eventProvider) GetSource() EventProvider { return ep }

type event struct {
	Object
	eventId    string
	eventName  string
	data       com.EventData
	source     Object
	parameters com.EventParam
}

func (e *event) EventId() string       { return e.eventId }
func (e *event) Data() com.EventData   { return e.data }
func (e *event) Source() Object        { return e.source }
func (e *event) Param() com.EventParam { return e.parameters }
func (e *event) Type() string          { return e.eventName }
func (e *event) Prev() com.Event       { return nil }

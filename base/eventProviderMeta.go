package base

type EventDef struct {
	Name    string
	Details map[string]interface{}
}

type EventProviderMetadata interface {
	MetaData
	AllEvents() []EventDef
	Event(s string) (EventDef, bool)
}

type eventProviderMetadata struct {
	MetaData
	events []EventDef
}

func (epm eventProviderMetadata) AllEvents() []EventDef {
	return epm.events
}

func (epm eventProviderMetadata) Event(s string) (e EventDef, ok bool) {
	for _, ed := range epm.events {
		if ed.Name == s {
			return ed, true
		}
	}
	return
}

func eventProviderMetadataBuilder(s string, ms ...MetaDataSetting) []MetaDataSetting {
	newEPM := &eventProviderMetadata{
		events: make([]EventDef, 0),
	}
	selector := func(mds MetaDataSetting) bool {
		switch item := mds.(type) {
		case MDSettingEventDef:
			newEPM.events = append(newEPM.events, item.EventDef)
			return true
		}
		return true
	}
	parentSettings := SelectMetadataSettings(ms, selector)
	parentSettings = append(
		parentSettings,
		MDSettingCallback{1,
			func(m MetaData, pre MetaData) MetaData {
				newEPM.MetaData = m
				if preEPM, ok := pre.(EventProviderMetadata); ok {
					newEPM.events = append(newEPM.events, preEPM.AllEvents()...)
				}
				return newEPM
			},
		},
	)
	return parentSettings
}

func constructEventProvider(m MetaData, embbed Constructor) Constructor {
	epm := m.(EventProviderMetadata)
	return func(id string, is ...InstanceSetting) Object {
		ep := createEventProvider(id, epm, embbed, is...)
		return ep
	}
}

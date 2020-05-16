package base

type InstanceSetting interface {
	IType() SettingType
}

type SettingType int

const (
	Unspecific SettingType = iota
	Required
	Overwrite
	Combine
)

type MetaDataSetting interface {
	Mtype() SettingType
}

func SelectInstanceSettings(is []InstanceSetting, f func(InstanceSetting) bool) []InstanceSetting {
	result := make([]InstanceSetting, 0)
	for _, i := range is {
		if f(i) {
			result = append(result, i)
		}
	}
	return result
}

func SelectMetadataSettings(is []MetaDataSetting, f func(MetaDataSetting) bool) []MetaDataSetting {
	result := make([]MetaDataSetting, 0)
	for _, i := range is {
		if f(i) {
			result = append(result, i)
		}
	}
	return result
}

type SelfSetting struct {
	Object
	MetaData
}

func (_ SelfSetting) IType() SettingType {
	return Required
}

func AdjustSelfSetting(self Object, m MetaData, ms []InstanceSetting) []InstanceSetting {
	for _, r := range ms {
		if _, ok := r.(SelfSetting); ok {
			return ms
		}
	}
	return append(ms, SelfSetting{self, m})
}

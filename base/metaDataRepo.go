package base

var (
	metaDataRepo map[string]MetaData
)

func registerMetaData(m MetaData) {
	if metaDataRepo == nil {
		metaDataRepo = make(map[string]MetaData)
	}
	metaDataRepo[m.GetName()] = m
}

func LoadMetaData(n string) (m MetaData, ok bool) {
	m, ok = metaDataRepo[n]
	return
}

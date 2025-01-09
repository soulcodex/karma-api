package xjsonapiresponse

type Metadata []MetadataItem

func NewMetadata(items ...MetadataItem) Metadata {
	return items
}

func (m Metadata) MetadataMap() map[string]interface{} {
	metadata := map[string]interface{}{}
	for _, item := range m {
		metadata[item.Key] = item.Value
	}

	return metadata
}

type MetadataItem struct {
	Key   string
	Value interface{}
}

func NewMetadataItem(key string, value interface{}) MetadataItem {
	return MetadataItem{
		Key:   key,
		Value: value,
	}
}

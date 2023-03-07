package gojson

type FormatJSONProvider interface {
	AddFormatOption(options ...JSONFormatOption)
	FormatJSONData(data []byte) ([]byte, error)
	UpdateTemplate(rawTemplate []byte) error
}

func NewFormatJSONProvider(rawTemplate []byte, options ...JSONFormatOption) (FormatJSONProvider, error) {
	return newFormatItemImpl(rawTemplate, options...)
}

func NewDefaultFormatJSONProvider(rawTemplate []byte) (FormatJSONProvider, error) {
	return NewFormatJSONProvider(rawTemplate, DefaultOptions...)
}

func (fii *formatItemsImpl) AddFormatOption(options ...JSONFormatOption) {
	fii.addFormatOption(options...)
}

func (fii *formatItemsImpl) UpdateTemplate(rawTemplate []byte) error {
	return fii.updateTemplate(rawTemplate)
}

func (fii *formatItemsImpl) FormatJSONData(data []byte) ([]byte, error) {
	return fii.formatJSONData(data)
}

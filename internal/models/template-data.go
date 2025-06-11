package models

// TemplateData represents the data passed into handlers and shown in the templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

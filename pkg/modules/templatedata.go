package modules

type TemplateData struct {
	StringMap                        map[string]string
	IntMap                           map[string]int
	FloatMap                         map[string]float32
	Data                             map[string]interface{}
	CSRFToken, Flash, Warning, Error string
}
package template

type FromData struct {
	Action   string
	Method   string
	FormData map[string]string
}
type JsRes struct {
	Url    string
	Method string
	Param  map[string]string
	IsForm bool
}

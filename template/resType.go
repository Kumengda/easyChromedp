package template

type FormData struct {
	Name  string
	Type  string
	Value string
}

type FormDatas struct {
	Action   string
	Method   string
	FormData []FormData
}
type JsRes struct {
	Url          string
	Method       string
	Param        []FormData
	IsForm       bool
	IsFileUpload bool
}

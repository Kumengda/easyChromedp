package template

type FormData struct {
	Enctype string
	Name    string
	Type    string
	Value   string
}

type FormDataset struct {
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

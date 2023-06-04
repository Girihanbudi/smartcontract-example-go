package filesystem

type File struct {
	Filename string      `json:"filename"`
	Contents interface{} `json:"contents"`
}

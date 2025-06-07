package models

import "mime/multipart"

type File struct {
	Name      string //file name with extension
	Parameter string //form parameter
	Extension string //extension such pdf, doc
	Data      *multipart.File
	Path      string //full file path
}

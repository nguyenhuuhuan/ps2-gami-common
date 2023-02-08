package fileupload

import "io"

// Adapter is File Upload adapter.
type Adapter interface {
	Upload(uploadFile io.Reader, uploadFileType, folderName string) (string, error)
	UploadWithName(uploadFile io.Reader, uploadFileType, folderName string, fileName string) (string, error)
	GetURL(url string) (string, error)
}

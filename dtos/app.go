package dtos

// AppError contains code and message of errors
type AppError struct {
	Meta `json:"meta"`
}

// NewAppError build and returns a new Merchant error.
func NewAppError(code int, messages ...string) *AppError {
	msg := ""
	if len(messages) > 0 {
		msg = messages[0]
	}
	return &AppError{
		Meta: Meta{
			Code:    code,
			Message: msg,
		},
	}
}

// Error returns the error as a string.
func (e AppError) Error() string {
	return e.Message
}

// Meta is common meta.
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    *Data  `json:"data,omitempty"`
}

type Data struct {
	AmountLeft int64 `json:"amount_left,omitempty"`
}

// PaginationMeta is meta information with pagination info.
type PaginationMeta struct {
	Meta
	Total int64 `json:"total"`
}

// PaginationMetaV2 is meta information with pagination info v2
type PaginationMetaV2 struct {
	Meta
	Total    int64 `json:"total_items"`
	Page     int   `json:"page_index"`
	PageSize int   `json:"page_size"`
}

type MetaV2 struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Cursor  string `json:"cursor"`
}

type CursorMeta struct {
	*MetaV2
	Total int64 `json:"total"`
}

// NewMeta returns a new meta with message.
func NewMeta(code int, messages ...string) Meta {
	var msg = ""
	if len(messages) > 0 {
		msg = messages[0]
	}
	return Meta{
		Code:    code,
		Message: msg,
	}
}

// GetAppMetaResponse represents response of GetAppMeta.
type GetAppMetaResponse struct {
	Meta Meta `json:"meta"`
}

// BaseCreateEntityResponse represents data in response when create new entity
type BaseCreateEntityResponse struct {
	ID int64 `json:"id,omitempty"`
}

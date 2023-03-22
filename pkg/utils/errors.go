package utils

type GetError struct {
	message    string
	statusCode int
}

type SQLError struct {
	message string
}

func (w *GetError) Error() (string, int) {
	return w.message, w.statusCode
}
func (w *GetError) StatusCode() int {
	return w.statusCode
}
func (w *GetError) ErrorMessage() string {
	return w.message
}

func (w *SQLError) Error() string {
	return w.message
}

//func (w *GetError) Wrap(err int64, info string) *GetError {
//	return &GetError{
//		message:    info,
//		statusCode: err,
//	}
//}

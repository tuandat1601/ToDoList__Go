package common

import "net/http"

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}
func NewErrorResponse(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest ,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}
func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}
 
func (e *AppError) RootError()error{
	if err,ok:=e.RootErr.(*AppError);ok{

		return err.RootError()
	}
	return e.RootErr
}
func (e *AppError) Error() string{
	return e.RootError().Error()
}
func ErrDB(err error)*AppError{
	return NewFullErrorResponse(http.StatusInternalServerError,err,"something went wrong with DB",err.Error(),"DB_ERROR")
}
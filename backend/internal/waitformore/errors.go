package waitformore

type RepositoryError struct {
	HttpStatus int
	ErrorText  string
}

func (r *RepositoryError) GetHttpStatus() int {
	return r.HttpStatus
}

func (r RepositoryError) Error() string {
	return r.ErrorText
}

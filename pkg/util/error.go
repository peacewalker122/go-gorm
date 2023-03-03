package util

import "strings"

type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	var tempErr string
	for _, err := range m.Errors {
		tempErr += err.Error() + ",\n "
	}
	tempErr = strings.TrimSuffix(tempErr, ",\n")
	return tempErr
}

func (m *MultiError) HasError() bool {
	return len(m.Errors) > 0
}

func (m *MultiError) Add(err error) {
	m.Errors = append(m.Errors, err)
}

type UsecaseError struct {
	MultiError
	Status int
}

func (u *UsecaseError) Error() string {
	return u.MultiError.Error()
}

func (u *UsecaseError) ResponseError() (int, string) {
	return u.Status, u.Error()
}

func (u *UsecaseError) HasError() bool {
	return u.MultiError.HasError()
}

func (u *UsecaseError) Add(err error, status int) *UsecaseError {
	u.MultiError.Add(err)
	u.Status = status

	return u
}

func NewUsecaseError() *UsecaseError {
	return &UsecaseError{}
}

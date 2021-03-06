package model

import (
	"fmt"
	"strings"
)

// RepositoryMethod is method of Repository.
type RepositoryMethod string

// methods of Repository.
const (
	RepositoryMethodREAD   RepositoryMethod = "READ"
	RepositoryMethodInsert RepositoryMethod = "INSERT"
	RepositoryMethodUPDATE RepositoryMethod = "UPDATE"
	RepositoryMethodDELETE RepositoryMethod = "DELETE"
	RepositoryMethodLIST   RepositoryMethod = "LIST"
)

// InvalidDataError expresses that given data is invalid.
type InvalidDataError struct {
	BaseErr       error
	DataName      string
	DataValue     interface{}
	InvalidReason string
}

// Error returns error message.
func (e *InvalidDataError) Error() string {
	return fmt.Sprintf("%s, %s", e.DataName, e.InvalidReason)
}

// AlreadyExistError expresses already specified data has existed.
type AlreadyExistError struct {
	BaseErr error
	PropertyName
	PropertyValue interface{}
	DomainModelName
}

// Error returns error message.
func (e *AlreadyExistError) Error() string {
	return fmt.Sprintf("%s, %s, is already exists", e.PropertyName, e.DomainModelName)
}

// RequiredError is not existing necessary value error.
type RequiredError struct {
	BaseErr error
	PropertyName
}

// Error returns error message.
func (e *RequiredError) Error() string {
	return fmt.Sprintf("%s is required", e.PropertyName)
}

// InvalidParamError is inappropriate parameter error。
type InvalidParamError struct {
	BaseErr error
	PropertyName
	PropertyValue interface{}
	InvalidReason string
}

// Error returns error message.
func (e *InvalidParamError) Error() string {
	return fmt.Sprintf("%s, %v, is invalid, %s", e.PropertyName, e.PropertyValue, e.InvalidReason)
}

// InvalidParamsError is inappropriate parameters error。
type InvalidParamsError struct {
	Errors []*InvalidParamError
}

// Error returns error message.
func (e *InvalidParamsError) Error() string {
	length := len(e.Errors)
	messages := make([]string, length, length)
	for i, err := range e.Errors {
		messages[i] = err.Error()
	}
	return strings.Join(messages, ",")
}

// NoSuchDataError is not existing specified data error.
type NoSuchDataError struct {
	BaseErr error
	PropertyName
	PropertyValue interface{}
	DomainModelName
}

// Error returns error message.
func (e *NoSuchDataError) Error() string {
	return fmt.Sprintf("no such data, %s: %v, %s", e.PropertyName, e.PropertyValue, e.DomainModelName)
}

// RepositoryError is Repository error.
type RepositoryError struct {
	BaseErr          error
	RepositoryMethod RepositoryMethod
	DomainModelName
}

// Error returns error message.
func (e *RepositoryError) Error() string {
	return fmt.Sprintf("failed Repository operation, %s, %s", e.RepositoryMethod, e.DomainModelName)
}

// SQLError is SQL error.
type SQLError struct {
	BaseErr                   error
	InvalidReasonForDeveloper InvalidReason
}

// Error returns error message.
func (e *SQLError) Error() string {
	return e.InvalidReasonForDeveloper.String()
}

// AuthenticationErr is Authentication error.
type AuthenticationErr struct {
	BaseErr error
}

// Error returns error message.
func (e *AuthenticationErr) Error() string {
	return "invalid name or password"
}

// OtherServerError is other server error.
type OtherServerError struct {
	BaseErr       error
	InvalidReason string
}

// Error returns error message.
func (e *OtherServerError) Error() string {
	return e.InvalidReason
}

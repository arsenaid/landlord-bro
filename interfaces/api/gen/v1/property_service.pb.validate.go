// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/property_service.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreatePropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePropertyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePropertyRequestMultiError, or nil if none found.
func (m *CreatePropertyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePropertyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProperty()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreatePropertyRequestValidationError{
					field:  "Property",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreatePropertyRequestValidationError{
					field:  "Property",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProperty()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePropertyRequestValidationError{
				field:  "Property",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreatePropertyRequestMultiError(errors)
	}

	return nil
}

// CreatePropertyRequestMultiError is an error wrapping multiple validation
// errors returned by CreatePropertyRequest.ValidateAll() if the designated
// constraints aren't met.
type CreatePropertyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePropertyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePropertyRequestMultiError) AllErrors() []error { return m }

// CreatePropertyRequestValidationError is the validation error returned by
// CreatePropertyRequest.Validate if the designated constraints aren't met.
type CreatePropertyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePropertyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePropertyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePropertyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePropertyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePropertyRequestValidationError) ErrorName() string {
	return "CreatePropertyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePropertyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePropertyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePropertyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePropertyRequestValidationError{}

// Validate checks the field values on UpdatePropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePropertyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePropertyRequestMultiError, or nil if none found.
func (m *UpdatePropertyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePropertyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProperty()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePropertyRequestValidationError{
					field:  "Property",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePropertyRequestValidationError{
					field:  "Property",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProperty()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePropertyRequestValidationError{
				field:  "Property",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePropertyRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePropertyRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePropertyRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdatePropertyRequestMultiError(errors)
	}

	return nil
}

// UpdatePropertyRequestMultiError is an error wrapping multiple validation
// errors returned by UpdatePropertyRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdatePropertyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePropertyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePropertyRequestMultiError) AllErrors() []error { return m }

// UpdatePropertyRequestValidationError is the validation error returned by
// UpdatePropertyRequest.Validate if the designated constraints aren't met.
type UpdatePropertyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePropertyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePropertyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePropertyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePropertyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePropertyRequestValidationError) ErrorName() string {
	return "UpdatePropertyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePropertyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePropertyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePropertyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePropertyRequestValidationError{}

// Validate checks the field values on GetPropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPropertyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPropertyRequestMultiError, or nil if none found.
func (m *GetPropertyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPropertyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PropertyId

	if len(errors) > 0 {
		return GetPropertyRequestMultiError(errors)
	}

	return nil
}

// GetPropertyRequestMultiError is an error wrapping multiple validation errors
// returned by GetPropertyRequest.ValidateAll() if the designated constraints
// aren't met.
type GetPropertyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPropertyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPropertyRequestMultiError) AllErrors() []error { return m }

// GetPropertyRequestValidationError is the validation error returned by
// GetPropertyRequest.Validate if the designated constraints aren't met.
type GetPropertyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPropertyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPropertyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPropertyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPropertyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPropertyRequestValidationError) ErrorName() string {
	return "GetPropertyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPropertyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPropertyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPropertyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPropertyRequestValidationError{}

// Validate checks the field values on ListPropertiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPropertiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPropertiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPropertiesRequestMultiError, or nil if none found.
func (m *ListPropertiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPropertiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for PageToken

	if len(errors) > 0 {
		return ListPropertiesRequestMultiError(errors)
	}

	return nil
}

// ListPropertiesRequestMultiError is an error wrapping multiple validation
// errors returned by ListPropertiesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListPropertiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPropertiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPropertiesRequestMultiError) AllErrors() []error { return m }

// ListPropertiesRequestValidationError is the validation error returned by
// ListPropertiesRequest.Validate if the designated constraints aren't met.
type ListPropertiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPropertiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPropertiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPropertiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPropertiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPropertiesRequestValidationError) ErrorName() string {
	return "ListPropertiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPropertiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPropertiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPropertiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPropertiesRequestValidationError{}

// Validate checks the field values on ListPropertiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPropertiesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPropertiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPropertiesResponseMultiError, or nil if none found.
func (m *ListPropertiesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPropertiesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProperties() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListPropertiesResponseValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPropertiesResponseValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPropertiesResponseValidationError{
					field:  fmt.Sprintf("Properties[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListPropertiesResponseMultiError(errors)
	}

	return nil
}

// ListPropertiesResponseMultiError is an error wrapping multiple validation
// errors returned by ListPropertiesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListPropertiesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPropertiesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPropertiesResponseMultiError) AllErrors() []error { return m }

// ListPropertiesResponseValidationError is the validation error returned by
// ListPropertiesResponse.Validate if the designated constraints aren't met.
type ListPropertiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPropertiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPropertiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPropertiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPropertiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPropertiesResponseValidationError) ErrorName() string {
	return "ListPropertiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPropertiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPropertiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPropertiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPropertiesResponseValidationError{}

// Validate checks the field values on DeletePropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeletePropertyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePropertyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePropertyRequestMultiError, or nil if none found.
func (m *DeletePropertyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePropertyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PropertyId

	if len(errors) > 0 {
		return DeletePropertyRequestMultiError(errors)
	}

	return nil
}

// DeletePropertyRequestMultiError is an error wrapping multiple validation
// errors returned by DeletePropertyRequest.ValidateAll() if the designated
// constraints aren't met.
type DeletePropertyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePropertyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePropertyRequestMultiError) AllErrors() []error { return m }

// DeletePropertyRequestValidationError is the validation error returned by
// DeletePropertyRequest.Validate if the designated constraints aren't met.
type DeletePropertyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePropertyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePropertyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePropertyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePropertyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePropertyRequestValidationError) ErrorName() string {
	return "DeletePropertyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePropertyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePropertyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePropertyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePropertyRequestValidationError{}

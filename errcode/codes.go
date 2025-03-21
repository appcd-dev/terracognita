package errcode

import "errors"

// List of all the error Codes used
var (
	ErrProviderResourceNotSupported  = errors.New("the resource type is not supported")
	ErrProviderResourceNotRead       = errors.New("the resource did not return an ID")
	ErrProviderResourceDoNotMatchTag = errors.New("the resource does not match the required tags")
	ErrProviderResourceFiltered      = errors.New("the resource is filtered")
	ErrProviderResourceAutogenerated = errors.New("the resource is autogenerated and should not be imported")

	ErrCacheKeyNotFound        = errors.New("the key used to search was not found")
	ErrCacheKeyAlreadyExisting = errors.New("the key already exists on the cache")

	ErrWriterRequiredKey      = errors.New("the key is required")
	ErrWriterRequiredValue    = errors.New("the value is required")
	ErrWriterInvalidKey       = errors.New("invalid key")
	ErrWriterInvalidTypeValue = errors.New("invalid type of value")
	ErrWriterAlreadyExistsKey = errors.New("the key already exists")

	ErrFilterTargetsInvalid = errors.New("the filter targets has an invalid format")

	ErrTagInvalidForamt = errors.New("invalid format for tag, the expected format is 'NAME:VALUE' or 'NAME=VALUE'")

	// ErrProviderAPI will be raised when an error occurs provider side while
	// using its APIs (authorization error, unavailable operation, ...)
	ErrProviderAPI = errors.New("error while requesting the provider APIs")

	ErrAccessDenied = errors.New("AccessDenied")
)

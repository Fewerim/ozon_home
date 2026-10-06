package core_errors

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument") // неправильный запрос
	ErrNotFound        = errors.New("not found")        // не найдено
	ErrUnimplemented   = errors.New("not implemented")  // не реализовано
)

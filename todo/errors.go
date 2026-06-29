package todo

import "errors"

var errTaskNotFound = errors.New("task not found")
var errTaskAlreadyExists = errors.New("task already exists")

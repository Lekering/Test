package todo

import "errors"

var ErrorTaskNotFound = errors.New("Not Found")
var ErrorTaskAlreadyExist = errors.New("Already Exist")

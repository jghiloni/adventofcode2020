package errors

import (
	goerr "errors"
)

// ErrUsage will be thrown in Main if command line args are wrong
var ErrUsage = goerr.New(`Usage: adventofcode2020 dayN [part1|part2]`)

// ErrUnimplemented will be returned if the method is unimplemented
var ErrUnimplemented = goerr.New(`unimplemented`)

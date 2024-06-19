package hermeskv

import "errors"

// kv store related errors
var ErrNoKey = errors.New("Key doesn't exist")
var ErrNoNode = errors.New("node not found")

// transaction related errors
var ErrExLockActive = errors.New("exclusive lock active")

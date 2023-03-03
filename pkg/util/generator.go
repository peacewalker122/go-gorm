package util

import "github.com/google/uuid"

type UUIDfn func() uuid.UUID

// UUIDGenerator this will be used to generate UUIDs for entities
// and also dependency injection to create testable code
var UUIDGenerator UUIDfn = uuid.New

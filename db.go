package main

import "github.com/google/uuid"

type Database = map[uuid.UUID]string

var db = Database{}

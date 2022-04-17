package main

type Database = map[string]chan (string)

var db = Database{}

package main

import (
	"time"
)

type ArkEventKind int
const (
	DefaultEvent ArkEventKind = iota
	KillEvent
	TameEvent
	AdminCmdEvent
	JoinEvent
	LeaveEvent
)

type ArkEvent struct {
	Kind ArkEventKind
	Info map[string]string
	Timestamp time.Time
	RawLog string
}


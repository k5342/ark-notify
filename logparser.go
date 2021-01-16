package main

import (
	"strings"
	"errors"
	"time"
	"log"
)

func validateLogLine(logLine string) (error) {
	// 50 is a length of timestamp part
	if (len(logLine) < 50) {
		return errors.New("line is too short")
	}
	if (string(logLine[0]) != "[") {
		return errors.New("wrong character at beginning of timestamp")
	}
	if (string(logLine[24]) != "]") {
		return errors.New("wrong character at ending of timestamp")
	}
	return nil
}

func ParseEventFromLogLine(logLine string) (*ArkEvent, error) {
	ae := ArkEvent{}
	var err error
	if err = validateLogLine(logLine); err != nil {
		return nil, errors.New("log format validation failed. maybe log format was changed?")
	}
	// get timestamp
	log.Println("time.Parse", logLine[1:24])
	var ts time.Time
	if ts, err = time.ParseInLocation("2006.01.02-15.04.05", logLine[1:20], time.Now().Location()); err != nil { // TODO: support milliseconds
		return nil, errors.New("parse timestamp has failed:" + err.Error())
	}
	ae.Timestamp = ts
	// detect event kind
	if (strings.Contains(logLine, " was killed")) {
		ae.Kind = KillEvent
	} else if (strings.Contains(logLine, " Tamed a")) {
		ae.Kind = TameEvent
	} else if (strings.Contains(logLine, " AdminCmd: ")) {
		ae.Kind = AdminCmdEvent
	} else if (strings.Contains(logLine, " joined this ARK")) {
		ae.Kind = JoinEvent
	} else if (strings.Contains(logLine, " left this ARK")) {
		ae.Kind = LeaveEvent
	} else {
		ae.Kind = DefaultEvent
	}
	ae.RawLog = logLine
	return &ae, nil
}


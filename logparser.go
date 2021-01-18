package main

import (
	"strings"
	"errors"
	"time"
	"log"
	"regexp"
	"ark-notify/event"
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

func parseKillEvent(ae *event.ArkEvent, logLine string) {
	r := regexp.MustCompile("(.+ - Lvl \\d+ \\(.+\\)) was killed( by ((an? |)(.+)))?!")
	m := r.FindStringSubmatch(logLine)
	if len(m) == 0 {
		log.Println("parseKillEvent failed: ", logLine)
		return
	}
	ae.Info["Victim"] = m[1];
	if m[5] != "" {
		ae.Info["Assailant"] = m[5];
	}
}

func parseJoinEvent(ae *event.ArkEvent, logLine string) {
	r := regexp.MustCompile("(.+) joined this ARK")
	m := r.FindStringSubmatch(logLine)
	if len(m) == 0 {
		log.Println("parseJoinEvent", logLine)
	}
	ae.Info["Player"] = m[1];
}

func parseLeaveEvent(ae *event.ArkEvent, logLine string) {
	r := regexp.MustCompile("(.+) left this ARK")
	m := r.FindStringSubmatch(logLine)
	if len(m) == 0 {
		log.Println("parseLeaveEvent", logLine)
	}
	ae.Info["Player"] = m[1];
}

func ParseEventFromLogLine(logLine string) (*event.ArkEvent, error) {
	ae := event.ArkEvent{}
	ae.Info = make(map[string]string)
	var err error
	if err = validateLogLine(logLine); err != nil {
		return nil, errors.New("log format validation failed. maybe log format was changed?")
	}
	// get timestamp
	log.Println("time.Parse", logLine[1:24])
	var ts time.Time
	if ts, err = time.Parse("2006.01.02-15.04.05", logLine[1:20]); err != nil { // TODO: support milliseconds
		return nil, errors.New("parse timestamp has failed:" + err.Error())
	}
	ae.Timestamp = ts
	logBody := logLine[50:len(logLine)-1]
	log.Print("logBody:", logBody)
	// detect event kind
	if (strings.Contains(logLine, " was killed")) {
		ae.Kind = event.KillEvent
		parseKillEvent(&ae, logBody)
	} else if (strings.Contains(logLine, " Tamed a")) {
		ae.Kind = event.TameEvent
	} else if (strings.Contains(logLine, " AdminCmd: ")) {
		ae.Kind = event.AdminCmdEvent
	} else if (strings.Contains(logLine, " joined this ARK")) {
		ae.Kind = event.JoinEvent
		parseJoinEvent(&ae, logBody)
	} else if (strings.Contains(logLine, " left this ARK")) {
		ae.Kind = event.LeaveEvent
		parseLeaveEvent(&ae, logBody)
	} else {
		ae.Kind = event.DefaultEvent
	}
	ae.RawLog = logLine
	return &ae, nil
}


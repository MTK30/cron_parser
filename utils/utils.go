package utils

import (
	"strconv"
	"strings"
)

const (
	Asterix             = "*"
	BackwardSlash       = "/"
	Comma               = ","
	Hypen               = "-"
	EmptyString         = ""
	CronExpressionRegeX = `^(\*|(\d+|\*\/\d+|\d+-\d+|\d+(,\d+)*)) (\*|(\d+|\*\/\d+|\d+-\d+|\d+(,\d+)*)) (\*|(\d+|\*\/\d+|\d+-\d+|\d+(,\d+)*)) (\*|(\d+|\*\/\d+|\d+-\d+|\d+(,\d+)*)) (\*|(\d+|\*\/\d+|\d+-\d+|\d+(,\d+)*)) ([^\s]+)$` //taken from internet
	Minutes             = "minutes"
	Hour                = "hour"
	DayOfMonth          = "DayOfMonth"
	Month               = "Month"
	DayOfWeek           = "DayOfWeek"
	Command             = "Command"
)

func GenerateSequence(start, end, interval int) string {
	var result []string
	for i := start; i <= end; i += interval {
		result = append(result, strconv.Itoa(i))
	}
	return strings.Join(result, " ")
}

package main

import (
	"log"
	"testing"

	"github.com/MTK30/cron_parser/parser"
	"github.com/MTK30/cron_parser/utils"
)

func TestValidMinute(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"*/15 0 1,15 * 1-5 /cmd/abc", "0 15 30 45"},
		{"* * * * * /cmd/abc", utils.GenerateSequence(0, 59, 1)},
		{"*/20 0 1,15 * 1-5 /cmd/abc", utils.GenerateSequence(0, 59, 20)},
		{"20,30 0 1,15 * 1-5 /cmd/abc", "20 30"},
		{"15 0 1,15 * 1-5 /cmd/abc", "15"},
	}
	for _, tc := range testCases {
		parse, err := parser.GetParserInstance(tc.input)
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		err = parse.MinuteParser()
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		log.Println("parse.ParserTypeIns.Minute : ", parse.ParserTypeIns.Minute)
		if parse.ParserTypeIns.Minute != tc.expected {
			t.Errorf("isValidCronExpression(%q) expected : %v, acutal : %v", tc.input, parse.ParserTypeIns.Minute, tc.expected)
		}
	}
}

func TestValidHours(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0 */15 1,15 * 1-5 /cmd/abc", utils.GenerateSequence(0, 23, 15)},
		{"* * * * * /cmd/abc", utils.GenerateSequence(0, 23, 1)},
		{"0 */20 1,15 * 1-5 /cmd/abc", utils.GenerateSequence(0, 23, 20)},
		{"0 15 1,15 * 1-5 /cmd/abc", "15"},
		{"0 3,6 1,15 * 1-5 /cmd/abc", "3 6"},
	}
	for _, tc := range testCases {
		parse, err := parser.GetParserInstance(tc.input)
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		// log.Println("len : ", len(parse.Words))
		err = parse.HourParser()
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		log.Println("parse.ParserTypeIns.Hour : ", parse.ParserTypeIns.Hours)
		if parse.ParserTypeIns.Hours != tc.expected {
			t.Errorf("isValidCronExpression(%q) expected : %v, acutal : %v", tc.input, parse.ParserTypeIns.Minute, tc.expected)
		}
	}
}

func TestValidMonths(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0 */15 1,15 * 1-5 /cmd/abc", utils.GenerateSequence(1, 12, 1)},
		{"* * * 1,12 1,12 /cmd/abc", "1 12"},
		{"0 */20 1,15 */12 1-5 /cmd/abc", utils.GenerateSequence(1, 12, 12)},
		{"0 15 1,15 1-5 1-5 /cmd/abc", "1 2 3 4 5"},
		{"0 * 1,15 3,6 1-5 /cmd/abc", "3 6"},
	}
	for _, tc := range testCases {
		parse, err := parser.GetParserInstance(tc.input)
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		// log.Println("len : ", len(parse.Words))
		err = parse.MonthParser()
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		log.Println("parse.ParserTypeIns.Hour : ", parse.ParserTypeIns.Month)
		if parse.ParserTypeIns.Month != tc.expected {
			t.Errorf("isValidCronExpression(%q) expected : %v, acutal : %v", tc.input, parse.ParserTypeIns.Minute, tc.expected)
		}
	}
}

func TestValidWeek(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0 */15 1,15 * 1-5 /cmd/abc", utils.GenerateSequence(1, 5, 1)},
		{"* * * 1,12 1,6 /cmd/abc", "1 6"},
		{"0 */20 1,15 */12 */7 /cmd/abc", utils.GenerateSequence(1, 7, 7)},
		{"0 15 1,15 1-5 1-5 /cmd/abc", "1 2 3 4 5"},
		{"0 * 1,15 3,6 */1 /cmd/abc", utils.GenerateSequence(1, 7, 1)},
	}
	for _, tc := range testCases {
		parse, err := parser.GetParserInstance(tc.input)
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		// log.Println("len : ", len(parse.Words))
		err = parse.DayOfWeek()
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		log.Println("parse.ParserTypeIns.Hour : ", parse.ParserTypeIns.DayOfWeek)
		if parse.ParserTypeIns.DayOfWeek != tc.expected {
			t.Errorf("isValidCronExpression(%q) expected : %v, acutal : %v", tc.input, parse.ParserTypeIns.Minute, tc.expected)
		}
	}
}

func TestValidDayOfMonth(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0 */15 1,15 * 1-5 /cmd/abc", "1 15"},
		{"* * * 1,12 1,6 /cmd/abc", utils.GenerateSequence(1, 30, 1)},
		{"0 */20 */7 */12 */7 /cmd/abc", utils.GenerateSequence(1, 30, 7)},
		{"0 15 1-5 1-5 1-5 /cmd/abc", "1 2 3 4 5"},
		{"0 * 1-15 3,6 */1 /cmd/abc", utils.GenerateSequence(1, 15, 1)},
	}
	for _, tc := range testCases {
		parse, err := parser.GetParserInstance(tc.input)
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		// log.Println("len : ", len(parse.Words))
		err = parse.DayOfMonthParser()
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		log.Println("parse.ParserTypeIns.Hour : ", parse.ParserTypeIns.DayOfMonth)
		if parse.ParserTypeIns.DayOfMonth != tc.expected {
			t.Errorf("isValidCronExpression(%q) expected : %v, acutal : %v", tc.input, parse.ParserTypeIns.Minute, tc.expected)
		}
	}
}

func TestCommand(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"0 */15 1,15 * 1-5 /cmd/abc", "/cmd/abc"},
	}
	for _, tc := range testCases {
		parse, err := parser.GetParserInstance(tc.input)
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		// log.Println("len : ", len(parse.Words))
		err = parse.CommandParser()
		if err != nil {
			t.Errorf("isValidCronExpression(%q) with err : %v", tc.input, err.Error())
		}
		log.Println("parse.ParserTypeIns.Command : ", parse.ParserTypeIns.Command)
		if parse.ParserTypeIns.Command != tc.expected {
			t.Errorf("isValidCronExpression(%q) expected : %v, acutal : %v", tc.input, parse.ParserTypeIns.Minute, tc.expected)
		}
	}
}

func TestParseCronExpression(t *testing.T) {
	testCases := []struct {
		input   string
		outputs *parser.ParsedType
	}{
		{
			input: "*/15 0 1,15 * 1-5 /usr/bin/find",
			outputs: &parser.ParsedType{
				Minute:     utils.GenerateSequence(0, 59, 15),
				Hours:      "0",
				DayOfMonth: "1 15",
				Month:      utils.GenerateSequence(1, 12, 1),
				DayOfWeek:  utils.GenerateSequence(1, 5, 1),
				Command:    "/usr/bin/find",
			},
		},
		{
			input: "*/10 * * * * /usr/bin/command",
			outputs: &parser.ParsedType{
				Minute:     utils.GenerateSequence(0, 59, 10),
				Hours:      utils.GenerateSequence(0, 23, 1),
				DayOfMonth: utils.GenerateSequence(1, 30, 1),
				Month:      utils.GenerateSequence(1, 12, 1),
				DayOfWeek:  utils.GenerateSequence(1, 7, 1),
				Command:    "/usr/bin/command",
			},
		},
		{
			input: "0 12 * * 1-5 /usr/bin/restart-service",
			outputs: &parser.ParsedType{
				Minute:     "0",
				Hours:      "12",
				DayOfMonth: utils.GenerateSequence(1, 30, 1),
				Month:      utils.GenerateSequence(1, 12, 1),
				DayOfWeek:  utils.GenerateSequence(1, 5, 1),
				Command:    "/usr/bin/restart-service",
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {

		parse, _ := parser.GetParserInstance(tc.input)
		parse.Handler()
		var isInValid bool
		if tc.outputs.Month != parse.ParserTypeIns.Month || tc.outputs.DayOfMonth != parse.ParserTypeIns.DayOfMonth || tc.outputs.Minute != parse.ParserTypeIns.Minute {
			isInValid = true
		}
		if tc.outputs.Hours != parse.ParserTypeIns.Hours || tc.outputs.DayOfWeek != parse.ParserTypeIns.DayOfWeek || tc.outputs.Command != parse.ParserTypeIns.Command {
			isInValid = true
		}
		if isInValid {
			t.Errorf("Error!!")
		}

	}
}

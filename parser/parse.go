package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MTK30/cron_parser/utils"
)

type ParsedType struct {
	Minute     string
	Hours      string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Command    string
}

type ParserInterface interface {
	MinuteParser() error
	HourParser() error
	DayOfMonthParser() error
	MonthParser() error
	DayOfWeek() error
	CommandParser() error
}

type Parser struct {
	Words         []string
	ParserTypeIns *ParsedType
}

var parse Parser

func GetParserInstance(cronString string) (Parser, error) {
	strArry, err := splitAndStore(cronString)
	if err != nil {
		return Parser{}, err
	}
	parse = Parser{Words: strArry, ParserTypeIns: &ParsedType{}}
	return parse, nil
}

func (parser Parser) Handler() error {
	if err := parse.MinuteParser(); err != nil {
		return err
	}
	if err := parse.HourParser(); err != nil {
		return err
	}
	if err := parse.DayOfMonthParser(); err != nil {
		return err
	}
	if err := parse.MonthParser(); err != nil {
		return err
	}
	if err := parse.DayOfWeek(); err != nil {
		return err
	}
	if err := parse.CommandParser(); err != nil {
		return err
	}
	return nil
}

func (p Parser) String() string {
	return fmt.Sprintf("minute\t %s \nhour\t %s\nday of month\t %s\nmonth \t %s\nday of week\t %s\ncommand\t %s\n", p.ParserTypeIns.Minute, p.ParserTypeIns.Hours, p.ParserTypeIns.DayOfMonth, p.ParserTypeIns.Month, p.ParserTypeIns.DayOfWeek, p.ParserTypeIns.Command)
}

func splitAndStore(cronString string) ([]string, error) {
	if cronString == utils.EmptyString {
		return nil, errors.New("invalid length cronString")
	}
	re := regexp.MustCompile(utils.CronExpressionRegeX)
	if !re.MatchString(cronString) {
		return nil, errors.New("invalid char in cronString")
	}
	strArry := strings.Split(cronString, " ")
	if len(strArry) != 6 || len(strArry) == 0 {
		return nil, errors.New("invalid parser attributes, valid attribute format : \t `Min hrs dayOfMont Month dayOfWeek Cmd`")
	}
	return strArry, nil
}

func (p *Parser) HourParser() error {
	hoursString := p.Words[1]
	err := p.handle(hoursString, utils.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) DayOfMonthParser() error {
	dayOfMonth := p.Words[2]
	err := p.handle(dayOfMonth, utils.DayOfMonth)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) MinuteParser() error {
	minutesString := p.Words[0]
	err := p.handle(minutesString, utils.Minutes)
	if err != nil {
		return err
	}
	return nil
}
func (p *Parser) MonthParser() error {
	month := p.Words[3]
	err := p.handle(month, utils.Month)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) DayOfWeek() error {
	week := p.Words[4]
	err := p.handle(week, utils.DayOfWeek)
	if err != nil {
		return err
	}
	return nil
}

func (p *Parser) CommandParser() error {
	week := p.Words[5]
	p.setToParsedType(utils.Command, week)
	return nil
}

func (p *Parser) handle(parserString, keyType string) error {
	var st string
	if strings.Contains(parserString, "-") && strings.Contains(parserString, "/") {
		return errors.New("wrong value for min string")
	}
	if parserString == utils.Asterix {
		in := 0
		if keyType == utils.DayOfMonth || keyType == utils.DayOfWeek || keyType == utils.Month {
			in = 1
		}
		st = getSequence(in, 1, keyType)
	} else if strings.Contains(parserString, utils.Hypen) {
		inp := strings.Split(parserString, utils.Hypen)
		s, err := getStringForHypen(inp, keyType)
		if err != nil {
			return err
		}
		st = s
	} else if strings.Contains(parserString, utils.BackwardSlash) {
		inp := strings.Split(parserString, utils.BackwardSlash)
		s, err := getStringForBackwardSlash(inp, keyType)
		if err != nil {
			return err
		}
		st = s
	} else if strings.Contains(parserString, utils.Comma) {
		inp := strings.Split(parserString, utils.Comma)
		s, err := getStringForComma(inp, keyType)
		if err != nil {
			return err
		}
		st = s
	} else {
		minute, err := strconv.Atoi(parserString)
		if err != nil || minute < 0 || minute > 59 {
			return errors.New("Invalid input")
		}
		st = parserString
	}
	p.setToParsedType(keyType, st)
	return nil
}

func (p *Parser) setToParsedType(keyType, s string) {
	switch keyType {
	case utils.Minutes:
		p.ParserTypeIns.Minute = s
	case utils.Hour:
		p.ParserTypeIns.Hours = s
	case utils.DayOfMonth:
		p.ParserTypeIns.DayOfMonth = s
	case utils.DayOfWeek:
		p.ParserTypeIns.DayOfWeek = s
	case utils.Month:
		p.ParserTypeIns.Month = s
	case utils.Command:
		p.ParserTypeIns.Command = s
	}
}

//private methods

func getStringForComma(inp []string, keyType string) (string, error) {
	for _, in := range inp {
		i, _ := strconv.Atoi(in)
		if keyType == utils.Hour && i > 23 {
			return utils.EmptyString, errors.New("invalid hour value")
		} else if keyType == utils.Minutes && i > 59 {
			return utils.EmptyString, errors.New("invalid min value")
		} else if keyType == utils.DayOfMonth && i > 30 {
			return utils.EmptyString, errors.New("invalid day Of Month value")
		} else if keyType == utils.DayOfWeek && i > 6 {
			return utils.EmptyString, errors.New("invalid day Of week value")
		} else if keyType == utils.Month && i > 12 {
			return utils.EmptyString, errors.New("invalid Month value")
		}

	}
	return strings.Join(inp, " "), nil
}

func getStringForHypen(inp []string, keyType string) (string, error) {
	if len(inp) != 2 {
		return utils.EmptyString, errors.New("invalid input")
	}
	startPoint, err := strconv.Atoi(inp[0])
	if err != nil {
		return utils.EmptyString, err
	}
	endPoint, err := strconv.Atoi(inp[1])
	if err := isValidEndPoint(endPoint, keyType); err != nil {
		return utils.EmptyString, err
	}
	if err != nil {
		return utils.EmptyString, err
	}
	s := utils.GenerateSequence(startPoint, endPoint, 1)
	return s, nil
}

func getStringForBackwardSlash(inp []string, keyType string) (string, error) {
	if len(inp) != 2 {
		return utils.EmptyString, errors.New("invalid input")
	}
	var s string
	interval, err := strconv.Atoi(inp[1])
	if err != nil {
		return utils.EmptyString, err
	}
	if inp[0] == "*" {
		in := 0
		if keyType == utils.DayOfMonth || keyType == utils.DayOfWeek || keyType == utils.Month {
			in = 1
		}
		s = getSequence(in, interval, keyType)
		return s, nil
	}
	startPoint, err := strconv.Atoi(inp[0])
	if err != nil {
		return utils.EmptyString, err
	}
	s = getSequence(startPoint, interval, keyType)
	return s, nil
}

func getSequence(strt, interval int, keyType string) string {
	var s string
	if keyType == utils.Minutes {
		s = utils.GenerateSequence(strt, 59, interval)
	} else if keyType == utils.Hour {
		s = utils.GenerateSequence(strt, 23, interval)
	} else if keyType == utils.DayOfMonth {
		s = utils.GenerateSequence(strt, 30, interval)
	} else if keyType == utils.Month {
		s = utils.GenerateSequence(strt, 12, interval)
	} else if keyType == utils.DayOfWeek {
		s = utils.GenerateSequence(strt, 7, interval)
	}
	return s
}

func isValidEndPoint(endpoint int, keyType string) error {
	switch keyType {
	case utils.Hour:
		if endpoint >= 24 {
			return errors.New("hours cant be more than 24")
		}
	case utils.Minutes:
		if endpoint >= 59 {
			return errors.New("hours cant be more than 59")
		}
	case utils.DayOfMonth:
		if endpoint > 30 {
			return errors.New("day of month cant be more than 30")
		}
	case utils.Month:
		if endpoint > 12 {
			return errors.New("month number cant be more than 12")
		}
	case utils.DayOfWeek:
		if endpoint > 7 {
			return errors.New("day of wekk cant be more or equal to 7")
		}
	}
	return nil
}

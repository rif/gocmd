package main

import (
	"log"
	"regexp"
)

var str = `
Tag,RatesTag,TimingProfile,Weight
STANDARD,RT_STANDARD,WORKDAYS_00,10
STANDARD,RT_STD_WEEKEND,WORKDAYS_18,10
STANDARD,RT_STD_WEEKEND,WEEKENDS,10
PREMIUM,RT_STD_WEEKEND,WEEKENDS,10
DEFAULT,RT_DEFAULT,WORKDAYS_00,10
EVENING,P1,WORKDAYS_00,10
EVENING,P2,WORKDAYS_18,10
EVENING,P2,WEEKENDS,10
`

func main() {
	re := regexp.MustCompile(`(?m)(?:\w+\s*,\s*){3}\d+$`)
	lines := re.FindAllStringSubmatch(str, -1)
	log.Print(len(lines))
}

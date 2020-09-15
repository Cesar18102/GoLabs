package Tasks

import (
	"fmt"
)

type daysInMonth struct {
	month string;
	days int;
}

var DAYS_IN_MONTHES = [MONTH_COUNT]daysInMonth {
	{ "January", 31 },
	{ "February", 28 },
	{ "March", 31 },
	{ "April", 30 },
	{ "May", 31 },
	{ "June", 30 },
	{ "July", 31 },
	{ "August", 31 },
	{ "September", 30 },
	{ "October", 31 },
	{ "November", 30 },
	{ "December", 31 },
}

var SEASONS = [SEASON_COUNT]string{
	"WINTER",
	"SPRING",
	"SUMMER",
	"AUTUMN",
}

const SEASON_COUNT = 4;
const MONTH_COUNT = 12;
const MONTHES_PER_SEASON = 3;
const FIRST_SEASON_MONTH_INDENT = 11;

func Task3() {
	season := int(randval(0, SEASON_COUNT)) + 1;

	result := monthMapBySeason(season);
	fmt.Printf("%s: %v", SEASONS[season - 1], result);
}

func monthMapBySeason(season int) [MONTHES_PER_SEASON]daysInMonth {
	firstMonth := (FIRST_SEASON_MONTH_INDENT + (season - 1) * MONTHES_PER_SEASON) % MONTH_COUNT;
	result := [MONTHES_PER_SEASON]daysInMonth { };

	for i := 0; i < MONTHES_PER_SEASON; i++ {
		monthNumber := (firstMonth + i) % MONTH_COUNT;
		result[i] = DAYS_IN_MONTHES[monthNumber];
	}

	return result;
}

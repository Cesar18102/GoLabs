package Tasks

import "fmt"

func Task3()  {
	monthNumber := randInt(0, 12);
	month, nextMonthNumber := monthName(monthNumber), nextMonthNumber(monthNumber);
	nextMonth := monthName(nextMonthNumber);
	fmt.Printf("The next month after %v(%d) is %v(%d)\n", month, monthNumber + 1, nextMonth, nextMonthNumber + 1);
}

const MONTH_COUNT = 12;
func nextMonthNumber(number int) int {
	return (number + 1) % MONTH_COUNT;
}

func monthName(number int) string {
	switch number {
		case 0: return "January";
		case 1: return "February";
		case 2: return "March";
		case 3: return "April";
		case 4: return "May";
		case 5: return "June";
		case 6: return "July";
		case 7: return "August";
		case 8: return "September";
		case 9: return "October";
		case 10: return "November";
		case 11: return "December";
		default: return "UNKNOWN";
	}
}
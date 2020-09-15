package Tasks

import (
	"fmt"
	"unicode/utf8"
)

type employee struct {
	name string;
	salary float64;
	stage int;
}

func Task8() {
	var emps []employee = []employee {
		{ "Олег", 5000, 3 },
		{ "Леонид", 1500, 1 },
		{ "Антон", 5000, 3 },
		{ "Лев", 500, 0 },
	};
	fmt.Printf("Emploees list: %v\n", emps);

	lemps := where(emps, func(emp employee) bool {
		rune, _ := utf8.DecodeRuneInString(emp.name);
		return rune == 'Л';
	});
	fmt.Printf("Emploees whose name starts with 'Л': %v\n", lemps);

	var toadd []employee = []employee {
		{ "Сергей", 10000, 5 },
		{ "Андрей", 7000, 4 },
		{ "Максим", 2500, 2 },
	};

	for _, emp := range toadd {
		fmt.Printf("Adding %v\n", emp);
		emps = prepend(emps, emp);
		fmt.Printf("Changed emploees list: %v\n", emps);
	}
}

func prepend(emps []employee, emp employee) []employee {
	result := make([]employee, len(emps) + 1);
	result[0] = emp;
	copy(result[1:], emps);
	return result;
}

func where(emps []employee, predicate func(emp employee) bool) []employee {
	var found []employee;
	for _, emp := range emps {
		if predicate(emp) {
			found = append(found, emp);
		}
	}
	return found;
}

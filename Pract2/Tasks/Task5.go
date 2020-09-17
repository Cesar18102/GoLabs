package Tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type worker struct {
	name string;
	pos string;
	year int;
}

var scanner = bufio.NewScanner(os.Stdin)

func Task5() {
	workers := make([]worker, 10);
	for i := 0; i < 10; i++ {
		workers[i] = readWorker();
	}

	fmt.Printf("Enter minimum experience: ");
	minexp := readInt(scanner);
	experienced := whereWorker(workers, func(w worker) bool {
		return time.Now().Year() - w.year > minexp;
	});

	if len(experienced) != 0 {
		fmt.Printf("Workers with experience of %d+ years are: ", minexp);
		for _, v := range experienced {
			fmt.Printf("%v; ", v.name);
		}
	} else {
		fmt.Printf("There is no workers with experience of %d+ years", minexp);
	}
}

func whereWorker(workers []worker, predicate func(w worker) bool) []worker {
	found := make([]worker, 0);
	for _, v := range workers {
		if predicate(v) {
			found = append(found, v);
		}
	}
	return found;
}

func readWorker() worker {
	fmt.Print("Input worker's name: ");
	name := readText(scanner);

	fmt.Print("Input worker's position: ");
	pos := readText(scanner);

	fmt.Print("Input worker's hire year: ");
	year := readInt(scanner);

	return worker { name, pos, year };
}

func readText(scanner *bufio.Scanner) string {
	scanner.Scan();
	return scanner.Text();
}

func readInt(scanner *bufio.Scanner) int {
	num, _ := strconv.Atoi(readText(scanner));
	return num;
}
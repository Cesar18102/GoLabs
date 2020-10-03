package Tasks

import "fmt"

type Fraction struct {
	numerator int;
	denominator int;
}

func gcd(a int, b int) int {
	if a == 0 || b == 0 {
		return a + b;
	}

	if a > b {
		return gcd(a % b, b);
	} else {
		return gcd(a, b % a);
	}
}

func plus(fraction *Fraction, otherFraction Fraction) {
	commonDenominator := gcd(fraction.denominator, otherFraction.denominator);
	fraction.numerator = fraction.numerator * commonDenominator / fraction.denominator +
		otherFraction.numerator * commonDenominator / otherFraction.denominator;
	fraction.denominator = commonDenominator;
}

func negate(fraction *Fraction) {
	fraction.numerator = -fraction.numerator;
}

func minus(fraction *Fraction, otherFraction Fraction)  {
	negate(fraction);
	plus(fraction, otherFraction);
	negate(fraction);
}

func mult(fraction *Fraction, otherFraction Fraction) {
	fraction.numerator *= otherFraction.numerator;
	fraction.denominator *= otherFraction.denominator;
}

func Task2() {
	a, b, c, e, f := readIntNoError("Input A: "), readIntNoError("Input B: "), readIntNoError("Input C: "),
		readIntNoError("Input E: "),readIntNoError("Input F: ");

	abc := Fraction { a, b + c };
	aca := Fraction { c, a - c };
	ef := Fraction { e, f };

	minus(&abc, aca);
	mult(&abc, ef);

	fmt.Printf("(a / (b + c) - c / (a - c)) * e / f = %d / %d", abc.numerator, abc.denominator);
}

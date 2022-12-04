package percent

import "strconv"

type PCent8 uint8
type PCent16 uint16

func (p PCent8) String() string {
	return strconv.Itoa(int(p))
}

// Percent - calculate what %[number1] of [number2] is.
// ex. 25% of 200 is 50
func (p PCent8) Percent(percent int, all int) float64 {
	return Percent(percent, all)
}

// PercentFloat - calculate what %[number1] of [number2] is.
// ex. 25% of 200 is 50
func (p PCent8) PercentFloat(percent float64, all float64) float64 {
	return PercentFloat(percent, all)
}

// PercentOf - calculate what percent [number1] is of [number2].
// ex. 300 is 12.5% of 2400
func (p PCent8) PercentOf(part int, total int) float64 {
	return PercentOf(part, total)
}

// PercentOfFloat - calculate what percent [number1] is of [number2].
// ex. 300 is 12.5% of 2400
func (p PCent8) PercentOfFloat(part float64, total float64) float64 {
	return PercentOfFloat(part, total)
}

// Change - calculate the percent increase/decrease from two numbers.
// ex. 60 is a 200.0% increase from 20
func (p PCent8) Change(before int, after int) float64 {
	return Change(before, after)
}

// ChangeFloat - calculate the percent increase/decrease from two numbers.
// ex. 60.0 is a 200.0% increase from 20.0
func (p PCent8) ChangeFloat(before float64, after float64) float64 {
	return ChangeFloat(before, after)
}

func (p PCent16) String() string {
	return strconv.Itoa(int(p))
}

// Percent - calculate what %[number1] of [number2] is.
// ex. 25% of 200 is 50
func (p PCent16) Percent(percent int, all int) float64 {
	return Percent(percent, all)
}

// PercentFloat - calculate what %[number1] of [number2] is.
// ex. 25% of 200 is 50
func (p PCent16) PercentFloat(percent float64, all float64) float64 {
	return PercentFloat(percent, all)
}

// PercentOf - calculate what percent [number1] is of [number2].
// ex. 300 is 12.5% of 2400
func (p PCent16) PercentOf(part int, total int) float64 {
	return PercentOf(part, total)
}

// PercentOfFloat - calculate what percent [number1] is of [number2].
// ex. 300 is 12.5% of 2400
func (p PCent16) PercentOfFloat(part float64, total float64) float64 {
	return PercentOfFloat(part, total)
}

// Change - calculate the percent increase/decrease from two numbers.
// ex. 60 is a 200.0% increase from 20
func (p PCent16) Change(before int, after int) float64 {
	return Change(before, after)
}

// ChangeFloat - calculate the percent increase/decrease from two numbers.
// ex. 60.0 is a 200.0% increase from 20.0
func (p PCent16) ChangeFloat(before float64, after float64) float64 {
	return ChangeFloat(before, after)
}

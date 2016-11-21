package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	x string
	y string
)

func init() {
	flag.StringVar(&x, "X", "0", "First number")
	flag.StringVar(&y, "Y", "0", "Second number")
	flag.Parse()
}

var Icounter uint64

func main() {
	Icounter = 0
	x := "1685287499328328297814655639278583667919355849391453456921116729"
	y := "7114192848577754587969744626558571536728983167954552999895348492"
	z, n := Mul(x, y, `12`)
	fmt.Printf(" %s -> %d\n", z, n)
}

func Add(x, y string) string {
	sign := ""

	if strings.HasPrefix(x, `-`) && strings.HasPrefix(y, `-`) {
		sign = "-"
		x = strings.TrimLeft(x, "-")
		y = strings.TrimLeft(y, "-")
	} else if strings.HasPrefix(x, `-`) {
		return Sub(y, strings.TrimLeft(x, "-"))
	} else if strings.HasPrefix(y, `-`) {
		return Sub(x, strings.TrimLeft(y, "-"))
	}

	if len(x) == 1 || len(y) == 1 {
		xn, _ := strconv.Atoi(x)
		yn, _ := strconv.Atoi(y)
		return sign + strconv.Itoa(xn+yn)
	}

	carry := 0
	result := ""
	xy := ""
	n, xs, ys := SameSize(x, y)

	for i := n - 1; i > -1; i-- {
		xl, _ := strconv.Atoi(string(xs[i]))
		yl, _ := strconv.Atoi(string(ys[i]))
		z := strconv.Itoa(xl + yl + carry)
		if len(z) > 1 {
			carry = 1
			xy = string(z[1])
		} else {
			carry = 0
			xy = string(z[0])
		}
		result = xy + result
	}

	if carry > 0 {
		result = "1" + result
	}
	trimmed := strings.TrimLeft(result, "0")
	if len(trimmed) == 0 {
		return `0`
	}
	return sign + trimmed
}

func Sub(x, y string) string {
	if strings.HasPrefix(x, `-`) && strings.HasPrefix(y, `-`) {
		x = strings.TrimLeft(x, "-")
		y = strings.TrimLeft(y, "-")
		return Sub(y, x)
	} else if strings.HasPrefix(x, `-`) {
		return "-" + Add(strings.TrimLeft(x, "-"), y)
	} else if strings.HasPrefix(y, `-`) {
		return Add(x, strings.TrimLeft(y, "-"))
	}

	if len(x) == 1 || len(y) == 1 {
		xn, _ := strconv.Atoi(x)
		yn, _ := strconv.Atoi(y)
		return strconv.Itoa(xn - yn)
	}

	carry := 0
	result := ""
	n, xs, ys := SameSize(x, y)
	z := ""
	for i := n - 1; i > -1; i-- {
		xl, _ := strconv.Atoi(string(xs[i]))
		yl, _ := strconv.Atoi(string(ys[i]))
		if yl > xl {
			z = strconv.Itoa(10 + xl - yl - carry)
			carry = 1
		} else {
			z = strconv.Itoa(xl - yl - carry)
			if z == "-1" {
				z = "9"
				carry = 1
			} else {
				carry = 0
			}
		}
		result = string(z[0]) + result
	}
	if carry > 0 {
		return `-` + Sub(y, x)
	}

	trimmed := strings.TrimLeft(result, "0")
	if len(trimmed) == 0 {
		return `0`
	}

	return trimmed
}

func Mul(x, y, inspected string) (string, uint64) {
	if strings.HasPrefix(x, `-`) {
		x = strings.TrimLeft(x, "-")
	} else if strings.HasPrefix(y, `-`) {
		y = strings.TrimLeft(y, "-")
	}

	if len(x) == 1 && len(y) == 1 {
		xn, _ := strconv.Atoi(x)
		yn, _ := strconv.Atoi(y)
		return strconv.Itoa(xn * yn), 0
	}

	n, xs, ys := SameSize(x, y)
	m := int(math.Ceil(float64(n) / 2.0))

	x1, x2 := SplitString(xs)
	y1, y2 := SplitString(ys)

	p1, _ := Mul(x1, y1, inspected)
	p2, _ := Mul(x2, y2, inspected)
	p3, _ := Mul(Add(x1, x2), Add(y1, y2), inspected)

	a := p1 + strings.Repeat("0", 2*m)
	a = Add(a, p2)

	b := Sub(Sub(p3, p1), p2)

	if b == inspected {
		Icounter = Icounter + 1
	}

	b = b + strings.Repeat("0", m)

	result := Add(a, b)
	return result, Icounter
}

func SplitString(s string) (sl, sr string) {
	sb := []byte(s)
	n := len(sb)
	return string(sb[:n/2]), string(sb[n/2:])
}

func SameSize(x, y string) (int, string, string) {
	n := len(x)
	if n > len(y) {
		y = strings.Repeat("0", n-len(y)) + y
	} else if n < len(y) {
		n = len(y)
		x = strings.Repeat("0", n-len(x)) + x
	}
	return n, x, y
}

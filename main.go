package main

import (
	"flag"
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

func main() {
}

func Add(x, y string) string {
	carry := 0
	result := ""
	xy := ""
	n := len(x)
	xa := []byte(x)
	ya := []byte(y)
	for i, _ := range x {
		xl, _ := strconv.Atoi(string(xa[n-i-1]))
		yl, _ := strconv.Atoi(string(ya[n-i-1]))
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
	return result
}

func Sub(x, y string) string {
	carry := 0
	result := ""
	n := len(x)
	xa := []byte(x)
	ya := []byte(y)
	z := ""
	for i := range x {
		xl, _ := strconv.Atoi(string(xa[n-i-1]))
		yl, _ := strconv.Atoi(string(ya[n-i-1]))
		if yl > xl {
			z = strconv.Itoa(10 + xl - yl - carry)
			carry = 1
		} else {
			z = strconv.Itoa(xl - yl - carry)
			carry = 0
		}
		result = string(z[0]) + result
	}
	if carry > 0 {
		carry = 0
		result = ""
		for i := range x {
			xl, _ := strconv.Atoi(string(xa[n-i-1]))
			yl, _ := strconv.Atoi(string(ya[n-i-1]))
			if xl > yl {
				z = strconv.Itoa(10 + yl - xl - carry)
				carry = 1
			} else {
				z = strconv.Itoa(yl - xl - carry)
				carry = 0
			}
			result = string(z[0]) + result
		}
		result = "-" + result
	}
	return result
}

func Mul(x, y string) string {
	n := len(x) // Let's assume, both numbers has the same length
	if n < 2 {
		xn, _ := strconv.Atoi(x)
		yn, _ := strconv.Atoi(y)
		return strconv.Itoa(xn * yn)
	}

	m := n / 2
	x1, x2 := SplitString(x)
	y1, y2 := SplitString(y)

	z2 := Mul(x1, y1)
	z1 := Mul(Add(x1, x2), Add(y1, y2))
	z0 := Mul(x2, y2)

	r0 := z2 + strings.Repeat("0", n)
	r1 := Sub(Sub(z1, z2), z0) + strings.Repeat("0", m)

	return r0 + r1 + z0
}

func SplitString(s string) (sl, sr string) {
	sb := []byte(s)
	n := len(sb)
	return string(sb[:n/2]), string(sb[n/2:])
}

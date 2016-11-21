package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	f, err := os.Open("add_cases.csv")
	if err != nil {
		t.Error("Can't open test file")
	}
	defer f.Close()

	r := csv.NewReader(f)
	var v string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("Can't read test file (may be not a csv?)")
		}

		v = Add(record[0], record[1])
		if v != record[2] {
			t.Error("Expected ", record[2], " got ", v)
		}

	}
}

func TestSub(t *testing.T) {
	f, err := os.Open("sub_cases.csv")
	if err != nil {
		t.Error("Can't open test file")
	}
	defer f.Close()

	r := csv.NewReader(f)
	var v string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("Can't read test file (may be not a csv?)")
		}

		v = Sub(record[0], record[1])
		if v != record[2] {
			t.Error("Expected ", record[2], " got ", v)
		}

	}
}

func TestSplitString(t *testing.T) {
	var sl, sr string
	sl, sr = SplitString("1234")
	if sl != "12" {
		t.Error("Expected 12, got ", sl)
	} else if sr != "34" {
		t.Error("Expected 34, got ", sr)
	}
	sl, sr = SplitString("12345")
	if sl != "12" {
		t.Error("Expected 12, got ", sl)
	} else if sr != "345" {
		t.Error("Expected 345, got ", sr)
	}
}

func TestSameSize(t *testing.T) {
	var sl, sr string
	var n int
	n, sl, sr = SameSize("1234", `12`)
	if sl != "1234" {
		t.Error("Expected 1234, got ", sl)
	} else if sr != "0012" {
		t.Error("Expected 0012, got ", sr)
	} else if n != 4 {
		t.Error("Expected 4, got ", n)
	}

	n, sl, sr = SameSize(`55`, "123")
	if sl != "055" {
		t.Error("Expected 055, got ", sl)
	} else if sr != "123" {
		t.Error("Expected 123, got ", sr)
	} else if n != 3 {
		t.Error("Expected 3, got ", n)
	}

	n, sl, sr = SameSize(`55321`, "12345")
	if sl != "55321" {
		t.Error("Expected 55321, got ", sl)
	} else if sr != "12345" {
		t.Error("Expected 12345, got ", sr)
	} else if n != 5 {
		t.Error("Expected 5, got ", n)
	}

	n, sl, sr = SameSize(`1`, "1234567890")
	if sl != "0000000001" {
		t.Error("Expected 0000000001, got ", sl)
	} else if sr != "1234567890" {
		t.Error("Expected 1234567890, got ", sr)
	} else if n != 10 {
		t.Error("Expected 10, got ", n)
	}
}

func TestMul(t *testing.T) {
	f, err := os.Open("mul_cases.csv")
	if err != nil {
		t.Error("Can't open test file")
	}
	defer f.Close()

	r := csv.NewReader(f)
	var v string
	var n uint64
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("Can't read test file (may be not a csv?)")
		}

		Icounter = 0
		v, n = Mul(record[0], record[1], record[3])
		if v != record[2] {
			t.Error("Expected ", record[2], " got ", v)
		}
		if record[3] != `0` && record[4] != fmt.Sprintf("%d", n) {
			t.Error("Expected ", record[3], " encountered ", record[4], "times, got ", n)
		}

	}
}

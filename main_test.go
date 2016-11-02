package main

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	var v string
	v = Add("1", "2")
	if v != "3" {
		t.Error("Expected 3, got ", v)
	}
	v = Add("111111", "666666")
	if v != "777777" {
		t.Error("Expected 777777, got ", v)
	}
	v = Add("567", "789")
	if v != "1356" {
		t.Error("Expected 1356, got ", v)
	}
	v = Add("123", "123")
	if v != "246" {
		t.Error("Expected 246, got ", v)
	}
}

func TestSub(t *testing.T) {
	var v string
	v = Sub("3", "2")
	if v != "1" {
		t.Error("Expected 1, got ", v)
	}
	v = Sub("777777", "666666")
	if v != "111111" {
		t.Error("Expected 111111, got ", v)
	}
	v = Sub("567", "789")
	if v != "-222" {
		t.Error("Expected -222, got ", v)
	}
	v = Sub("767", "589")
	if v != "178" {
		t.Error("Expected 178, got ", v)
	}
	v = Sub("876", "642")
	if v != "234" {
		t.Error("Expected 234, got ", v)
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
}

func TestMul(t *testing.T) {
	f, err := os.Open("test_cases.csv")
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

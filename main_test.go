package main

import "testing"

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

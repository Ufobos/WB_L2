package main

import (
	"regexp"
	"testing"
)

func TestCountEntries(t *testing.T) {
	strs := []string{"hello", "world", "hello world", "test"}
	re := regexp.MustCompile(`hello`)
	want := 2
	got := countEntries(strs, re)
	if got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}

func TestFindMatch(t *testing.T) {
	strs := []string{"hello", "world", "hello world", "test"}
	re := regexp.MustCompile(`world`)
	want := []int{1, 2} // "world" and "hello world"
	got := findMatch(strs, re)
	if len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

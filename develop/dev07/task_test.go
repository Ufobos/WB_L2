package main

import (
	"testing"
	"time"
)

// Функция для создания канала, который закрывается через заданное время.
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

// Тест на проверку, что канал закрывается при закрытии первого из входных каналов.
func TestOrClosesWhenOneCloses(t *testing.T) {
	short := sig(100 * time.Millisecond)
	long := sig(1 * time.Minute)

	start := time.Now()
	<-or(short, long)
	duration := time.Since(start)

	if duration >= 1*time.Second {
		t.Errorf("or did not close in expected time; took %s", duration)
	}
}

// Тест на проверку, что функция or корректно обрабатывает уже закрытые каналы.
func TestOrWithAlreadyClosedChannels(t *testing.T) {
	alreadyClosed := make(chan interface{})
	close(alreadyClosed)
	other := sig(1 * time.Minute)

	start := time.Now()
	<-or(alreadyClosed, other)
	duration := time.Since(start)

	if duration >= 1*time.Second {
		t.Errorf("or did not close immediately when one channel was already closed; took %s", duration)
	}
}

// Тест на проверку, что or корректно обрабатывает ситуацию, когда все каналы закрываются одновременно.
func TestOrWithSimultaneousClose(t *testing.T) {
	a := sig(100 * time.Millisecond)
	b := sig(100 * time.Millisecond)

	start := time.Now()
	<-or(a, b)
	duration := time.Since(start)

	if duration > 200*time.Millisecond {
		t.Errorf("or did not close as expected when both channels closed simultaneously; took %s", duration)
	}
}

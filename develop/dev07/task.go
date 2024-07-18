package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	done := make(chan struct{})
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			select {
			case <-ch:
				done <- struct{}{} // Отправляет сигнал в канал done, если получен сигнал из одного из каналов ch
			case <-done:
				return // Выходит из горутины, если был получен сигнал завершения из канала done
			}
		}(ch)
	}
	go func() {
		<-done        // Ожидает сигнала завершения
		close(done)   // Закрывает канал done после получения сигнала, чтобы прекратить работу всех горутин
		close(single) // Закрывает канал single, уведомляя о завершении
	}()
	return single
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(20200*time.Millisecond),
		sig(5*time.Minute),
		sig(15*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}

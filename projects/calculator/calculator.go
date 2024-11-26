package main

import (
 "fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
 outputChan := make(chan int)

 go func() {
  defer close(outputChan) // Закрываем выходной канал при завершении работы

  select {
  case num := <-firstChan:
   outputChan <- num * num // Отправляем квадрат числа
  case num := <-secondChan:
   outputChan <- num * 3 // Отправляем результат умножения на 3
  case <-stopChan:
   // Если получен сигнал остановки, просто завершаем работу
   return
  }
 }()

 return outputChan
}

func main() {
 firstChan := make(chan int)
 secondChan := make(chan int)
 stopChan := make(chan struct{})

 outputChan := calculator(firstChan, secondChan, stopChan)

 // Пример использования
 go func() {
  firstChan <- 4 // Отправляем число в firstChan
 }()

 go func() {
  secondChan <- 5 // Отправляем число в secondChan
 }()

 go func() {
  stopChan <- struct{}{} // Отправляем сигнал остановки
 }()

 for result := range outputChan {
  fmt.Println(result) // Печатаем результаты
 }
}
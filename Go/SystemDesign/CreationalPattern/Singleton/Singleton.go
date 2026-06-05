package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	logCount int
}

var (
	instance *Logger
	once sync.Once
)

func GetLogger() *Logger{
	once.Do(func() {
		fmt.Println("Creating Logger instance")
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) Log(message string) {
	l.logCount++
	fmt.Printf("[LOG %d] %s\n", l.logCount, message)
}


// func main()  {

// 	logger1 := GetLogger()
// 	logger2 := GetLogger()

// 	logger1.Log("Application started")
// 	logger2.Log("Database connected")

// 	fmt.Println("Same instance?", logger1 == logger2)

// }

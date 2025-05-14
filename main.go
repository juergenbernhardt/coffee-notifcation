package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	reset := make(chan bool)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadByte()
			if input == 'x' {
				reset <- true
			}
		}
	}()

	interval := 1 * time.Second // Beispiel: alle 10 Sekunden
	timer := time.NewTimer(interval)

	for {
		select {
		case <-timer.C:
			ClearConsole()
			fmt.Println("Kaffepause bis ")
			timer.Reset(interval)
		case <-reset:
			ClearConsole()
			fmt.Println("Heiß Schwarz und lecker!")
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(interval)
		}
	}

}

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

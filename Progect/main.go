package main

import (
	"fmt"
	"strings"
	"time"
)

// Stage описывает одну стадию конвейера: читает из входного канала и пишет в выходной.
type Stage func(<-chan string) <-chan string

// echoStage имитирует echo / cat: из набора строк создаёт исходный канал.
func echoStage(lines ...string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()

	return out
}

// grepStage имитирует grep pattern: пропускает только строки, содержащие подстроку pattern.
func grepStage(pattern string) Stage {
	return func(in <-chan string) <-chan string {
		out := make(chan string)

		go func() {
			defer close(out)
			for line := range in {
				if strings.Contains(line, pattern) {
					out <- line
				}
			}
		}()

		return out
	}
}

// toUpperStage имитирует tr a-z A-Z: переводит строку в верхний регистр.
func toUpperStage() Stage {
	return func(in <-chan string) <-chan string {
		out := make(chan string)

		go func() {
			defer close(out)
			for line := range in {
				out <- strings.ToUpper(line)
			}
		}()

		return out
	}
}

// prefixStage добавляет префикс к каждой строке.
func prefixStage(prefix string) Stage {
	return func(in <-chan string) <-chan string {
		out := make(chan string)

		go func() {
			defer close(out)
			for line := range in {
				out <- prefix + line
			}
		}()

		return out
	}
}

// pipeline последовательно соединяет несколько стадий, как команды через | в shell.
func pipeline(in <-chan string, stages ...Stage) <-chan string {
	out := in
	for _, stage := range stages {
		out = stage(out)
	}
	return out
}

func main() {
	// Имитация Unix-пайплайна:
	// echo "error: disk full" "ok: done" "error: network" |
	//   grep "error" |
	//   tr a-z A-Z |
	//   prefix "[LOG] "

	source := echoStage(
		"error: disk full",
		"info: starting service",
		"ok: done",
		"error: network unreachable",
	)

	result := pipeline(
		source,
		grepStage("error"),
		toUpperStage(),
		prefixStage("[LOG] "),
	)

	for line := range result {
		fmt.Println(time.Now().Format("15:04:05"), line)
	}
}

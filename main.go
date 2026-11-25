package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func contextFor(mode string, d time.Duration) (context.Context, context.CancelFunc, error) {
	parent := context.Background()
	switch mode {
	case "cancel":
		return context.WithCancel(parent)
	case "deadline":
		return context.WithDeadline(parent, time.Now().Add(d))
	case "timeout":
		return context.WithTimeout(parent, d)
	default:
		return nil, nil, fmt.Errorf("unknown scenario %q", mode)
	}
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(250 * time.Millisecond)
	defer timer.Stop()
	fmt.Printf("worker %d: started\n", id)
	select {
	case <-ctx.Done():
		fmt.Printf("worker %d: stopped: %v\n", id, ctx.Err())
	case <-timer.C:
		fmt.Printf("worker %d: completed before cancellation\n", id)
	}
}

func run(mode string, workers int) error {
	ctx, cancel, err := contextFor(mode, 65*time.Millisecond)
	if err != nil { return err }
	defer cancel()

	var wg sync.WaitGroup
	for id := 1; id <= workers; id++ {
		wg.Add(1)
		go worker(ctx, id, &wg)
	}
	if mode == "cancel" {
		time.Sleep(35 * time.Millisecond)
		fmt.Println("main: calling cancel()")
		cancel()
	}
	wg.Wait()
	fmt.Printf("context result: %v\n\n", ctx.Err())
	return nil
}

func main() {
	scenario := flag.String("scenario", "all", "cancel, deadline, timeout, or all")
	workers := flag.Int("workers", 3, "number of worker goroutines")
	flag.Parse()
	if *workers < 1 || *workers > 20 {
		fmt.Println("workers must be between 1 and 20")
		return
	}
	modes := []string{*scenario}
	if *scenario == "all" { modes = []string{"cancel", "deadline", "timeout"} }
	for _, mode := range modes {
		fmt.Printf("=== %s ===\n", mode)
		if err := run(mode, *workers); err != nil { fmt.Println(err); return }
	}
}

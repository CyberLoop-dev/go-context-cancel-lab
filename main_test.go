package main

import (
	"context"
	"testing"
	"time"
)

func TestContextCancellationPropagates(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	child, childCancel := context.WithCancel(parent)
	defer childCancel()
	cancel()
	select {
	case <-child.Done():
	case <-time.After(time.Second):
		t.Fatal("child context did not receive parent cancellation")
	}
}

func TestContextForRejectsUnknownMode(t *testing.T) {
	if _, _, err := contextFor("unknown", time.Millisecond); err == nil {
		t.Fatal("expected an error for an unknown mode")
	}
}

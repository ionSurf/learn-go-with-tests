package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Correct output", func(t *testing.T) {
		buffer_ptr := &bytes.Buffer{}
		spySleeper_ptr := &SpyCountdownOperations{}

		Countdown(buffer_ptr, spySleeper_ptr)

		got := buffer_ptr.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Verify correct order of operations", func(t *testing.T) {
		spySleeper := &SpyCountdownOperations{}

		Countdown(spySleeper, spySleeper)

		got := spySleeper.Calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted calls %v got %v", want, got)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("Sleeps for x time", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

		// Action
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("Should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}

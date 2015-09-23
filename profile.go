package main

import (
	"os"
	"os/signal"
	"runtime/pprof"
)

func startCPUProfile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return nil
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		return err
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			switch sig.String() {
			case "interrupt":
				pprof.StopCPUProfile()
				f.Close()
				os.Exit(0)
			}
		}
	}()
	return nil
}

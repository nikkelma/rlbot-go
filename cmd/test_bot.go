package main

import (
	"github.com/nikkelma/rlbot-go/flat"
	"github.com/nikkelma/rlbot-go/native"

	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	bridge, err := native.NewBridge()
	defer bridge.Close()

	if err != nil {
		fmt.Println("NewBridge error")
		fmt.Println(err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	init, err := bridge.IsInitialized()

	for {
		if init {
			break
		}
		<-time.After(1 * time.Second)
		init, err = bridge.IsInitialized()
	}

	fieldInfo, err := bridge.UpdateFieldInfo()
	if err != nil {
		fmt.Println("UpdateFieldInfo error")
		fmt.Println(err)
		return
	}

	boostPadLength := fieldInfo.BoostPadsLength()
	if boostPadLength == 0 {
		fmt.Println("BoostPads length is 0")
		return
	}
	boostPad := &flat.BoostPad{}
	success := fieldInfo.BoostPads(boostPad, 0)
	if !success {
		fmt.Println("BoostPads error")
		fmt.Println(err)
		return
	}
	fmt.Println("BoostPad")
	fmt.Println(boostPad.IsFullBoost())
}

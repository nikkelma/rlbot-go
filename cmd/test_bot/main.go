package main

import (
	"github.com/nikkelma/rlbot-go/pkg/flat"
	"github.com/nikkelma/rlbot-go/pkg/native"

	"fmt"
	// "os"
	// "os/signal"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	bridge, err := native.NewBridge()
	if err != nil {
		fmt.Println("NewBridge error")
		fmt.Println(err)
		return
	}
	defer bridge.Close()

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// signal.Notify(c, os.Kill)

	init, err := bridge.IsInitialized()

	for {
		if err != nil {
			fmt.Println("IsInitialized error")
			fmt.Println(err)
			return
		}
		if init {
			break
		} else {
			fmt.Println("Not initialized yet...")
		}
		<-time.After(2 * time.Second)
		init, err = bridge.IsInitialized()
	}

	fmt.Println("Initialized!")

	matchSettings, err := bridge.GetMatchSettings()
	if err != nil {
		fmt.Println("GetMatchSettings error")
		fmt.Println(err)
		return
	}

	gameModeStr := flat.EnumNamesGameMode[int(matchSettings.GameMode())]
	gameMapStr := flat.EnumNamesGameMap[int(matchSettings.GameMap())]

	fmt.Printf("GameMode: %v\n", gameModeStr)
	fmt.Printf("GameMap: %v\n", gameMapStr)

	builder := flatbuffers.NewBuilder(0)
	flat.QuickChatStart(builder)
	flat.QuickChatAddQuickChatSelection(builder, 5)
	flat.QuickChatAddPlayerIndex(builder, 1)
	offset := flat.QuickChatEnd(builder)

	builder.Finish(offset)

	err = bridge.SendQuickChat(flat.GetRootAsQuickChat(builder.FinishedBytes(), offset))
	if err != nil {
		fmt.Printf("SendQuickChat error: %v", err)
		return
	}
}

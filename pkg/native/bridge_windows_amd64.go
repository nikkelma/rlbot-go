// +build windows amd64

package native

import (
	"github.com/nikkelma/rlbot-go/pkg/flat"
	rlbotstatus "github.com/nikkelma/rlbot-go/pkg/native/status"

	"fmt"
	"reflect"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type byteBuffer struct {
	ptr  uintptr
	size int32
}

const dllName string = "RLBot_Core_Interface.dll"

// TODO - combine architecture-specific implementations?
func newBridge() (Bridge, error) {
	fmt.Println(dllName)

	// TODO - confirm avoidance of DLL pre-loading attacks
	rlBotInterfaceDll, err := windows.LoadDLL(dllName)
	if err != nil {
		return nil, err
	}

	if rlBotInterfaceDll == nil {
		err = fmt.Errorf("nil rlBotInterfaceDll error")
		return nil, err
	}
	fmt.Println(rlBotInterfaceDll)

	bridge := &bridgeWindowsAmd64{
		rlBotInterfaceDll:                  rlBotInterfaceDll,
		freeProc:                           newWindowsProc(rlBotInterfaceDll, "Free"),
		isInitializedProc:                  newWindowsProc(rlBotInterfaceDll, "IsInitialized"),
		getBallPredictionProc:              newWindowsProc(rlBotInterfaceDll, "GetBallPrediction"),
		setGameStateProc:                   newWindowsProc(rlBotInterfaceDll, "SetGameState"),
		startMatchFlatbufferProc:           newWindowsProc(rlBotInterfaceDll, "StartMatchFlatbuffer"),
		updateFieldInfoFlatbufferProc:      newWindowsProc(rlBotInterfaceDll, "UpdateFieldInfoFlatbuffer"),
		updateLiveDataPacketFlatbufferProc: newWindowsProc(rlBotInterfaceDll, "UpdateLiveDataPacketFlatbuffer"),
		updateRigidBodyTickFlatbufferProc:  newWindowsProc(rlBotInterfaceDll, "UpdateRigidBodyTickFlatbuffer"),
		getMatchSettingsProc:               newWindowsProc(rlBotInterfaceDll, "GetMatchSettings"),
		sendQuickChatProc:                  newWindowsProc(rlBotInterfaceDll, "SendQuickChat"),
		receiveChatProc:                    newWindowsProc(rlBotInterfaceDll, "ReceiveChat"),
		updatePlayerInputFlatbufferProc:    newWindowsProc(rlBotInterfaceDll, "UpdatePlayerInput"),
		renderGroupProc:                    newWindowsProc(rlBotInterfaceDll, "RenderGroup"),
	}
	return bridge, nil
}

// TODO - combine architecture-specific implementations?
type bridgeWindowsAmd64 struct {
	rlBotInterfaceDll                  *windows.DLL
	freeProc                           windowsProc
	isInitializedProc                  windowsProc
	getBallPredictionProc              windowsProc
	setGameStateProc                   windowsProc
	startMatchFlatbufferProc           windowsProc
	updateFieldInfoFlatbufferProc      windowsProc
	updateLiveDataPacketFlatbufferProc windowsProc
	updateRigidBodyTickFlatbufferProc  windowsProc
	getMatchSettingsProc               windowsProc
	sendQuickChatProc                  windowsProc
	receiveChatProc                    windowsProc
	updatePlayerInputFlatbufferProc    windowsProc
	renderGroupProc                    windowsProc
}

func (b *bridgeWindowsAmd64) Close() error {
	fmt.Println("bridgeWindowsAmd64 Close")

	if b.rlBotInterfaceDll == nil {
		return fmt.Errorf("nil rlBotInterfaceDll")
	}
	return b.rlBotInterfaceDll.Release()
}

// Interface.hpp
func (b *bridgeWindowsAmd64) IsInitialized() (bool, error) {
	b.isInitializedProc.Lock()
	defer b.isInitializedProc.Unlock()

	res, _, errno := b.isInitializedProc.Call()
	if errno != syscall.Errno(0) {
		return false, errno
	}
	return res != 0, nil
}

// GameFunctions/BallPrediction.hpp
func (b *bridgeWindowsAmd64) GetBallPrediction() (*flat.BallPrediction, error) {
	b.getBallPredictionProc.Lock()
	defer b.getBallPredictionProc.Unlock()

	ballPredictionByteBuffer := byteBuffer{}
	_, _, errno := b.getBallPredictionProc.Call(uintptr(unsafe.Pointer(&ballPredictionByteBuffer)))

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("GetBallPrediction error: %v", errno)
	}

	ptr := ballPredictionByteBuffer.ptr
	size := ballPredictionByteBuffer.size

	ballPredictionBytes := make([]byte, size)
	for i := 0; i < int(size); i++ {
		ballPredictionBytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	_, _, errno = b.freeProc.Call(ptr)
	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("Free error: %v", errno)
	}

	ballPrediction := flat.GetRootAsBallPrediction(ballPredictionBytes, 0)

	return ballPrediction, nil
}

// GameFunctions/GameFunctions.hpp
func (b *bridgeWindowsAmd64) SetGameState(gameState *flat.DesiredGameState) error {
	gameStateBytes := gameState.Table().Bytes
	gameStateSize := len(gameStateBytes)
	gameStateBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&gameStateBytes))

	b.setGameStateProc.Lock()
	defer b.setGameStateProc.Unlock()

	status, _, errno := b.setGameStateProc.Call(gameStateBytesHeader.Data, uintptr(gameStateSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("SetGameState error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("SetGameState bad status: %v", rlbotStatus)
	}

	return nil
}

func (b *bridgeWindowsAmd64) StartMatch(matchSettings *flat.MatchSettings) error {
	matchSettingsBytes := matchSettings.Table().Bytes
	matchSettingsSize := len(matchSettingsBytes)
	matchSettingsBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&matchSettingsBytes))

	b.startMatchFlatbufferProc.Lock()
	defer b.startMatchFlatbufferProc.Unlock()

	status, _, errno := b.startMatchFlatbufferProc.Call(matchSettingsBytesHeader.Data, uintptr(matchSettingsSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("StartMatch error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("StartMatch bad status: %v", rlbotStatus)
	}

	return nil
}

// GameFunctions/GamePacket.hpp
func (b *bridgeWindowsAmd64) GetFieldInfo() (*flat.FieldInfo, error) {
	b.updateFieldInfoFlatbufferProc.Lock()
	defer b.updateFieldInfoFlatbufferProc.Unlock()

	fieldInfoByteBuffer := byteBuffer{}
	_, _, errno := b.updateFieldInfoFlatbufferProc.Call(uintptr(unsafe.Pointer(&fieldInfoByteBuffer)))

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("GetFieldInfo error: %v", errno)
	}

	ptr := fieldInfoByteBuffer.ptr
	size := fieldInfoByteBuffer.size

	fieldInfoBytes := make([]byte, size)
	for i := 0; i < int(size); i++ {
		fieldInfoBytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	_, _, errno = b.freeProc.Call(ptr)
	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("Free error: %v", errno)
	}

	fieldInfo := flat.GetRootAsFieldInfo(fieldInfoBytes, 0)

	return fieldInfo, nil
}

func (b *bridgeWindowsAmd64) GetLiveGameTickPacket() (*flat.GameTickPacket, error) {
	b.updateLiveDataPacketFlatbufferProc.Lock()
	defer b.updateLiveDataPacketFlatbufferProc.Unlock()

	gameTickPacketByteBuffer := byteBuffer{}
	_, _, errno := b.updateLiveDataPacketFlatbufferProc.Call(uintptr(unsafe.Pointer(&gameTickPacketByteBuffer)))

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("GetLiveGameTickPacket error: %v", errno)
	}

	ptr := gameTickPacketByteBuffer.ptr
	size := gameTickPacketByteBuffer.size

	gameTickPacketBytes := make([]byte, size)
	for i := 0; i < int(size); i++ {
		gameTickPacketBytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	_, _, errno = b.freeProc.Call(ptr)
	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("Free error: %v", errno)
	}

	gameTickPacket := flat.GetRootAsGameTickPacket(gameTickPacketBytes, 0)

	return gameTickPacket, nil
}

func (b *bridgeWindowsAmd64) GetMatchSettings() (*flat.MatchSettings, error) {
	b.getMatchSettingsProc.Lock()
	defer b.getMatchSettingsProc.Unlock()

	matchSettingsByteBuffer := byteBuffer{}
	_, _, errno := b.getMatchSettingsProc.Call(uintptr(unsafe.Pointer(&matchSettingsByteBuffer)))

	if errno != syscall.Errno(0) {
		fmt.Printf("GetMatchSettings error: %v", errno)
		return nil, fmt.Errorf("GetMatchSettings error: %v", errno)
	}

	ptr := matchSettingsByteBuffer.ptr
	size := matchSettingsByteBuffer.size

	matchSettingsBytes := make([]byte, size)
	for i := 0; i < int(size); i++ {
		matchSettingsBytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	_, _, errno = b.freeProc.Call(ptr)
	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("Free error: %v", errno)
	}

	matchSettings := flat.GetRootAsMatchSettings(matchSettingsBytes, 0)

	return matchSettings, nil
}

// GameFunctions/PlayerInfo.hpp
func (b *bridgeWindowsAmd64) SendQuickChat(quickChat *flat.QuickChat) error {
	quickChatBytes := quickChat.Table().Bytes
	quickChatSize := len(quickChatBytes)
	quickChatBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&quickChatBytes))

	b.sendQuickChatProc.Lock()
	defer b.sendQuickChatProc.Unlock()

	status, _, errno := b.sendQuickChatProc.Call(quickChatBytesHeader.Data, uintptr(quickChatSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("SendQuickChat error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("SendQuickChat bad status: %v", rlbotStatus)
	}

	return nil
}

func (b *bridgeWindowsAmd64) ReceiveChat(botIndex, teamIndex, lastMessageIndex int) (*flat.QuickChatMessages, error) {
	b.receiveChatProc.Lock()
	defer b.receiveChatProc.Unlock()

	ptr, size, errno := b.receiveChatProc.Call(uintptr(botIndex), uintptr(teamIndex), uintptr(lastMessageIndex))

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("ReceiveChat error: %v", errno)
	}

	quickChatMessagesBytes := make([]byte, size)
	for i := 0; i < int(size); i++ {
		quickChatMessagesBytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	_, _, errno = b.freeProc.Call(ptr)
	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("Free error: %v", errno)
	}

	quickChatMessages := flat.GetRootAsQuickChatMessages(quickChatMessagesBytes, 0)

	return quickChatMessages, nil
}

func (b *bridgeWindowsAmd64) UpdatePlayerInput(playerInput *flat.PlayerInput) error {
	playerInputBytes := playerInput.Table().Bytes
	playerInputSize := len(playerInputBytes)
	playerInputBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&playerInputBytes))

	b.updatePlayerInputFlatbufferProc.Lock()
	defer b.updatePlayerInputFlatbufferProc.Unlock()

	status, _, errno := b.updatePlayerInputFlatbufferProc.Call(playerInputBytesHeader.Data, uintptr(playerInputSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("UpdatePlayerInput error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("UpdatePlayerInput bad status: %v", rlbotStatus)
	}

	return nil
}

// RenderFunctions/RenderFunctions.hpp
func (b *bridgeWindowsAmd64) RenderGroup(renderGroup *flat.RenderGroup) error {
	renderGroupBytes := renderGroup.Table().Bytes
	renderGroupSize := len(renderGroupBytes)
	renderGroupBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&renderGroupBytes))

	b.renderGroupProc.Lock()
	defer b.renderGroupProc.Unlock()

	status, _, errno := b.renderGroupProc.Call(renderGroupBytesHeader.Data, uintptr(renderGroupSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("RenderGroup error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("RenderGroup bad status: %v", rlbotStatus)
	}

	return nil
}

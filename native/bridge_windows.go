// +build windows

package native

import (
	"github.com/nikkelma/rlbot-go/flat"
	rlbotstatus "github.com/nikkelma/rlbot-go/native/status"

	"fmt"
	"reflect"
	// "runtime"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type windowsProc struct {
	sync.Mutex
	*windows.Proc
}

func createWindowsProc(dll *windows.DLL, procName string) *windowsProc {
	return &windowsProc{
		Proc: dll.MustFindProc(procName),
	}
}

func newBridge() (Bridge, error) {
	var dllName string = getDllName()
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

	bridge := &bridgeWindows{
		rlBotInterfaceDll: rlBotInterfaceDll,
		freeProc:          createWindowsProc(rlBotInterfaceDll, "Free"),
	}
	return bridge, nil
}

type bridgeWindows struct {
	rlBotInterfaceDll                  *windows.DLL
	freeProc                           *windowsProc
	isInitializedProc                  *windowsProc
	getBallPredictionProc              *windowsProc
	setGameStateProc                   *windowsProc
	startMatchFlatbufferProc           *windowsProc
	updateFieldInfoFlatbufferProc      *windowsProc
	updateLiveDataPacketFlatbufferProc *windowsProc
	updateRigidBodyTickFlatbufferProc  *windowsProc
	getMatchSettingsProc               *windowsProc
	sendQuickChatProc                  *windowsProc
	sendChatProc                       *windowsProc
	receiveChatProc                    *windowsProc
	updatePlayerInputFlatbufferProc    *windowsProc
	renderGroupProc                    *windowsProc
}

func (b *bridgeWindows) Close() error {
	fmt.Println("bridgeWindows Close")

	if b.rlBotInterfaceDll == nil {
		return fmt.Errorf("nil rlBotInterfaceDll")
	}
	return b.rlBotInterfaceDll.Release()
}

// Interface.hpp
func (b *bridgeWindows) IsInitialized() (bool, error) {
	if b.isInitializedProc == nil {
		b.isInitializedProc = createWindowsProc(b.rlBotInterfaceDll, "IsInitialized")
	}

	b.isInitializedProc.Lock()
	defer b.isInitializedProc.Unlock()

	res, _, errno := b.isInitializedProc.Call()
	if errno != syscall.Errno(0) {
		return false, errno
	}
	return res != 0, nil
}

// GameFunctions/BallPrediction.hpp
func (b *bridgeWindows) GetBallPrediction() (*flat.BallPrediction, error) {
	if b.getBallPredictionProc == nil {
		b.getBallPredictionProc = createWindowsProc(b.rlBotInterfaceDll, "GetBallPrediction")
	}

	b.getBallPredictionProc.Lock()
	defer b.getBallPredictionProc.Unlock()

	ptr, size, errno := b.getBallPredictionProc.Call()

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("GetBallPrediction error: %v", errno)
	}

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
func (b *bridgeWindows) SetGameState(gameState *flat.DesiredGameState) error {
	if b.setGameStateProc == nil {
		b.setGameStateProc = createWindowsProc(b.rlBotInterfaceDll, "SetGameState")
	}

	gameStateBytes := gameState.Table().Bytes
	gameStateSize := len(gameStateBytes)
	gameStateBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&gameStateBytes))
	
	b.setGameStateProc.Lock()
	defer b.setGameStateProc.Unlock()

	status, _, errno :=	b.setGameStateProc.Call(gameStateBytesHeader.Data, uintptr(gameStateSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("SetGameState error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("SetGameState bad status: %v", rlbotStatus)
	}

	return nil
}

func (b *bridgeWindows) StartMatch(matchSettings *flat.MatchSettings) error {
	if b.startMatchFlatbufferProc == nil {
		b.startMatchFlatbufferProc = createWindowsProc(b.rlBotInterfaceDll, "StartMatchFlatbuffer")
	}

	matchSettingsBytes := matchSettings.Table().Bytes
	matchSettingsSize := len(matchSettingsBytes)
	matchSettingsBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&matchSettingsBytes))
	
	b.startMatchFlatbufferProc.Lock()
	defer b.startMatchFlatbufferProc.Unlock()

	status, _, errno :=	b.startMatchFlatbufferProc.Call(matchSettingsBytesHeader.Data, uintptr(matchSettingsSize))

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
func (b *bridgeWindows) UpdateFieldInfo() (*flat.FieldInfo, error) {
	if b.updateFieldInfoFlatbufferProc == nil {
		b.updateFieldInfoFlatbufferProc = createWindowsProc(b.rlBotInterfaceDll, "UpdateFieldInfoFlatbuffer")
	}

	b.updateFieldInfoFlatbufferProc.Lock()
	defer b.updateFieldInfoFlatbufferProc.Unlock()

	ptr, size, errno := b.updateFieldInfoFlatbufferProc.Call()

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("UpdateFieldInfo error: %v", errno)
	}

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

func (b *bridgeWindows) UpdateLiveDataPacket() (*flat.GameTickPacket, error) {
	if b.updateLiveDataPacketFlatbufferProc == nil {
		b.updateLiveDataPacketFlatbufferProc = createWindowsProc(b.rlBotInterfaceDll, "UpdateLiveDataPacketFlatbuffer")
	}

	b.updateLiveDataPacketFlatbufferProc.Lock()
	defer b.updateLiveDataPacketFlatbufferProc.Unlock()

	ptr, size, errno := b.updateLiveDataPacketFlatbufferProc.Call()

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("UpdateLiveDataPacket error: %v", errno)
	}

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

func (b *bridgeWindows) UpdateRigidBodyTick() (*flat.RigidBodyTick, error) {
	if b.updateRigidBodyTickFlatbufferProc == nil {
		b.updateRigidBodyTickFlatbufferProc = createWindowsProc(b.rlBotInterfaceDll, "UpdateRigidBodyTickFlatbuffer")
	}

	b.updateRigidBodyTickFlatbufferProc.Lock()
	defer b.updateRigidBodyTickFlatbufferProc.Unlock()

	ptr, size, errno := b.updateRigidBodyTickFlatbufferProc.Call()

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("UpdateRigidBodyTick error: %v", errno)
	}

	rigidBodyTickBytes := make([]byte, size)
	for i := 0; i < int(size); i++ {
		rigidBodyTickBytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	_, _, errno = b.freeProc.Call(ptr)
	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("Free error: %v", errno)
	}

	rigidBodyTick := flat.GetRootAsRigidBodyTick(rigidBodyTickBytes, 0)

	return rigidBodyTick, nil
}

func (b *bridgeWindows) GetMatchSettings() (*flat.MatchSettings, error) {
	if b.getMatchSettingsProc == nil {
		b.getMatchSettingsProc = createWindowsProc(b.rlBotInterfaceDll, "GetMatchSettings")
	}

	b.getMatchSettingsProc.Lock()
	defer b.getMatchSettingsProc.Unlock()

	ptr, size, errno := b.getMatchSettingsProc.Call()

	if errno != syscall.Errno(0) {
		return nil, fmt.Errorf("GetMatchSettings error: %v", errno)
	}

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
func (b *bridgeWindows) SendQuickChat(quickChat *flat.QuickChat) error {
	if b.sendQuickChatProc == nil {
		b.sendQuickChatProc = createWindowsProc(b.rlBotInterfaceDll, "SendQuickChat")
	}

	quickChatBytes := quickChat.Table().Bytes
	quickChatSize := len(quickChatBytes)
	quickChatBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&quickChatBytes))
	
	b.sendQuickChatProc.Lock()
	defer b.sendQuickChatProc.Unlock()

	status, _, errno :=	b.sendQuickChatProc.Call(quickChatBytesHeader.Data, uintptr(quickChatSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("SendQuickChat error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("SendQuickChat bad status: %v", rlbotStatus)
	}

	return nil
}

func (b *bridgeWindows) ReceiveChat(botIndex, teamIndex, lastMessageIndex int) (*flat.QuickChatMessages, error) {
	if b.receiveChatProc == nil {
		b.receiveChatProc = createWindowsProc(b.rlBotInterfaceDll, "ReceiveChat")
	}

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

func (b *bridgeWindows) UpdatePlayerInput(playerInput *flat.PlayerInput) error {
	if b.updatePlayerInputFlatbufferProc == nil {
		b.updatePlayerInputFlatbufferProc = createWindowsProc(b.rlBotInterfaceDll, "UpdatePlayerInputFlatbuffer")
	}

	playerInputBytes := playerInput.Table().Bytes
	playerInputSize := len(playerInputBytes)
	playerInputBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&playerInputBytes))
	
	b.updatePlayerInputFlatbufferProc.Lock()
	defer b.updatePlayerInputFlatbufferProc.Unlock()

	status, _, errno :=	b.updatePlayerInputFlatbufferProc.Call(playerInputBytesHeader.Data, uintptr(playerInputSize))

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
func (b *bridgeWindows) RenderGroup(renderGroup *flat.RenderGroup) error {
	if b.renderGroupProc == nil {
		b.renderGroupProc = createWindowsProc(b.rlBotInterfaceDll, "RenderGroup")
	}

	renderGroupBytes := renderGroup.Table().Bytes
	renderGroupSize := len(renderGroupBytes)
	renderGroupBytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&renderGroupBytes))
	
	b.renderGroupProc.Lock()
	defer b.renderGroupProc.Unlock()

	status, _, errno :=	b.renderGroupProc.Call(renderGroupBytesHeader.Data, uintptr(renderGroupSize))

	if errno != syscall.Errno(0) {
		return fmt.Errorf("RenderGroup error: %v", errno)
	}

	rlbotStatus := rlbotstatus.RLBotCoreStatus(status)

	if rlbotStatus != rlbotstatus.Success {
		return fmt.Errorf("RenderGroup bad status: %v", rlbotStatus)
	}

	return nil
}

// ensure *bridgeWindows satisfies Bridge interface
var _ Bridge = &bridgeWindows{}

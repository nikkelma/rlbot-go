// +build windows

package native

import (
	"fmt"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/nikkelma/rlbot-go/flat"
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

// TODO - determine if locks are needed, refactor to struct to allow for locks
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
	return nil, fmt.Errorf("Not implemented")
}

// GameFunctions/GameFunctions.hpp
func (b *bridgeWindows) SetGameState(*flat.DesiredGameState) error {
	return fmt.Errorf("Not implemented")
}

func (b *bridgeWindows) StartMatch(*flat.MatchSettings) error {
	return fmt.Errorf("Not implemented")
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
	return nil, fmt.Errorf("Not implemented")
}

func (b *bridgeWindows) UpdateRigidBodyTick() (*flat.RigidBodyTick, error) {
	return nil, fmt.Errorf("Not implemented")
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
func (b *bridgeWindows) SendQuickChat(*flat.QuickChat) error {
	return fmt.Errorf("Not implemented")
}

func (b *bridgeWindows) ReceiveChat() (*flat.QuickChatMessages, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (b *bridgeWindows) UpdatePlayerInput(*flat.PlayerInput) error {
	return fmt.Errorf("Not implemented")
}

// RenderFunctions/RenderFunctions.hpp
func (b *bridgeWindows) RenderGroup(*flat.RenderGroup) error {
	return fmt.Errorf("Not implemented")
}

// ensure *bridgeWindows satisfies Bridge interface
var _ Bridge = &bridgeWindows{}

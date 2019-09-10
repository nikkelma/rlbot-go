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
	*sync.Mutex
	*windows.Proc
}
func createWindowsProc(dll *windows.DLL, procName string) *windowsProc {
	return &windowsProc{
		Proc: dll.MustFindProc(procName),
	}
}

func newBridge() (bridge Bridge, err error) {
	var dllName string = getDllName()
	fmt.Println(dllName)
	// TODO - confirm avoidance of DLL pre-loading attacks
	rlBotInterfaceDll, err := windows.LoadDLL(dllName)
	if err != nil {
		return nil, err
	}

	// TODO - only load Proc when needed
	_bridge := &bridgeWindows{
		rlBotInterfaceDll: rlBotInterfaceDll,
		freeProc:          createWindowsProc(rlBotInterfaceDll, "Free"),
	}
	return _bridge, nil
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
	return nil, nil
}

// GameFunctions/GameFunctions.hpp
func (b *bridgeWindows) SetGameState(*flat.DesiredGameState) error {
	return nil
}
func (b *bridgeWindows) StartMatch(*flat.MatchSettings) error {
	return nil
}

// GameFunctions/GamePacket.hpp
func (b *bridgeWindows) UpdateFieldInfo() (*flat.FieldInfo, error) {
	if b.updateFieldInfoFlatbufferProc == nil {
		b.updateFieldInfoFlatbufferProc = createWindowsProc(b.rlBotInterfaceDll, "UpdateLiveDataPacketFlatbuffer")
	}

	b.updateFieldInfoFlatbufferProc.Lock()
	defer b.updateFieldInfoFlatbufferProc.Unlock()

	res, _, errno := b.updateFieldInfoFlatbufferProc.Call()
	if errno != syscall.Errno(0) {
		return nil, errno
	}

	fmt.Println(res)

	// myvar := *(*C.ByteBuffer)(unsafe.Pointer(res))
	// fieldInfo := flat.GetRootAsFieldInfo(*myvar.ptr, myvar.size)

	// fmt.Println(fieldInfo)

	return nil, nil
}
func (b *bridgeWindows) UpdateLiveDataPacket() (*flat.GameTickPacket, error) {
	return nil, nil
}
func (b *bridgeWindows) UpdateRigidBodyTick() (*flat.RigidBodyTick, error) {
	return nil, nil
}
func (b *bridgeWindows) GetMatchSettings() (*flat.MatchSettings, error) {
	return nil, nil
}

// GameFunctions/PlayerInfo.hpp
func (b *bridgeWindows) SendQuickChat(*flat.QuickChat) error {
	return nil
}
func (b *bridgeWindows) ReceiveChat() (*flat.QuickChatMessages, error) {
	return nil, nil
}
func (b *bridgeWindows) UpdatePlayerInput(*flat.PlayerInput) error {
	return nil
}

// RenderFunctions/RenderFunctions.hpp
func (b *bridgeWindows) RenderGroup(*flat.RenderGroup) error {
	return nil
}

// ensure *bridgeWindows satisfies Bridge interface
var _ Bridge = &bridgeWindows{}

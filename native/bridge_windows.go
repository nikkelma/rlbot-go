// +build windows

package native

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"

	"github.com/nikkelma/rlbot-go/flat"
)

func newBridge() (bridge Bridge, err error) {
	var dllName string = getDllName()
	fmt.Print(dllName)
	// TODO - confirm avoidance of DLL pre-loading attacks
	rlBotInterfaceDll, err := windows.LoadDLL(dllName)
	if err != nil {
		return nil, err
	}

	// TODO - only load Proc when needed
	_bridge := &bridgeWindows{
		isInitializedProc:                  rlBotInterfaceDll.MustFindProc("IsInitialized"),
		getBallPredictionProc:              rlBotInterfaceDll.MustFindProc("GetBallPrediction"),
		freeProc:                           rlBotInterfaceDll.MustFindProc("Free"),
		setGameStateProc:                   rlBotInterfaceDll.MustFindProc("SetGameState"),
		startMatchFlatbufferProc:           rlBotInterfaceDll.MustFindProc("StartMatchFlatbuffer"),
		updateFieldInfoFlatbufferProc:      rlBotInterfaceDll.MustFindProc("UpdateFieldInfoFlatbuffer"),
		updateLiveDataPacketFlatbufferProc: rlBotInterfaceDll.MustFindProc("UpdateLiveDataPacketFlatbuffer"),
		updateRigidBodyTickFlatbufferProc:  rlBotInterfaceDll.MustFindProc("UpdateRigidBodyTickFlatbuffer"),
		getMatchSettingsProc:               rlBotInterfaceDll.MustFindProc("GetMatchSettings"),
		sendQuickChatProc:                  rlBotInterfaceDll.MustFindProc("SendQuickChat"),
		sendChatProc:                       rlBotInterfaceDll.MustFindProc("SendChat"),
		receiveChatProc:                    rlBotInterfaceDll.MustFindProc("ReceiveChat"),
		updatePlayerInputFlatbufferProc:    rlBotInterfaceDll.MustFindProc("UpdatePlayerInputFlatbuffer"),
		renderGroupProc:                    rlBotInterfaceDll.MustFindProc("RenderGroup"),
	}
	return _bridge, nil
}

// TODO - determine if locks are needed, refactor to struct to allow for locks
type bridgeWindows struct {
	isInitializedProc                  *windows.Proc
	getBallPredictionProc              *windows.Proc
	freeProc                           *windows.Proc
	setGameStateProc                   *windows.Proc
	startMatchFlatbufferProc           *windows.Proc
	updateFieldInfoFlatbufferProc      *windows.Proc
	updateLiveDataPacketFlatbufferProc *windows.Proc
	updateRigidBodyTickFlatbufferProc  *windows.Proc
	getMatchSettingsProc               *windows.Proc
	sendQuickChatProc                  *windows.Proc
	sendChatProc                       *windows.Proc
	receiveChatProc                    *windows.Proc
	updatePlayerInputFlatbufferProc    *windows.Proc
	renderGroupProc                    *windows.Proc
}

// Interface.hpp
func (b *bridgeWindows) IsInitialized() (bool, error) {
	res, _, errno := b.isInitializedProc.Call()
	if errno != syscall.Errno(0) {
		return false, errno
	}
	return res != 0, nil
}

// GameFunctions/BallPrediction.hpp
func (b *bridgeWindows) GetBallPrediction() (flat.BallPrediction, error) {
	return flat.BallPrediction{}, nil
}

// GameFunctions/GameFunctions.hpp
func (b *bridgeWindows) SetGameState(flat.DesiredGameState) error {
	return nil
}
func (b *bridgeWindows) StartMatch(flat.MatchSettings) error {
	return nil
}

// GameFunctions/GamePacket.hpp
func (b *bridgeWindows) UpdateFieldInfo() (flat.FieldInfo, error) {
	return flat.FieldInfo{}, nil
}
func (b *bridgeWindows) UpdateLiveDataPacket() (flat.GameTickPacket, error) {
	return flat.GameTickPacket{}, nil
}
func (b *bridgeWindows) UpdateRigidBodyTick() (flat.RigidBodyTick, error) {
	return flat.RigidBodyTick{}, nil
}
func (b *bridgeWindows) GetMatchSettings() (flat.MatchSettings, error) {
	return flat.MatchSettings{}, nil
}

// GameFunctions/PlayerInfo.hpp
func (b *bridgeWindows) SendQuickChat(flat.QuickChat) error {
	return nil
}
func (b *bridgeWindows) ReceiveChat() (flat.QuickChatMessages, error) {
	return flat.QuickChatMessages{}, nil
}
func (b *bridgeWindows) UpdatePlayerInput(flat.PlayerInput) error {
	return nil
}

// RenderFunctions/RenderFunctions.hpp
func (b *bridgeWindows) RenderGroup(flat.RenderGroup) error {
	return nil
}

// ensure *bridgeWindows satisfies Bridge interface
var _ Bridge = &bridgeWindows{}

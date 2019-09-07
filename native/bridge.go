package native

import (
	"github.com/nikkelma/rlbot-go/flat"
)

// Bridge is the interface implemented by platform-specific bridges
type Bridge interface {
	// Interface.hpp
	IsInitialized() bool

	// GameFunctions/BallPrediction.hpp
	GetBallPrediction() flat.BallPrediction

	// GameFunctions/GameFunctions.hpp
	SetGameState(flat.DesiredGameState) error
	StartMatch(flat.MatchSettings) error

	// GameFunctions/GamePacket.hpp
	UpdateFieldInfo() flat.FieldInfo
	UpdateLiveDataPacket() flat.GameTickPacket
	UpdateRigidBodyTick() flat.RigidBodyTick
	GetMatchSettings() flat.MatchSettings

	// GameFunctions/PlayerInfo.hpp
	SendQuickChat(flat.QuickChat) error
	ReceiveChat() flat.QuickChatMessages
	UpdatePlayerInput(flat.PlayerInput) error

	// RenderFunctions/RenderFunctions.hpp
	RenderGroup(flat.RenderGroup) error
}

// NewBridge returns a new platform-specific bridge
func NewBridge() Bridge {
	return newBridge()
}

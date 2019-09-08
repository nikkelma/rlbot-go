package native

import (
	"github.com/nikkelma/rlbot-go/flat"
)

// Bridge is the interface implemented by platform-specific bridges
type Bridge interface {
	// Interface.hpp
	IsInitialized() (bool, error)

	// GameFunctions/BallPrediction.hpp
	GetBallPrediction() (flat.BallPrediction, error)

	// GameFunctions/GameFunctions.hpp
	SetGameState(flat.DesiredGameState) error
	StartMatch(flat.MatchSettings) error

	// GameFunctions/GamePacket.hpp
	UpdateFieldInfo() (flat.FieldInfo, error)
	UpdateLiveDataPacket() (flat.GameTickPacket, error)
	UpdateRigidBodyTick() (flat.RigidBodyTick, error)
	GetMatchSettings() (flat.MatchSettings, error)

	// GameFunctions/PlayerInfo.hpp
	SendQuickChat(flat.QuickChat) error
	ReceiveChat() (flat.QuickChatMessages, error)
	UpdatePlayerInput(flat.PlayerInput) error

	// RenderFunctions/RenderFunctions.hpp
	RenderGroup(flat.RenderGroup) error
}

// NewBridge returns a new platform-specific bridge
func NewBridge() Bridge {
	return newBridge()
}

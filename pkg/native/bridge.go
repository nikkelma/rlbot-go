package native

import (
	"github.com/nikkelma/rlbot-go/pkg/flat"
)

// Bridge is the interface implemented by platform-specific bridges
type Bridge interface {
	// Close stops the bridge and attempts to unload the shared library
	Close() error

	// Interface.hpp
	IsInitialized() (bool, error)

	// GameFunctions/BallPrediction.hpp
	GetBallPrediction() (*flat.BallPrediction, error)

	// GameFunctions/GameFunctions.hpp
	SetGameState(*flat.DesiredGameState) error
	StartMatch(*flat.MatchSettings) error

	// GameFunctions/GamePacket.hpp
	GetFieldInfo() (*flat.FieldInfo, error)
	GetLiveGameTickPacket() (*flat.GameTickPacket, error)
	GetMatchSettings() (*flat.MatchSettings, error)

	// GameFunctions/PlayerInfo.hpp
	SendQuickChat(*flat.QuickChat) error
	ReceiveChat(botIndex, teamIndex, lastMessageIndex int) (*flat.QuickChatMessages, error)
	UpdatePlayerInput(*flat.PlayerInput) error

	// RenderFunctions/RenderFunctions.hpp
	RenderGroup(*flat.RenderGroup) error
}

// NewBridge returns a new platform-specific bridge
func NewBridge() (Bridge, error) {
	return newBridge()
}

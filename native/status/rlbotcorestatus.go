//go:generate stringer -type=RLBotCoreStatus

package status

// RLBotCoreStatus represents the RLBotCoreStatus enum in C
type RLBotCoreStatus int

const (
  Success RLBotCoreStatus = iota
  BufferOverfilled
  MessageLargerThanMax
  InvalidNumPlayers
  InvalidBotSkill
  InvalidHumanIndex
  InvalidName
  InvalidTeam
  InvalidTeamColorID
  InvalidCustomColorID
  InvalidGameValues
  InvalidThrottle
  InvalidSteer
  InvalidPitch
  InvalidYaw
  InvalidRoll
  InvalidPlayerIndex
  InvalidQuickChatPreset
  InvalidRenderType
  QuickChatRateExceeded
  NotInitialized
)

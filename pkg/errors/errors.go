package errors

import (
	"fmt"

	"github.com/elojah/game_01/pkg/geometry"
	perrors "github.com/pkg/errors"
)

// #Dev errors

// ErrNotImplementedYet is raised when a resource is not implemented yet.
type ErrNotImplementedYet struct {
	Version string
}

func (err ErrNotImplementedYet) Error() string {
	return fmt.Sprintf("not implemented in version %s", err.Version)
}

// #Token/login errors

// ErrWrongIP is raised when a source IP doesn't match with token-associated IP.
type ErrWrongIP struct {
	TokenID  string
	Expected string
	Actual   string
}

func (err ErrWrongIP) Error() string {
	return fmt.Sprintf("token %s ip is %s, received %s", err.TokenID, err.Expected, err.Actual)
}

// ErrInvalidTS is raised when a packet has a TS out of valid range.
type ErrInvalidTS struct {
	MsgID string
	TS    uint64
	Now   uint64
}

func (err ErrInvalidTS) Error() string {
	return fmt.Sprintf("msg %s with TS %d received too late at %d", err.MsgID, err.TS, err.Now)
}

// ErrWrongCredentials is raised when user logs with invalid username/password.
type ErrWrongCredentials struct {
	Username string
}

func (err ErrWrongCredentials) Error() string {
	return fmt.Sprintf("invalid credentials for user %s", err.Username)
}

// ErrInsufficientACLs is raised when a user apply an action without valid rights.
type ErrInsufficientACLs struct {
	Value  int
	Source string
	Target string
}

func (err ErrInsufficientACLs) Error() string {
	return fmt.Sprintf("insufficient rights %d for source %s on target %s", err.Value, err.Source, err.Target)
}

// ErrMultipleLogin is raised zhen an account is already logged.
type ErrMultipleLogin struct {
	AccountID string
}

func (err ErrMultipleLogin) Error() string {
	return fmt.Sprintf("account %s already logged", err.AccountID)
}

// ErrInvalidEntityType is raised when an entity doesn't respect the correct type.
type ErrInvalidEntityType struct {
	EntityType string
}

func (err ErrInvalidEntityType) Error() string {
	return fmt.Sprintf("invalid entity type %s", err.EntityType)
}

// #Common api/core errors

// ErrInvalidMove is raised when an invalid move event is processed.
type ErrInvalidMove struct {
	TargetID string

	SectorID       string
	SectorDim      geometry.Vec3
	TargetPosition geometry.Vec3

	NewSectorID  string
	NewSectorDim geometry.Vec3
	NewPosition  geometry.Vec3
}

func (err ErrInvalidMove) Error() string {
	return fmt.Sprintf(
		"invalid move from sector %s (%f, %f, %f) to sector %s (%f, %f, %f) for entity %s from (%f, %f, %f) to (%f, %f, %f)",
		err.SectorID,
		err.SectorDim.X,
		err.SectorDim.Y,
		err.SectorDim.Z,
		err.NewSectorID,
		err.NewSectorDim.X,
		err.NewSectorDim.Y,
		err.NewSectorDim.Z,
		err.TargetID,
		err.TargetPosition.X,
		err.TargetPosition.Y,
		err.TargetPosition.Z,
		err.NewPosition.X,
		err.NewPosition.Y,
		err.NewPosition.Z,
	)
}

// ErrInvalidNeighbourSector is raised when a sector is not an expected neighbour.
type ErrInvalidNeighbourSector struct {
	SectorID        string
	SectorNeighbour string
}

func (err ErrInvalidNeighbourSector) Error() string {
	return fmt.Sprintf("sector %s is not neighbour to %s", err.SectorNeighbour, err.SectorID)
}

// ErrFullPCCreated is raised when a pc is created but the account limit for pc is already reached.
type ErrFullPCCreated struct {
	AccountID string
}

func (err ErrFullPCCreated) Error() string {
	return fmt.Sprintf("account %s cannot create anymore pc", err.AccountID)
}

// ErrMissingMP is raised when an entity cast a skill with not enough mp to perform it.
type ErrMissingMP struct {
	EntityID      string
	AbilityID     string
	MPLeft        uint64
	MPConsumption uint64
}

func (err ErrMissingMP) Error() string {
	return fmt.Sprintf(
		"entity %s (%d MP) cannot cast %s (%d MP)",
		err.EntityID,
		err.MPLeft,
		err.AbilityID,
		err.MPConsumption,
	)
}

// ErrAbilityCDDown is raised when an ability is cast but previous CD is still up.
type ErrAbilityCDDown struct {
	EntityID  string
	AbilityID string
	TS        uint64
	LastUsed  uint64
	CD        uint64
}

func (err ErrAbilityCDDown) Error() string {
	return fmt.Sprintf(
		"entity %s cannot cast ability %s at %d last used %d + CD %d",
		err.EntityID,
		err.AbilityID,
		err.TS,
		err.LastUsed,
		err.CD,
	)
}

// ErrNotFound is raised when a mandatory resource is not found in storage.
type ErrNotFound struct {
	Store string
	Index string
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("no results found in store %s for index %s", err.Store, err.Index)
}

// ErrMissingTarget is raised when a ability is performed with a missing component targets.
type ErrMissingTarget struct {
	AbilityID   string
	ComponentID string
}

func (err ErrMissingTarget) Error() string {
	return fmt.Sprintf("missing target for ability %s component %s", err.AbilityID, err.ComponentID)
}

// ErrTooManyTargets is raised when a ability is casted/performed with too many targets for a component.
type ErrTooManyTargets struct {
	NTargets    int
	Max         uint64
	AbilityID   string
	ComponentID string
}

func (err ErrTooManyTargets) Error() string {
	return fmt.Sprintf("too many targets %d for ability %s component %s", err.NTargets, err.AbilityID, err.ComponentID)
}

// ErrOutOfRange is raised when a component ability is performed out of range on a target or when an item is looted in a further position.
type ErrOutOfRange struct {
	Dist  float64
	Range float64
}

func (err ErrOutOfRange) Error() string {
	return fmt.Sprintf("out of range %f for %f", err.Dist, err.Range)
}

// ErrMissingItem is raised when a loot action is performed on an non present item in target inventory.
type ErrMissingItem struct {
	ItemID      string
	InventoryID string
}

func (err ErrMissingItem) Error() string {
	return fmt.Sprintf("item %s missing in inventory %s", err.ItemID, err.InventoryID)
}

// ErrFullInventory is raised when a loot action is performed and the source inventory is full.
type ErrFullInventory struct {
	InventoryID string
}

func (err ErrFullInventory) Error() string {
	return fmt.Sprintf("inventory %s full", err.InventoryID)
}

// ErrIneffectiveCancel is raised when a cancel event is applied but no event was found to cancel.
type ErrIneffectiveCancel struct {
	TriggerID string
}

func (err ErrIneffectiveCancel) Error() string {
	return fmt.Sprintf("ineffective cancel trigger %s", err.TriggerID)
}

// IsGameLogicError returns if error type is a game logic error and needs a cancel propagation.
func IsGameLogicError(err error) bool {
	switch err := perrors.Cause(err).(type) {
	case ErrNotImplementedYet:
	case ErrWrongIP:
	case ErrInvalidTS:
	case ErrWrongCredentials:
	case ErrInsufficientACLs:
	case ErrMultipleLogin:
	case ErrInvalidEntityType:
	case ErrNotFound:
	case ErrMissingTarget:
	case ErrTooManyTargets:
	case ErrOutOfRange:
	case ErrMissingItem:
	case ErrFullInventory:
	case ErrIneffectiveCancel:
		_ = err
	}

	return false
}

package touch

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/exp/maps"
)

var (
	allTouchIDs          []ebiten.TouchID
	currentTouchIDs      map[ebiten.TouchID]bool
	justPressedTouchIDs  map[ebiten.TouchID]bool
	justReleasedTouchIDs map[ebiten.TouchID]bool
)

func UpdateTouchIDs() {
	newPressedTouchIDs := []ebiten.TouchID{}
	newPressedTouchIDs = inpututil.AppendJustPressedTouchIDs(newPressedTouchIDs)
	justPressedTouchIDs = map[ebiten.TouchID]bool{}
	for i := 0; i < len(newPressedTouchIDs); i++ {
		justPressedTouchIDs[newPressedTouchIDs[i]] = true
		currentTouchIDs[newPressedTouchIDs[i]] = true
	}
	justReleasedTouchIDs = map[ebiten.TouchID]bool{}
	allTouchIDs = maps.Keys(currentTouchIDs)
	newReleasedTouchIDs := []ebiten.TouchID{}
	newReleasedTouchIDs = inpututil.AppendJustReleasedTouchIDs(newReleasedTouchIDs)
	for i := 0; i < len(newReleasedTouchIDs); i++ {
		justReleasedTouchIDs[newReleasedTouchIDs[i]] = true
		delete(currentTouchIDs, newReleasedTouchIDs[i])
	}
}

func GetTouchIDs() []ebiten.TouchID {
	return allTouchIDs
}

func IsTouchJustPressed(touchID ebiten.TouchID) bool {
	return justPressedTouchIDs[touchID]
}

func IsTouchJustReleased(touchID ebiten.TouchID) bool {
	return justReleasedTouchIDs[touchID]
}

func init() {
	allTouchIDs = []ebiten.TouchID{}
	currentTouchIDs = map[ebiten.TouchID]bool{}
	justPressedTouchIDs = map[ebiten.TouchID]bool{}
	justReleasedTouchIDs = map[ebiten.TouchID]bool{}
}

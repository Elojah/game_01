package pkg

import (
	"github.com/elojah/game_01/pkg/account"
	"github.com/elojah/game_01/pkg/entity"
	"github.com/elojah/game_01/pkg/event"
)

// Token wraps use cases around token object.
type Token struct {
	account.TokenMapper
	EntityMapper entity.Mapper
	event.QRecurrerMapper
}

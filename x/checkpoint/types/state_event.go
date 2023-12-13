package types

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/stateevent"
)

type StateEvent interface {
	Subscribe(subscriber stateevent.EventSubscriber)
}

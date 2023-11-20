package types

import (
	stateevent "github.com/PlatONnetwork/AppChain-SDK/x/state_event"
)

type StateEvent interface {
	Subscribe(subscriber stateevent.EventSubscriber)
}

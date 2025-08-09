package notifier

import "lld/behavioral_design_pattens/observer/observer"

type Notifier interface {
	Subscribe(o observer.Observer)
	Unsubscribe(o observer.Observer)
	NotifyAll()
}

package input

import (
	"github.com/clktmr/n64/drivers/controller"
)

// Manager handles controller input polling
type Manager struct {
	states chan [4]controller.Controller
}

// NewManager creates a new controller manager
func NewManager() *Manager {
	cm := &Manager{
		states: make(chan [4]controller.Controller),
	}
	cm.start()
	return cm
}

// start begins polling controllers
func (cm *Manager) start() {
	go func() {
		var states [4]controller.Controller
		for {
			controller.Poll(&states)
			cm.states <- states
		}
	}()
}

// GetInputs returns the current controller states
func (cm *Manager) GetInputs() [4]controller.Controller {
	return <-cm.states
}

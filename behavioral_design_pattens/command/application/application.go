package application

import (
	"lld/behavioral_design_pattens/command/commands"
	"lld/behavioral_design_pattens/command/history"
)

type Application struct {
	historyInstance *history.CommandHistory
}

func NewApplication(history *history.CommandHistory) *Application {
	return &Application{
		historyInstance: history,
	}
}

func (app *Application) ExecuteCommand(cmd commands.Command) {
	if cmd.Execute() {
		app.historyInstance.Push(cmd)
	}
}

func (app *Application) Undo() {
	if cmd, err := app.historyInstance.Pop(); err == nil {
		cmd.Undo()
	}
}

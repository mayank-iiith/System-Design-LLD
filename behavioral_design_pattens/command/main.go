package main

// Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations.

import (
	"fmt"
	"lld/behavioral_design_pattens/command/application"
	"lld/behavioral_design_pattens/command/commands"
	"lld/behavioral_design_pattens/command/editor"
	"lld/behavioral_design_pattens/command/history"
)

func main() {
	editorInstance := editor.NewEditor()
	historyInstance := history.NewCommandHistory()
	app := application.NewApplication(historyInstance)

	editorInstance.Text = "Hello, World!"
	fmt.Println("Editor text:", editorInstance.Text)

	app.ExecuteCommand(&commands.CutCommand{
		BaseCommand: commands.BaseCommand{
			Editor: editorInstance,
		},
	})
	fmt.Println("Editor text:", editorInstance.Text)

	app.Undo()
	fmt.Println("Editor text (after undo):", editorInstance.Text)

	editorInstance.Text = "Hello, Mayank!"
	fmt.Println("Editor text:", editorInstance.Text)

	app.ExecuteCommand(&commands.PasteCommand{
		BaseCommand: commands.BaseCommand{
			Editor: editorInstance,
		},
	})
	fmt.Println("Editor text (after paste):", editorInstance.Text)
}

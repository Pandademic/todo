package main

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    berry "github.com/pandademic/raspberry"
    "os"
    "github.com/mitchellh/go-homedir"
)

type Todo struct {
    dueDate string
    name    string
    done    bool
}

var app = tview.NewApplication()

func AddTodoForm() {
  newTodo := Todo{}

  var form = tview.NewForm()

  form.AddInputField("Name","",20,nil,func(name string){
      newTodo.name = name
      newTodo.done = false
  })

  form.AddInputField("Due Date","",20,nil,func(dueDate string){
    newTodo.dueDate = dueDate
  })

  form.AddButton("Save",func(){
      f , err := os.OpenFile(string(homedir.Dir())+string(os.PathSeparator)+"todo"+string(os.PathSeparator)+"todos",
           os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
      if err != nil {
        panic(err)
      }
      defer f.Close()

      if _, err := f.WriteString(string(newTodo)); err != nil {
          panic(err)
      }
  })
}


func tuiMain() {
  var quitmsg = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("use Ctrl-c to quit")
  app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    if event.Rune() == 113 {
        app.Stop()
    }
    return event
  })
  if err := app.SetRoot(quitmsg,true).EnableMouse(true).Run(); err != nil {
    panic(err)
  }
}

func main() {
  cli := berry.Cli{AcceptedCommands:[]string{"tui"},HelpMsg:"Use 'todo tui' to see the tui",Version: 0.1}
  cli.Setup()
  cli.SetHandler("tui",tuiMain)

}


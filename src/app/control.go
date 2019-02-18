package app

import(
  "github.com/nsf/termbox-go"
  "github.com/ryomak/onepass/src/display"
)

func insertLoop(){
  insertloop:
  for{
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventResize:
      display.DisplayMainFrame()
    case termbox.EventKey:
      switch ev.Key {
      case termbox.KeyEsc:
        break insertloop
      case termbox.KeyCtrlC:
        break insertloop
      }
    }
  }
}

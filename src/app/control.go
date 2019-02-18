package app

import(
  "github.com/nsf/termbox-go"
  "github.com/ryomak/onepass/src/display"
  "github.com/ryomak/onepass/src/util"
)

func insertLoop(){
  insertloop:
  for{
    display.DisplayMainFrame()
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventResize:
      display.DisplayMainFrame()
    case termbox.EventKey:
      switch ev.Key {
      case termbox.KeyEsc:
        util.Mode = util.NORMALMODE
        break insertloop
      case termbox.KeyCtrlC:
        util.Mode = util.NORMALMODE
        break insertloop
      }
    }
    display.DisplayMainFrame()
  }
}

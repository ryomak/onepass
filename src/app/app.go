package app

import(
  "github.com/ryomak/onepass/src/display"
  "github.com/ryomak/onepass/src/util"
  termbox"github.com/nsf/termbox-go"
  "log"
)

const(
  UP = iota
  Down
  Left
  Right
)


type cursor struct {
  x int
  y int
}


func Run(){
  if err := termbox.Init(); err != nil {
    log.Fatal(err)
  }
  defer termbox.Close()
  display.DisplayMainFrame()
  mainloop:
  for{
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventResize:
      display.DisplayMainFrame()
    case termbox.EventKey:
      switch ev.Key {
      case termbox.KeyEsc:
        break mainloop
      case termbox.KeyCtrlC:
        break mainloop
      case termbox.KeyCtrlI:
        util.Mode = util.CREATEMODE
        insertLoop()
      }
    }
    display.DisplayMainFrame()
  }
}



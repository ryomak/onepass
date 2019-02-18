package app

import(
  "github.com/ryomak/onepass/src/display"
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
  display.DisplayFrame()
  mainloop:
  for{
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventResize:
      display.DisplayFrame()
    case termbox.EventKey:
      switch ev.Key {
      case termbox.KeyEsc:
        break mainloop
      case termbox.KeyCtrlC:
        break mainloop
      }
    }
    display.DisplayFrame()
  }
}


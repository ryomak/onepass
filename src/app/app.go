package app

import(
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
  keyEventCh := make(chan termbox.Key)
  charEventCh := make(chan rune)
  go keyEventLoop(keyEventCh,charEventCh)
  control(keyEventCh,charEventCh)
}

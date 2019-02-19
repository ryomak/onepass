package app

import (
  termbox"github.com/nsf/termbox-go"
)
func keyEventLoop(keyEventCh chan termbox.Key){
  for{
    switch ev := termbox.PollEvent(); ev.Type {
    case termbox.EventKey:
      keyEventCh <- ev.Key
    }
  }
}

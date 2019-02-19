package app

import(
  "github.com/nsf/termbox-go"
  "github.com/ryomak/onepass/src/display"
  "github.com/ryomak/onepass/src/util"
  "log"
)

func control(keyCh chan termbox.Key,charCh chan rune){
  mainloop:
  for{
    display.DisplayAll()
    select{
    case key := <-keyCh:
      switch key {
      case termbox.KeyEsc:
        break mainloop
      case termbox.KeyCtrlC:
        break mainloop
      case termbox.KeyCtrlI:
        util.Mode = util.CREATEMODE
        insertLoop(keyCh)
      }
    case chr := <-charCh:
        log.Println(chr)
    }
  }
}

func insertLoop(keyCh chan termbox.Key){
  insertloop:
  for{
    display.DisplayAll()
    select{
    case key := <-keyCh:
      switch key {
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

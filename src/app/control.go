package app

import(
  "github.com/nsf/termbox-go"
  "github.com/ryomak/onepass/src/display"
  "github.com/ryomak/onepass/src/util"
)

func control(keyCh chan termbox.Key){
  mainloop:
  for{
    display.DisplayMainFrame()
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
    }
  }
}

func insertLoop(keyCh chan termbox.Key){
  insertloop:
  for{
    display.DisplayMainFrame()
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

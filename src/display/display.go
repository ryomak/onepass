package display
import(
  termbox"github.com/nsf/termbox-go"
  "github.com/ryomak/onepass/src/model"
)
var TerminalWidth,TerminalHeight int = 0,0
var panes = map[string]Pane{}

func init(){
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  panes["AccountListPane"] = new(AccountListPane)
  panes["ModePane"] = new(ModePane)
}

func DisplayMainFrame(){
  TerminalWidth ,TerminalHeight = termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  //縦のライン
  writeVerticalLine(panes["AccountListPane"].MaxX(),0,LineOption{
    Length:panes["ModePane"].Y(),
  })
  writeLine(0,panes["ModePane"].Y(),LineOption{})
  DisplayAccount([]model.Account{
    model.Account{Name:"ryoma"} ,
    model.Account{Name:"ryoma"} ,
    model.Account{Name:"ryoma"} ,
    model.Account{Name:"ryoma"} ,
  })
	termbox.Flush()
}

func DisplayAccount(acc []model.Account){
  for i,v := range acc{
    writeStr(v.Name,SentenceOption{
      X:1,
      Y:i+1,
      MaxX:panes["AccountListPane"].MaxX(),
    })
  }
}




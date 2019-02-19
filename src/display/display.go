package display
import(
  termbox"github.com/nsf/termbox-go"
  "github.com/ryomak/onepass/src/model"
  "github.com/ryomak/onepass/src/util"
  "github.com/mattn/go-runewidth"
)
var TerminalWidth,TerminalHeight int = 0,0
var panes = map[string]Pane{}

func init(){
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  panes["AccountListPane"] = new(AccountListPane)
  panes["ModePane"] = new(ModePane)
}

func DisplayAll(){
  DisplayMainFrame()
  DisplayMode(util.EchoMode())
  termbox.Flush()
}

func DisplayMainFrame(){
  TerminalWidth ,TerminalHeight = termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  //縦のライン
  writeVerticalLine(panes["AccountListPane"].MaxX(),0,LineOption{
    Length:panes["ModePane"].Y(),
  })
  //下のライン
  writeLine(0,panes["ModePane"].Y(),LineOption{})
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

func DisplayMode(mode string){
  writeStr(mode,SentenceOption{
    X:panes["ModePane"].MaxX()-runewidth.StringWidth(mode)-1,
    Y:panes["ModePane"].Y()+1,
    MaxX:panes["ModePane"].MaxX(),
  })
}



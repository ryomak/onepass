package display

import(
  termbox"github.com/nsf/termbox-go"
  "github.com/mattn/go-runewidth"
)

type LineOption struct{
  Reverse bool
  Length  int
  CharColor termbox.Attribute
  BGColor termbox.Attribute
}

type SentenceOption struct{
  X int
  Y int
  MaxX int
  MaxY int
  CharColor termbox.Attribute
  BGColor termbox.Attribute
}


func writeLine(x,y int,op LineOption){
  width := TerminalWidth
  if !op.Reverse {
    if op.Length != 0 && width > (x+op.Length){
        width = x+op.Length
    }
    for col:=x;col < width;col++{
      termbox.SetCell(col, y, '-', op.CharColor, op.BGColor)
    }
  }else{
    if op.Length != 0 {
      width = x-op.Length
      if width < 0{
        width = 0
      }
    }
    for col:=x;col > width;col--{
      termbox.SetCell(col, y, '-', op.CharColor, op.BGColor)
    }
  }
}

func writeVerticalLine(x,y int , op LineOption){
  height := TerminalHeight
  if !op.Reverse {
    if op.Length != 0 && height > (y+op.Length){
      height = y+op.Length
    }
    for row :=y ;row < height;row++{
      termbox.SetCell(x, row, '|', op.CharColor, op.BGColor)
    }
  }else{
    if op.Length != 0 {
      height = x-op.Length
      if height < 0{
        height = 0
      }
    }
    for row:=y ; row > height ;row--{
      termbox.SetCell(x, row, '|', op.CharColor, op.BGColor)
    }
  }
}

func writeStr(str string ,op SentenceOption){
  runes := []rune(str)
  for _,r := range runes{
		termbox.SetCell(op.X, op.Y, r, op.CharColor, op.BGColor)
    op.X += runewidth.RuneWidth(r)
    if(op.X) > op.MaxX{return}
	}
}

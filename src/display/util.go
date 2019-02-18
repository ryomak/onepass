package display

import(
  termbox"github.com/nsf/termbox-go"
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
      height = x+op.Length
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
	for col := 0; col < len(str); col++ {
    if(op.X+col) > op.MaxX{return}
		char := rune(str[col])
		termbox.SetCell(op.X+col, op.Y, char, op.CharColor, op.BGColor)
	}
}

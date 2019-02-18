package app

const (
  NORMAL = 1
  CREATE =-1
)

var Mode = NORMAL

func ChangeMode(mode *int){
  if mode == nil{
    Mode = Mode*-1
    return
  }
  Mode = *mode
}

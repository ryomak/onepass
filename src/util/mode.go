package util

const (
  NORMALMODE = 1
  CREATEMODE =-1
)

var Mode = NORMALMODE

func EchoMode()string{
  switch(Mode){
  case NORMALMODE:
    return "NORMALMODE"
  case CREATEMODE:
    return "CREATEMODE"
  }
  return "Error"
}

func ChangeMode(mode *int){
  if mode == nil{
    Mode = Mode*-1
    return
  }
  Mode = *mode
}

package display

import(
)


type Pane interface{
  X()int
  Y()int
  MaxX()int
  MaxY()int
}

type AccountListPane struct{}
func(a AccountListPane)X()int{return 0}
func(a AccountListPane)Y()int{return 0}
func(a AccountListPane)MaxX()int{return TerminalWidth/3}
func(a AccountListPane)MaxY()int{return TerminalWidth}

type ModePane struct{}
func(a ModePane)X()int{return 0}
func(a ModePane)Y()int{return TerminalHeight-2}
func(a ModePane)MaxX()int{return TerminalWidth}
func(a ModePane)MaxY()int{return TerminalHeight}

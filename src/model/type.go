package model

import(
  "time"
)

type Account struct{
  ID uint
  CreatedAt time.Time
  UpdatedAt time.Time
  DeleteAt *time.Time
  Name string
  User string
  Pass string
  Discription string
  Type string
}

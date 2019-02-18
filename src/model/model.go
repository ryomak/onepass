package model

type Model interface{
  GetAccounts()([]Account,error)
  CreateAccount(*Account)error
  UpdateAccount(*Account)error
  DeleteAccount(id uint)error
}


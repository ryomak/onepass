package model

type DB struct{

}


func (d DB)GetAccunts()([]Account,error){
  return nil,nil
}
func (d DB)CreateAccunts(acc *Account)error{
  return nil
}
func(d DB) UpdateAccount(*Account)error{
  return nil
}
func(d DB) DeleteAccount(id uint)error{
  return nil
}

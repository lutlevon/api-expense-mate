package models

type Bill struct {
	ID          uint `gorm:"column:id;primaryKey"`
	UserID      uint `gorm:"column:userid;unique"`
	BillTypeID  uint `gorm:"column:billtypeid;unique"`
	TotalAmount uint `gorm:"column:totalamount"`
}

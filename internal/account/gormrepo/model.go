package gormrepo

type AccountGorm struct {
	ID      string `gorm:"type:uuid;column:id;primaryKey"`
	Name    string `gorm:"type:varchar(255);column:name;not null"`
	Balance int    `gorm:"type:int;column:balance;not null"`
}

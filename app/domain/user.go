package domain

type User struct {
	Id       int64  `gorm:"column_name: id type: bigint(20) not null auto_increment; primary_key" json:"id"`
	Account  string `gorm:"column_name: account type: varchar(50) not null; index:idx_account,unique" json:"account"`
	Password string `gorm:"column_name: password type: varchar(50) not null" json:"-"`
	Status   string `gorm:"column_name: statustype: varchar(1) not null default '0';" json:"status"`
}

func (User) TableName() string {
	return "user"
}

type IUserRepository interface {
	FindById(id *int64) (*User, error)
	FindAllUsers() ([]*User, error)
	CreateUser(user *User) (id *int64, error error)
	UpdateUser(user *User) error
	DeleteUser(id *int64) error
	UnlockUser(id *int64) error
}

type IUserUsecase interface {
	FindById(id *int64) (*User, error)
	FindUsers() ([]*User, error)
	CreateUser(user *User) (id *int64, error error)
	UpdateUser(user *User) error
	DeleteUser(id *int64) error
	UnlockUser(id *int64) error
}

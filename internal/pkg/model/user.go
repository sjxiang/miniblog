package model

type UserM struct {

}

// TableName sets the insert table name for this struct type
func (u *UserM) TableName() string {
	return "user"
}



package model

type PostM struct {
}

// TableName sets the insert table name for this struct type
func (u *PostM) TableName() string {
	return "post"
}

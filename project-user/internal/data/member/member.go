package member

type Member struct {
	Id   int64
	Name string
	Pwd  string
}

type Rg struct {
	Id         int64
	Name       string
	Password   string
	CreateTime string
}

func (*Member) TableName() string {
	return "user"
}

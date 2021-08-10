package cmd

//模拟用户启动方式
func init() {
	globalObjectAPI = &User{}
}

type User struct {
}

func (this *User) A1(a string) string {
	return a
}

type Teacher struct {
}

func (this *Teacher) A1(a string) string {
	return a
}

package kgin

const (
	CtxKeyUser = "user"
)

type User struct {
	TalId   string
	Version string
	GradeId int
}

func (c *Context) SetUser(user *User) {
	c.Set(CtxKeyUser, user)
}

func (c *Context) GetUser() *User {
	if user, ok := c.Get(CtxKeyUser); ok {
		return user.(*User)
	}
	return nil
}

package req

type User struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

//SingUp 注册参数
type SingUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//Login 登录参数
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	//RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

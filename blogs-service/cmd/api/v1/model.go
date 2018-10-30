package v1

type SignUpStruct struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Sex      int64  `json:"sex"`
	TelNum   string `json:"tel_num"`
}

type SignInStruct struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

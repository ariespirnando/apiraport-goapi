package model

type Requestraportdet struct {
	Kodejuruan string `json:"kodejurusan"`
	Kodekelas  string `json:"kodekelas"`
	Guidraport string `json:"guidraport"`
}

type Requestlgoin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

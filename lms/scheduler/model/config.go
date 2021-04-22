package model

type ConfigLms struct {
	Login string     `json:"login"`
	Psw   string     `json:"psw"`
	Ids   []ConfigId `json:"ids"`
}

type ConfigId struct {
	Lms   int    `json:"lms"`
	Gpee  int    `json:"gpee"`
	Login string `json:"login"`
	Psw   string `json:"psw"`
}

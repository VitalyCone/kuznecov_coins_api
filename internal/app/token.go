package app

var CurrentToken string

type TokenData struct{
	CurrentToken string `yaml:"token"`
}

func (s *TokenData) Init(){
	CurrentToken = s.CurrentToken
}


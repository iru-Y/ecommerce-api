package dto

type UserDto struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Cpf      string `json:"cpf"`
	City     string `json:"city"`
	State    string `json:"state"`
	Address  string `json:"address"`
	Password []byte `json:"password"`
	Mail     string `json:"mail"`
	Role     string `json:"role"`
}

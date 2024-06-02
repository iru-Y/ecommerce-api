package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	LastName     string             `json:"lastName"`
	Cpf          string             `json:"cpf"`
	City         string             `json:"city"`
	State        string             `json:"state"`
	Address      string             `json:"address"`
	Password     []byte             `json:"password"`
	Mail         string             `json:"mail"`
	Role         string             `json:"role"`
	ProfileImage string             `json:"profileImage"`
}

func NoArgsUserModel(name, lastName, cpf, city, state, address, mail, role, profileImage string, password []byte) *UserModel {
	return &UserModel{
		Name:         "",
		LastName:     "",
		Cpf:          "",
		City:         "",
		State:        "",
		Address:      "",
		Password:     []byte("0"),
		Mail:         "",
		Role:         "",
		ProfileImage: "",
	}
}

func NewUserModel(name, lastName, cpf, city, state, address, mail, role string, password []byte,
) *UserModel {
	return &UserModel{
		Name:     name,
		LastName: lastName,
		Cpf:      cpf,
		City:     city,
		State:    state,
		Address:  address,
		Password: password,
		Mail:     mail,
		Role:     role,
	}
}

func NewUserModeWithProfileImage(name, lastName, cpf, city, state, address, mail, role, profileImage string, password []byte,
) *UserModel {
	return &UserModel{
		Name:         name,
		LastName:     lastName,
		Cpf:          cpf,
		City:         city,
		State:        state,
		Address:      address,
		Password:     password,
		Mail:         mail,
		Role:         role,
		ProfileImage: profileImage,
	}
}

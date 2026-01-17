package dto

import model "gen-concept-api/domain/model"

type TokenDetail struct {
	AccessToken            string     `json:"accessToken"`
	RefreshToken           string     `json:"refreshToken"`
	AccessTokenExpireTime  int64      `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64      `json:"refreshTokenExpireTime"`
	User                   model.User `json:"user"`
}

type RegisterUserByUsername struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func ToUserModel(from RegisterUserByUsername) model.User {
	return model.User{Username: from.Username,
		FirstName: from.FirstName,
		LastName:  from.LastName,
		Email:     from.Email,
	}
}

type RegisterLoginByMobile struct {
	MobileNumber string
	Otp          string
}

type LoginByUsername struct {
	Username string
	Password string
}

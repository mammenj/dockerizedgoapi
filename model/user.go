package model

import (
	"github.com/jinzhu/gorm"
)

type Users []*User
type Roles []Role

type User struct {
	gorm.Model
	Userid   string
	Name     string
	Password string
	Roles    []Role `gorm:"foreignkey:Userid;association_foreignkey:id"`
}

type Role struct {
	gorm.Model
	Roleid string
	Userid int64
}

type UsersDtos []*UserDto
type RolesDtos []RoleDto

type UserDto struct {
	ID       uint   `json:"id"`
	Userid   string `json:"userid"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Roles    []Role `json:"roles"`
}

type RoleDto struct {
	ID     uint   `json:"id"`
	Roleid string `json:"roleid"`
}

func (u User) ToDto() *UserDto {
	return &UserDto{
		ID:       u.ID,
		Userid:   u.Userid,
		Name:     u.Name,
		Password: u.Password,
		Roles:    u.Roles,
	}
}

func (r Role) ToDto() RoleDto {
	return RoleDto{
		ID:     r.ID,
		Roleid: r.Roleid,
	}
}

func (us Users) ToDto() UsersDtos {
	dtos := make([]*UserDto, len(us))
	for i, u := range us {
		dtos[i] = u.ToDto()
	}

	return dtos
}

func (rs Roles) ToDto() RolesDtos {
	dtos := make([]RoleDto, len(rs))
	for i, r := range rs {
		dtos[i] = r.ToDto()
	}

	return dtos
}

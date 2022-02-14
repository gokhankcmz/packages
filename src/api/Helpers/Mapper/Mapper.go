package Mapper

import (
	"Packages/src/api/Type/DtoTypes"
	"Packages/src/api/Type/EntityTypes"
)

func GetUserFromCreateUserDto(UserDto *DtoTypes.CreateUserDto) *EntityTypes.User {
	return &EntityTypes.User{
		Name:     UserDto.Name,
		Email:    UserDto.Email,
		Age:      UserDto.Age,
		Document: EntityTypes.Document{},
	}
}

func GetUserFromUpdateUserDto(UserDto *DtoTypes.UpdateUserDto, UserEntity *EntityTypes.User) {
	UserEntity.Name = UserDto.Name
	UserEntity.Email = UserDto.Email
	UserEntity.Age = UserDto.Age
}

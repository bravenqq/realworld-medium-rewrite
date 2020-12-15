// Package domain provides ...
package domain

import "errors"

//User ....
type User struct {
	ID       NUUID
	Name     string
	Email    Email
	Password string
	Phone    PhoneNumber
}

func (u User) Check() error {
	if len(u.Name) == 0 {
		return errors.New("请输入用户名")
	}
	if len(u.Password) == 0 {
		return errors.New("请输入密码")
	}
	if err := u.Email.Check(); err != nil {
		return err
	}
	return nil
}

// PhoneNumber 类型
// Derived: PhoneNumber, TelPhoneNumber, MobilePhoneNumber
type PhoneNumber string

func (phone PhoneNumber) Check() error {
	if len(phone) == 0 {
		return errors.New("请输入电话号码")
	}
	return nil
}

//Email
type Email string

func (e Email) Check() error {
	if len(e) == 0 {
		return errors.New("请输入邮箱")
	}
	return nil
}

//UserRepository ....
type UserRepository interface {
	FindByPhone(phone PhoneNumber) (User, error)
	Create(u User) error
	GetByEmail(e Email) (User, error)
}

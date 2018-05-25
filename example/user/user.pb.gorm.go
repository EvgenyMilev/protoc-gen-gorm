// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	example/user/user.proto

It has these top-level messages:
	User
	Email
	Address
	Language
	CreditCard
	Task
*/
package user

import context "context"
import errors "errors"
import time "time"

import gorm "github.com/jinzhu/gorm"
import ops "github.com/infobloxopen/atlas-app-toolkit/gorm"
import ptypes "github.com/golang/protobuf/ptypes"

import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type UserORM struct {
	BillingAddress    *AddressORM `gorm:"foreignkey:BillingAddressId;association_foreignkey:Id"`
	BillingAddressId  int32
	Birthday          time.Time
	CreatedAt         time.Time
	CreditCard        *CreditCardORM `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	Emails            []*EmailORM    `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	Friends           []*UserORM     `gorm:"many2many:user_friend;foreignkey:Id;association_foreignkey:Id;jointable_foreignkey:user_id;association_jointable_foreignkey:friend_id"`
	Id                int32
	Languages         []*LanguageORM `gorm:"many2many:user_language;foreignkey:Id;association_foreignkey:Id;jointable_foreignkey:user_id;association_jointable_foreignkey:language_id"`
	Num               uint32
	ShippingAddress   *AddressORM `gorm:"foreignkey:ShippingAddressId;association_foreignkey:Id"`
	ShippingAddressId int32
	Tasks             []*TaskORM `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	UpdatedAt         time.Time
}

// TableName overrides the default tablename generated by GORM
func (UserORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *User) ToORM(ctx context.Context) (UserORM, error) {
	to := UserORM{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
	}
	if m.Birthday != nil {
		if to.Birthday, err = ptypes.Timestamp(m.Birthday); err != nil {
			return to, err
		}
	}
	to.Num = m.Num
	if m.CreditCard != nil {
		tempCreditCard, err := m.CreditCard.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.CreditCard = &tempCreditCard
	}
	for _, v := range m.Emails {
		if v != nil {
			if tempEmails, cErr := v.ToORM(ctx); cErr == nil {
				to.Emails = append(to.Emails, &tempEmails)
			} else {
				return to, cErr
			}
		} else {
			to.Emails = append(to.Emails, nil)
		}
	}
	for _, v := range m.Tasks {
		if v != nil {
			if tempTasks, cErr := v.ToORM(ctx); cErr == nil {
				to.Tasks = append(to.Tasks, &tempTasks)
			} else {
				return to, cErr
			}
		} else {
			to.Tasks = append(to.Tasks, nil)
		}
	}
	if m.BillingAddress != nil {
		tempAddress, err := m.BillingAddress.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.BillingAddress = &tempAddress
	}
	if m.ShippingAddress != nil {
		tempAddress, err := m.ShippingAddress.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.ShippingAddress = &tempAddress
	}
	for _, v := range m.Languages {
		if v != nil {
			if tempLanguages, cErr := v.ToORM(ctx); cErr == nil {
				to.Languages = append(to.Languages, &tempLanguages)
			} else {
				return to, cErr
			}
		} else {
			to.Languages = append(to.Languages, nil)
		}
	}
	for _, v := range m.Friends {
		if v != nil {
			if tempFriends, cErr := v.ToORM(ctx); cErr == nil {
				to.Friends = append(to.Friends, &tempFriends)
			} else {
				return to, cErr
			}
		} else {
			to.Friends = append(to.Friends, nil)
		}
	}
	if posthook, ok := interface{}(m).(UserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserORM) ToPB(ctx context.Context) (User, error) {
	to := User{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if to.CreatedAt, err = ptypes.TimestampProto(m.CreatedAt); err != nil {
		return to, err
	}
	if to.UpdatedAt, err = ptypes.TimestampProto(m.UpdatedAt); err != nil {
		return to, err
	}
	if to.Birthday, err = ptypes.TimestampProto(m.Birthday); err != nil {
		return to, err
	}
	to.Num = m.Num
	if m.CreditCard != nil {
		tempCreditCard, err := m.CreditCard.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.CreditCard = &tempCreditCard
	}
	for _, v := range m.Emails {
		if v != nil {
			if tempEmails, cErr := v.ToPB(ctx); cErr == nil {
				to.Emails = append(to.Emails, &tempEmails)
			} else {
				return to, cErr
			}
		} else {
			to.Emails = append(to.Emails, nil)
		}
	}
	for _, v := range m.Tasks {
		if v != nil {
			if tempTasks, cErr := v.ToPB(ctx); cErr == nil {
				to.Tasks = append(to.Tasks, &tempTasks)
			} else {
				return to, cErr
			}
		} else {
			to.Tasks = append(to.Tasks, nil)
		}
	}
	if m.BillingAddress != nil {
		tempAddress, err := m.BillingAddress.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.BillingAddress = &tempAddress
	}
	if m.ShippingAddress != nil {
		tempAddress, err := m.ShippingAddress.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.ShippingAddress = &tempAddress
	}
	for _, v := range m.Languages {
		if v != nil {
			if tempLanguages, cErr := v.ToPB(ctx); cErr == nil {
				to.Languages = append(to.Languages, &tempLanguages)
			} else {
				return to, cErr
			}
		} else {
			to.Languages = append(to.Languages, nil)
		}
	}
	for _, v := range m.Friends {
		if v != nil {
			if tempFriends, cErr := v.ToPB(ctx); cErr == nil {
				to.Friends = append(to.Friends, &tempFriends)
			} else {
				return to, cErr
			}
		} else {
			to.Friends = append(to.Friends, nil)
		}
	}
	if posthook, ok := interface{}(m).(UserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type User the arg will be the target, the caller the one being converted from

// UserBeforeToORM called before default ToORM code
type UserWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserORM) error
}

// UserAfterToORM called after default ToORM code
type UserWithAfterToORM interface {
	AfterToORM(context.Context, *UserORM) error
}

// UserBeforeToPB called before default ToPB code
type UserWithBeforeToPB interface {
	BeforeToPB(context.Context, *User) error
}

// UserAfterToPB called after default ToPB code
type UserWithAfterToPB interface {
	AfterToPB(context.Context, *User) error
}

type EmailORM struct {
	Email      string
	Id         int32
	Subscribed bool
	UserId     int32
}

// TableName overrides the default tablename generated by GORM
func (EmailORM) TableName() string {
	return "emails"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Email) ToORM(ctx context.Context) (EmailORM, error) {
	to := EmailORM{}
	var err error
	if prehook, ok := interface{}(m).(EmailWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Email = m.Email
	to.Subscribed = m.Subscribed
	if posthook, ok := interface{}(m).(EmailWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *EmailORM) ToPB(ctx context.Context) (Email, error) {
	to := Email{}
	var err error
	if prehook, ok := interface{}(m).(EmailWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Email = m.Email
	to.Subscribed = m.Subscribed
	if posthook, ok := interface{}(m).(EmailWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Email the arg will be the target, the caller the one being converted from

// EmailBeforeToORM called before default ToORM code
type EmailWithBeforeToORM interface {
	BeforeToORM(context.Context, *EmailORM) error
}

// EmailAfterToORM called after default ToORM code
type EmailWithAfterToORM interface {
	AfterToORM(context.Context, *EmailORM) error
}

// EmailBeforeToPB called before default ToPB code
type EmailWithBeforeToPB interface {
	BeforeToPB(context.Context, *Email) error
}

// EmailAfterToPB called after default ToPB code
type EmailWithAfterToPB interface {
	AfterToPB(context.Context, *Email) error
}

type AddressORM struct {
	Address_1 string
	Address_2 string
	Id        int32
	Post      string
}

// TableName overrides the default tablename generated by GORM
func (AddressORM) TableName() string {
	return "addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Address) ToORM(ctx context.Context) (AddressORM, error) {
	to := AddressORM{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Address_1 = m.Address_1
	to.Address_2 = m.Address_2
	to.Post = m.Post
	if posthook, ok := interface{}(m).(AddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AddressORM) ToPB(ctx context.Context) (Address, error) {
	to := Address{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Address_1 = m.Address_1
	to.Address_2 = m.Address_2
	to.Post = m.Post
	if posthook, ok := interface{}(m).(AddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Address the arg will be the target, the caller the one being converted from

// AddressBeforeToORM called before default ToORM code
type AddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *AddressORM) error
}

// AddressAfterToORM called after default ToORM code
type AddressWithAfterToORM interface {
	AfterToORM(context.Context, *AddressORM) error
}

// AddressBeforeToPB called before default ToPB code
type AddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *Address) error
}

// AddressAfterToPB called after default ToPB code
type AddressWithAfterToPB interface {
	AfterToPB(context.Context, *Address) error
}

type LanguageORM struct {
	Code string
	Id   int32
	Name string
}

// TableName overrides the default tablename generated by GORM
func (LanguageORM) TableName() string {
	return "languages"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Language) ToORM(ctx context.Context) (LanguageORM, error) {
	to := LanguageORM{}
	var err error
	if prehook, ok := interface{}(m).(LanguageWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.Code = m.Code
	if posthook, ok := interface{}(m).(LanguageWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *LanguageORM) ToPB(ctx context.Context) (Language, error) {
	to := Language{}
	var err error
	if prehook, ok := interface{}(m).(LanguageWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Name = m.Name
	to.Code = m.Code
	if posthook, ok := interface{}(m).(LanguageWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Language the arg will be the target, the caller the one being converted from

// LanguageBeforeToORM called before default ToORM code
type LanguageWithBeforeToORM interface {
	BeforeToORM(context.Context, *LanguageORM) error
}

// LanguageAfterToORM called after default ToORM code
type LanguageWithAfterToORM interface {
	AfterToORM(context.Context, *LanguageORM) error
}

// LanguageBeforeToPB called before default ToPB code
type LanguageWithBeforeToPB interface {
	BeforeToPB(context.Context, *Language) error
}

// LanguageAfterToPB called after default ToPB code
type LanguageWithAfterToPB interface {
	AfterToPB(context.Context, *Language) error
}

type CreditCardORM struct {
	CreatedAt time.Time
	Id        int32
	Number    string
	UpdatedAt time.Time
	UserId    int32
}

// TableName overrides the default tablename generated by GORM
func (CreditCardORM) TableName() string {
	return "credit_cards"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *CreditCard) ToORM(ctx context.Context) (CreditCardORM, error) {
	to := CreditCardORM{}
	var err error
	if prehook, ok := interface{}(m).(CreditCardWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
	}
	if m.UpdatedAt != nil {
		if to.UpdatedAt, err = ptypes.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
	}
	to.Number = m.Number
	if posthook, ok := interface{}(m).(CreditCardWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *CreditCardORM) ToPB(ctx context.Context) (CreditCard, error) {
	to := CreditCard{}
	var err error
	if prehook, ok := interface{}(m).(CreditCardWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if to.CreatedAt, err = ptypes.TimestampProto(m.CreatedAt); err != nil {
		return to, err
	}
	if to.UpdatedAt, err = ptypes.TimestampProto(m.UpdatedAt); err != nil {
		return to, err
	}
	to.Number = m.Number
	if posthook, ok := interface{}(m).(CreditCardWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type CreditCard the arg will be the target, the caller the one being converted from

// CreditCardBeforeToORM called before default ToORM code
type CreditCardWithBeforeToORM interface {
	BeforeToORM(context.Context, *CreditCardORM) error
}

// CreditCardAfterToORM called after default ToORM code
type CreditCardWithAfterToORM interface {
	AfterToORM(context.Context, *CreditCardORM) error
}

// CreditCardBeforeToPB called before default ToPB code
type CreditCardWithBeforeToPB interface {
	BeforeToPB(context.Context, *CreditCard) error
}

// CreditCardAfterToPB called after default ToPB code
type CreditCardWithAfterToPB interface {
	AfterToPB(context.Context, *CreditCard) error
}

type TaskORM struct {
	Description string
	Name        string
	Priority    int
	UserId      int32
}

// TableName overrides the default tablename generated by GORM
func (TaskORM) TableName() string {
	return "tasks"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Task) ToORM(ctx context.Context) (TaskORM, error) {
	to := TaskORM{}
	var err error
	if prehook, ok := interface{}(m).(TaskWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Name = m.Name
	to.Description = m.Description
	if posthook, ok := interface{}(m).(TaskWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *TaskORM) ToPB(ctx context.Context) (Task, error) {
	to := Task{}
	var err error
	if prehook, ok := interface{}(m).(TaskWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Name = m.Name
	to.Description = m.Description
	if posthook, ok := interface{}(m).(TaskWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Task the arg will be the target, the caller the one being converted from

// TaskBeforeToORM called before default ToORM code
type TaskWithBeforeToORM interface {
	BeforeToORM(context.Context, *TaskORM) error
}

// TaskAfterToORM called after default ToORM code
type TaskWithAfterToORM interface {
	AfterToORM(context.Context, *TaskORM) error
}

// TaskBeforeToPB called before default ToPB code
type TaskWithBeforeToPB interface {
	BeforeToPB(context.Context, *Task) error
}

// TaskAfterToPB called after default ToPB code
type TaskWithAfterToPB interface {
	AfterToPB(context.Context, *Task) error
}

// DefaultCreateUser executes a basic gorm create call
func DefaultCreateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	for i, e := range ormObj.Tasks {
		e.Priority = i
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadUser executes a basic gorm read call
func DefaultReadUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadUser")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := UserORM{}
	db = db.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Order("priority")
	})
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateUser executes a basic gorm update call
func DefaultUpdateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	for i, e := range ormObj.Tasks {
		e.Priority = i
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteUser(ctx context.Context, in *User, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&UserORM{}).Error
	return err
}

// DefaultStrictUpdateUser clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUser(ctx context.Context, in *User, db *gorm.DB) (*User, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	filterCreditCard := CreditCardORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for UserORM")
	}
	filterCreditCard.UserId = ormObj.Id
	if err = db.Where(filterCreditCard).Delete(CreditCardORM{}).Error; err != nil {
		return nil, err
	}
	filterEmails := EmailORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for UserORM")
	}
	filterEmails.UserId = ormObj.Id
	if err = db.Where(filterEmails).Delete(EmailORM{}).Error; err != nil {
		return nil, err
	}
	filterTasks := TaskORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for UserORM")
	}
	filterTasks.UserId = ormObj.Id
	if err = db.Where(filterTasks).Delete(TaskORM{}).Error; err != nil {
		return nil, err
	}
	for i, e := range ormObj.Tasks {
		e.Priority = i
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListUser executes a gorm list call
func DefaultListUser(ctx context.Context, db *gorm.DB) ([]*User, error) {
	ormResponse := []UserORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	db = db.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Order("priority")
	})
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*User{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateEmail executes a basic gorm create call
func DefaultCreateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadEmail executes a basic gorm read call
func DefaultReadEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadEmail")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := EmailORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateEmail executes a basic gorm update call
func DefaultUpdateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteEmail(ctx context.Context, in *Email, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&EmailORM{}).Error
	return err
}

// DefaultStrictUpdateEmail clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListEmail executes a gorm list call
func DefaultListEmail(ctx context.Context, db *gorm.DB) ([]*Email, error) {
	ormResponse := []EmailORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Email{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateAddress executes a basic gorm create call
func DefaultCreateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadAddress executes a basic gorm read call
func DefaultReadAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadAddress")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := AddressORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateAddress executes a basic gorm update call
func DefaultUpdateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteAddress(ctx context.Context, in *Address, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&AddressORM{}).Error
	return err
}

// DefaultStrictUpdateAddress clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListAddress executes a gorm list call
func DefaultListAddress(ctx context.Context, db *gorm.DB) ([]*Address, error) {
	ormResponse := []AddressORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Address{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateLanguage executes a basic gorm create call
func DefaultCreateLanguage(ctx context.Context, in *Language, db *gorm.DB) (*Language, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateLanguage")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadLanguage executes a basic gorm read call
func DefaultReadLanguage(ctx context.Context, in *Language, db *gorm.DB) (*Language, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadLanguage")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := LanguageORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateLanguage executes a basic gorm update call
func DefaultUpdateLanguage(ctx context.Context, in *Language, db *gorm.DB) (*Language, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateLanguage")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteLanguage(ctx context.Context, in *Language, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteLanguage")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&LanguageORM{}).Error
	return err
}

// DefaultStrictUpdateLanguage clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateLanguage(ctx context.Context, in *Language, db *gorm.DB) (*Language, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateLanguage")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListLanguage executes a gorm list call
func DefaultListLanguage(ctx context.Context, db *gorm.DB) ([]*Language, error) {
	ormResponse := []LanguageORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Language{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateCreditCard executes a basic gorm create call
func DefaultCreateCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateCreditCard")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadCreditCard executes a basic gorm read call
func DefaultReadCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadCreditCard")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := CreditCardORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateCreditCard executes a basic gorm update call
func DefaultUpdateCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateCreditCard")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteCreditCard")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&CreditCardORM{}).Error
	return err
}

// DefaultStrictUpdateCreditCard clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateCreditCard(ctx context.Context, in *CreditCard, db *gorm.DB) (*CreditCard, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateCreditCard")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListCreditCard executes a gorm list call
func DefaultListCreditCard(ctx context.Context, db *gorm.DB) ([]*CreditCard, error) {
	ormResponse := []CreditCardORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*CreditCard{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateTask executes a basic gorm create call
func DefaultCreateTask(ctx context.Context, in *Task, db *gorm.DB) (*Task, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateTask")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultListTask executes a gorm list call
func DefaultListTask(ctx context.Context, db *gorm.DB) ([]*Task, error) {
	ormResponse := []TaskORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Task{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

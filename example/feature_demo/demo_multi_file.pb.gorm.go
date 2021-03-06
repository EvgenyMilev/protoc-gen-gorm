// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/feature_demo/demo_multi_file.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example/feature_demo/demo_multi_file.proto
	example/feature_demo/demo_types.proto
	example/feature_demo/demo_service.proto

It has these top-level messages:
	ExternalChild
	TestTypes
	TypeWithID
	MultiaccountTypeWithID
	MultiaccountTypeWithoutID
	APIOnlyType
	PrimaryUUIDType
	PrimaryStringType
	PrimaryIncluded
	IntPoint
	CreateIntPointRequest
	CreateIntPointResponse
	ReadIntPointRequest
	ReadIntPointResponse
	UpdateIntPointRequest
	UpdateIntPointResponse
	DeleteIntPointRequest
	DeleteIntPointResponse
	ListIntPointResponse
	Something
	ListIntPointRequest
*/
package example

import context "context"
import errors "errors"

import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gateway1 "github.com/infobloxopen/atlas-app-toolkit/gateway"
import go_uuid1 "github.com/satori/go.uuid"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import query1 "github.com/infobloxopen/atlas-app-toolkit/query"

import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type ExternalChildORM struct {
	Id                  string
	PrimaryIncludedId   *go_uuid1.UUID
	PrimaryStringTypeId *string
	PrimaryUUIDTypeId   *go_uuid1.UUID
}

// TableName overrides the default tablename generated by GORM
func (ExternalChildORM) TableName() string {
	return "external_children"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *ExternalChild) ToORM(ctx context.Context) (ExternalChildORM, error) {
	to := ExternalChildORM{}
	var err error
	if prehook, ok := interface{}(m).(ExternalChildWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if posthook, ok := interface{}(m).(ExternalChildWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ExternalChildORM) ToPB(ctx context.Context) (ExternalChild, error) {
	to := ExternalChild{}
	var err error
	if prehook, ok := interface{}(m).(ExternalChildWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if posthook, ok := interface{}(m).(ExternalChildWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type ExternalChild the arg will be the target, the caller the one being converted from

// ExternalChildBeforeToORM called before default ToORM code
type ExternalChildWithBeforeToORM interface {
	BeforeToORM(context.Context, *ExternalChildORM) error
}

// ExternalChildAfterToORM called after default ToORM code
type ExternalChildWithAfterToORM interface {
	AfterToORM(context.Context, *ExternalChildORM) error
}

// ExternalChildBeforeToPB called before default ToPB code
type ExternalChildWithBeforeToPB interface {
	BeforeToPB(context.Context, *ExternalChild) error
}

// ExternalChildAfterToPB called after default ToPB code
type ExternalChildWithAfterToPB interface {
	AfterToPB(context.Context, *ExternalChild) error
}

// DefaultCreateExternalChild executes a basic gorm create call
func DefaultCreateExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateExternalChild")
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

// DefaultReadExternalChild executes a basic gorm read call
func DefaultReadExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadExternalChild")
	}
	db = db.Set("gorm:auto_preload", true)
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormParams.Id == "" {
		return nil, errors.New("Read requires a non-zero primary key")
	}
	ormResponse := ExternalChildORM{}
	if err = db.Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateExternalChild executes a basic gorm update call
func DefaultUpdateExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateExternalChild")
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

func DefaultDeleteExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteExternalChild")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&ExternalChildORM{}).Error
	return err
}

// DefaultStrictUpdateExternalChild clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateExternalChild")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	count := 1
	err = db.Model(&ormObj).Where("id=?", ormObj.Id).Count(&count).Error
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
	if count == 0 {
		err = gateway1.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

// DefaultPatchExternalChild executes a basic gorm update call with patch behavior
func DefaultPatchExternalChild(ctx context.Context, in *ExternalChild, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultPatchExternalChild")
	}
	pbReadRes, err := DefaultReadExternalChild(ctx, &ExternalChild{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj := *pbReadRes
	if _, err := DefaultApplyFieldMaskExternalChild(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ExternalChildWithBeforePatchSave); ok {
		if ctx, db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return DefaultStrictUpdateExternalChild(ctx, &pbObj, db)
}

type ExternalChildWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *ExternalChild, *field_mask1.FieldMask, *gorm1.DB) (context.Context, *gorm1.DB, error)
}

// DefaultApplyFieldMaskExternalChild patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskExternalChild(ctx context.Context, patchee *ExternalChild, patcher *ExternalChild, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*ExternalChild, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.New("Patchee inputs to DefaultApplyFieldMaskExternalChild must be non-nil")
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// getCollectionOperators takes collection operator values from corresponding message fields
func getCollectionOperators(in interface{}) (*query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection, error) {
	f := &query1.Filtering{}
	err := gateway1.GetCollectionOp(in, f)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	s := &query1.Sorting{}
	err = gateway1.GetCollectionOp(in, s)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	p := &query1.Pagination{}
	err = gateway1.GetCollectionOp(in, p)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	fs := &query1.FieldSelection{}
	err = gateway1.GetCollectionOp(in, fs)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return f, s, p, fs, nil
}

// DefaultListExternalChild executes a gorm list call
func DefaultListExternalChild(ctx context.Context, db *gorm1.DB, req interface{}) ([]*ExternalChild, error) {
	ormResponse := []ExternalChildORM{}
	f, s, p, fs, err := getCollectionOperators(req)
	if err != nil {
		return nil, err
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &ExternalChildORM{}, &ExternalChild{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	if fs.GetFields() == nil {
		db = db.Set("gorm:auto_preload", true)
	}
	in := ExternalChild{}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	db = db.Where(&ormParams)
	db = db.Order("id")
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*ExternalChild{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

package github

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"reflect"
)

func (D *DBGORM) Insert(input interface{}) error {
	id := D.db.Create(input)
	return id.Error
}

func (D *DBGORM) GetAll(dest interface{}) error {
	result := D.db.Find(dest)

	sqlString := D.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return D.db.Find(dest)
	})
	log.Print(sqlString)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Ошибка 0 полей ")
	}
	return nil
}

func (D *DBGORM) GetById(dest interface{}, id interface{}) error {
	result := D.db.Where("ID = ?", id).Find(dest)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Ошибка 0 полей ")
	}
	return nil
}

func (D *DBGORM) Update(dest interface{}, id interface{}) error {
	result := D.db.Model(dest).Where("ID = ?", id).Updates(dest)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Ошибка 0 полей ")
	}
	return nil

}

func (D *DBGORM) Delete(dest interface{}, id interface{}, softDelete bool) error {
	if softDelete {
		result := D.db.Delete(dest, id)
		return result.Error
	}

	result := D.db.Unscoped().Delete(dest, id)
	return result.Error
}

func (D *DBGORM) GetPaginatedResult(db *gorm.DB, query *gorm.DB, page, limit int) (*PaginatedResult, error) {
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit
	var list []interface{}
	if err := query.Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}

	result := &PaginatedResult{
		total: int(count),
		limit: limit,
		page:  page,
		list:  list,
	}

	return result, nil
}

func (D *DBGORM) GetPaginatedResultFromSlice(data interface{}, page, limit int) (*PaginatedResult, error) {
	dataSlice := reflect.ValueOf(data)
	if dataSlice.Kind() != reflect.Slice {
		return nil, errors.New("data is not a slice")
	}

	if page <= 0 || limit <= 0 {
		return nil, errors.New("invalid page or limit")
	}
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= dataSlice.Len() {
		return &PaginatedResult{
			total: dataSlice.Len(),
			limit: limit,
			page:  page,
			list:  nil,
		}, nil
	}

	if endIndex > dataSlice.Len() {
		endIndex = dataSlice.Len()
	}

	paginatedData := make([]interface{}, endIndex-startIndex)

	return &PaginatedResult{
		total: dataSlice.Len(),
		limit: limit,
		page:  page,
		list:  paginatedData,
	}, nil
}

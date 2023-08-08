package SebbiaDB

import (
	"errors"
	"gorm.io/gorm"
	"log"
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

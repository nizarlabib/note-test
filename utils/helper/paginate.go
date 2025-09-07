package helper

import (
	"math"

	"gorm.io/gorm"
)

func PaginateV2(model any, pg *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	_ = db.Model(model).Count(&totalRows).Error

	pg.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(pg.GetLimit())))
	pg.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(pg.GetOffset()).
			Limit(pg.GetLimit())
	}
}

func Paginate(model interface{}, pagination *Pagination, db *gorm.DB, args ...interface{}) func(db *gorm.DB) *gorm.DB {
	var query interface{}
	var rest []interface{}

	if len(args) > 0 {
		query = args[0]
	}

	if len(args) > 1 {
		rest = args[1:]
	}

	var totalRows int64

	if len(args) > 0 {
		db.Model(model).Where(query, rest...).Count(&totalRows)
	} else {
		db.Model(model).Count(&totalRows)
	}

	pagination.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
		// .Order("id asc")
	}
}

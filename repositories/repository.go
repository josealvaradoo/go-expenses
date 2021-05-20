package repositories

import (
	"github.com/josealvaradoo/expenses/database"
)

type Repository interface {
	List(model interface{})
	Get(model interface{}, id int)
	Create(model interface{})
	Update(model interface{})
	Delete(model interface{}, id int)
}

func List(model interface{}) {
	db := database.DBConnection
	db.Find(model)
}

func Get(model interface{}, id int) {
	db := database.DBConnection
	db.Find(model, id)
}

func Create(model interface{}) {
	db := database.DBConnection
	db.Create(model)
}

func Update(model interface{}) {
	db := database.DBConnection
	db.Updates(model)
}

func Delete(model interface{}, id int) {
	db := database.DBConnection
	db.Delete(model, id)
}

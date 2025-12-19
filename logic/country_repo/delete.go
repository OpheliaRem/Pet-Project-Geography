package country_repo

import (
	"backPet0/db"
	"backPet0/logger"
	"strconv"
)

func Remove(id int) error {
	query := "delete from countries where id=" + strconv.Itoa(id)
	logger.Log(query)

	_, err := db.Db.Exec(query)

	return err
}

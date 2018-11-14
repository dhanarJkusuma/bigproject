package user

import (
	"bigproject/entity"
	"bigproject/util/nonpanic"
	"bigproject/util/parser"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

type UserModule struct {
	DB *sql.DB
}

func (um *UserModule) FetchData(page, size int) []entity.User {
	offset := page * size
	const querySelect = `SELECT user_id, full_name, msisdn, user_email, COALESCE (birth_date, NOW()), AGE(NOW(), COALESCE (birth_date, NOW())) as age FROM ws_user ORDER BY full_name ASC LIMIT $1 OFFSET $2`

	result, err := um.DB.Query(querySelect, size, offset)
	if err != nil {
		log.Printf("[BigProject] : Error %v\n", err.Error())
	}

	data := []entity.User{}
	for result.Next() {
		var id int
		var fullname, msisdn, email, birthDate, age string
		var date time.Time

		err = result.Scan(&id, &fullname, &msisdn, &email, &birthDate, &age)
		if err != nil {
			log.Printf("[BigProject] : Error %v\n", err.Error())
		}
		user := entity.User{}
		user.UserID = id
		user.FullName = fullname
		user.Msisdn = msisdn
		user.UserEmail = email
		user.BirthDate = birthDate
		user.UserAge = age
		if user.BirthDate != "" {
			date, err = parser.ParseTimeFromDB(user.BirthDate)
			nonpanic.HandleParsingDBValueError(err)

			now := parser.FormatDate(time.Now())
			user.BirthDate = parser.FormatDate(date)
			if strings.Compare(now, user.BirthDate) == 0 {
				user.BirthDate = " - "
				user.UserAge = " - "
			}else{
				idx := strings.Index(user.UserAge, "days") + 5
				user.UserAge = user.UserAge[:idx]
			}
		}

		data = append(data, user)
	}
	return data
}

func (um *UserModule) FetchCount() int {
	const querySelect = `SELECT count(*) FROM ws_user`

	result, err := um.DB.Query(querySelect)
	if err != nil {
		log.Printf("[BigProject] : Error %v\n", err.Error())
	}

	var count int
	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			log.Printf("[BigProject] : Error %v\n", err.Error())
		}
	}
	return count
}

func (um *UserModule) FetchSearch(page, size int, keyword string)[]entity.User {
	offset := page * size
	const querySelect = `SELECT user_id, full_name, msisdn, user_email, COALESCE(birth_date, NOW()), AGE(NOW(), COALESCE (birth_date, NOW())) as age FROM ws_user WHERE lower(full_name) LIKE $1 ORDER BY full_name ASC LIMIT $2 OFFSET $3`

	result, err := um.DB.Query(querySelect, fmt.Sprintf("%%%s%%", strings.ToLower(keyword)), size, offset)
	if err != nil {
		log.Printf("[BigProject] : Error %v\n", err.Error())
	}

	data := []entity.User{}
	for result.Next() {
		var id int
		var fullname, msisdn, email, birthDate, age string
		var date time.Time

		err = result.Scan(&id, &fullname, &msisdn, &email, &birthDate, &age)
		if err != nil {
			log.Printf("[BigProject] : Error %v\n", err.Error())
		}
		user := entity.User{}
		user.UserID = id
		user.FullName = fullname
		user.Msisdn = msisdn
		user.UserEmail = email
		user.BirthDate = birthDate
		user.UserAge = age
		if user.BirthDate != "" {
			date, err = parser.ParseTimeFromDB(user.BirthDate)
			nonpanic.HandleParsingDBValueError(err)

			now := parser.FormatDate(time.Now())
			user.BirthDate = parser.FormatDate(date)
			if strings.Compare(now, user.BirthDate) == 0 {
				user.BirthDate = " - "
				user.UserAge = " - "
			} else {
				idx := strings.Index(user.UserAge, "days") + 5
				user.UserAge = user.UserAge[:idx]
			}
		}

		data = append(data, user)
	}
	return data
}

func (um *UserModule) FetchCountSearch(keyword string) int {
	const querySelect = `SELECT count(*) FROM ws_user WHERE lower(full_name) LIKE $1`

	result, err := um.DB.Query(querySelect, fmt.Sprintf("%%%s%%", strings.ToLower(keyword)))
	if err != nil {
		log.Printf("[BigProject] : Error %v\n", err.Error())
	}

	var count int
	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			log.Printf("[BigProject] : Error %v\n", err.Error())
		}
	}
	return count
}


func RegisterUserModule(DB *sql.DB) *UserModule {
	return &UserModule{
		DB: DB,
	}
}
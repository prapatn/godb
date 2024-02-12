package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

type MatchLog struct {
	ID          uint      `db:"id"`
	MatchNumber string    `db:"match_number"`
	Player      string    `db:"player"`
	Turn        int       `db:"turn"`
	BallPower   int       `db:"ball_power"`
	Time        time.Time `db:"time"`
}

func main() {
	db, err = sqlx.Open("mysql", "root:example@tcp(localhost)/local?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("connect database")

	// err = Get()

	// err = GetByID(3731)

	// err = Insert(
	// 	MatchLog{
	// 		MatchNumber: "test",
	// 		Player:      "1",
	// 		Turn:        1,
	// 		BallPower:   100,
	// 	},
	// )

	// err = Update(
	// 	MatchLog{
	// 		ID:     3743,
	// 		Player: "w",
	// 	},
	// )

	// err = Delete(3743)

	// GetX()
	err = GetByIDX(3731)

	if err != nil {
		panic(err)
	}

}

func GetX() error {
	query := "SELECT id,match_number,player,turn,ball_power,time FROM match_logs ORDER BY `id` DESC LIMIT 10 "
	logs := []MatchLog{}
	err = db.Select(&logs, query)
	if err != nil {
		return err
	}
	fmt.Println(logs)
	return nil
}

func Get() error {
	query := "SELECT id,match_number,player,turn,ball_power,time FROM match_logs ORDER BY `id` DESC LIMIT 100 "
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var logs MatchLog

	for rows.Next() {
		err = rows.Scan(&logs.ID, &logs.MatchNumber, &logs.Player, &logs.Turn, &logs.BallPower, &logs.Time)
		if err != nil {
			return err
		}
		log.Println(logs)
	}

	return nil
}

func GetByIDX(id int) error {
	var logs MatchLog
	query := "SELECT id,match_number,player,turn,ball_power,time FROM match_logs WHERE id = ?"
	err := db.Get(&logs, query, id)
	if err != nil {
		return err
	}

	log.Println(logs)

	return nil
}

func GetByID(id int) error {
	query := "SELECT id,match_number,player,turn,ball_power,time FROM match_logs WHERE id = ?"
	rows := db.QueryRow(query, id)

	var logs MatchLog

	err = rows.Scan(&logs.ID, &logs.MatchNumber, &logs.Player, &logs.Turn, &logs.BallPower, &logs.Time)
	if err != nil {
		return err
	}
	log.Println(logs)

	return nil
}

func Insert(log MatchLog) error {
	query := "INSERT INTO match_logs (match_number,player,turn,ball_power) VALUES (?,?,?,?)"

	result, err := db.Exec(query, log.MatchNumber, log.Player, log.Turn, log.BallPower)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected <= 0 {
		return errors.New("Inset fail")
	}

	return nil
}

func Update(log MatchLog) error {
	query := "UPDATE match_logs SET player = ? WHERE id = ?"

	result, err := db.Exec(query, log.Player, log.ID)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected <= 0 {
		return errors.New("Update fail")
	}

	return nil
}

func Delete(id int) error {
	query := "DELETE FROM match_logs  WHERE id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected <= 0 {
		return errors.New("Delete fail")
	}

	return nil
}

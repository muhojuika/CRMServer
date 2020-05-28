package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"time"
)
type logBook struct {
	id string
	desc string
}
const dbName ="crmDB.db"

func LoginUser(login string, password string) (User, error) { // load user from DB
	var user User
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	row := db.QueryRow("select * from Users where login=$1 and password=$2", login, password)
	err = row.Scan(&user.id, &user.login, &user.password, &user.fio, &user.perms, &user.email, &user.phone)
	if err != nil{
		return user, err
	}
	return user, nil
}
func GetUser(userId int) (User, error) { // load user from DB
	var user User
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	row := db.QueryRow("select * from Users where id =$1", userId)
	err = row.Scan(&user.id, &user.login, &user.password, &user.fio, &user.perms, &user.email, &user.phone)
	if err != nil{
		return user, err
	}
	return user, nil
}
func RegisterUser(user User)error{ //register new user
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	_, err = db.Exec("insert into users(login, password, fio, perms, email, phone)  values  ($1,$2,$3,$4,$5,$6)",
		user.login, user.password, user.fio, user.perms, user.email, user.phone)
	if err!=nil{
		return err
	}
	return nil
}
func RegisterLog(event_type int, user User)error{ //
	var eventName string
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	switch event_type {
	case 0:
		eventName = "Auth user: "+user.login+"  Time: "+ time.Now().Format(time.RFC822)
		break
	case 1:
		eventName = "Create user: "+user.login+"  Time: "+ time.Now().Format(time.RFC822)
		break
	case 2:
		eventName = "Clear Table: "+user.login+"  Time: "+ time.Now().Format(time.RFC822)
		break

	}
	_, err = db.Exec("insert into logbook(event)  values  ($1)",
		eventName)
	if err!=nil{
		return err
	}
	return nil
}
func GetLog() (string, error) { // load logbook
	logBooks := []logBook{}
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	rows, err := db.Query("select * from logbook")
	if err != nil{
		return "", err
	}
	lg:= logBook{}
	for rows.Next(){
		err := rows.Scan(&lg.id, &lg.desc)
		if err != nil {
			log.Println(err)
			return "", err
		}
		logBooks = append(logBooks, lg)
	}
	var msg string
	msg+="log:"
	for _, l := range logBooks{
		msg+= l.id+"-"+l.desc+";"
	}
	return msg, nil
}
func ClearLog(){
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	result, err := db.Exec("delete from logbook")
	if err != nil{
		panic(err)
	}
	log.Println(result.RowsAffected())
}
func GetUserList()(string, error){
	msg := "userlist:{"
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	rows, err := db.Query("select * from Users")
	if err != nil{
		return "", err
	}
	for rows.Next(){
		var user User
		err := rows.Scan(&user.id, &user.login, &user.password, &user.fio, &user.perms, &user.email, &user.phone)
		if err != nil {
			log.Println(err)
			return "", err
		}
		msg+=strconv.Itoa(user.id)+":"+user.login+":"+user.fio+":"+strconv.Itoa(user.perms)+":"+user.phone+":"+user.email+";"
	}
	msg+="}"

	return msg, nil

}
func RegisterTaskLog(t Task)error{ //
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}

	_, err = db.Exec("insert into tasks(name, description, time_create, time_work, userF_id, userT_id, prir, status)  values  ($1,$2,$3,$4,$5,$6,$7,$8)",
		t.name, t.desc, t.time_cr, t.time_work, t.userF, t.userT, t.pr, t.status)
	if err!=nil{
		return err
	}
	return nil
}
func GetTasksList()(string, error){
	msg := "taskslist:{"
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	rows, err := db.Query("select * from tasks")
	if err != nil{
		return "", err
	}
	for rows.Next(){
		var ts Task
		err := rows.Scan(&ts.id, &ts.name, &ts.desc, &ts.time_cr, &ts.time_work, &ts.userT, &ts.userF, &ts.pr, &ts.status)
		if err != nil {
			log.Println(err)
			return "", err
		}
		msg+=ts.id+":"+ts.name+":"+ts.desc+":"+ts.time_cr+":"+ts.time_work+":"+ts.userT+":"+ts.userF+":"+ts.pr+":"+ts.status+";"
	}
	msg+="}"

	return msg, nil
}
func GetTasksListForUser(user_id string)(string, error){
	msg := "taskslist:{"
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	rows, err := db.Query("select * from tasks WHERE userT_id=$1",user_id)
	if err != nil{
		return "", err
	}
	for rows.Next(){
		var ts Task
		err := rows.Scan(&ts.id, &ts.name, &ts.desc, &ts.time_cr, &ts.time_work, &ts.userT, &ts.userF, &ts.pr, &ts.status)
		if err != nil {
			log.Println(err)
			return "", err
		}
		msg+=ts.id+":"+ts.name+":"+ts.desc+":"+ts.time_cr+":"+ts.time_work+":"+ts.userT+":"+ts.userF+":"+ts.pr+":"+ts.status+";"
	}
	msg+="}"

	return msg, nil
}
func CompeteTask(id string)(string, error){
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	_, err = db.Exec("update tasks set status=1 where id=$1", id)
	if err != nil{
		return "error", err
	}
	return "complete", nil
}
func DeleteTask(id string)(string, error){
	db, err := sql.Open("sqlite3",dbName)
	if err != nil{
		panic(err)
	}
	_, err = db.Exec("delete from tasks where id=$1", id)
	if err != nil{
		return "error", err
	}
	return "complete", nil
}
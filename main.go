package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
type User struct {
	id int
	login string
	password string
	fio string
	perms int
	email string
	phone string
}
type Task struct {
	id string
	name string
	desc string
	time_cr string
	time_work string
	userT string
	userF string
	pr string
	status string
}

const Server_adress string = "localhost:9000"//82.146.63.120

func GetLogHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	log.Println("path", r.URL.Path, r.RemoteAddr)
	msg, err := GetLog()
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
		return
	}
	if msg!=""{
		fmt.Fprintf(w, msg)
	}
}
func ClearLogHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	userId, err := strconv.Atoi(r.Form.Get("user_id"))
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
		return
	}
	log.Println("path", r.URL.Path, r.RemoteAddr)
	ClearLog()
	user,err :=GetUser(userId)
	err = RegisterLog(2,user)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	fmt.Fprintf(w, "Успешно!")
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path", r.URL.Path, r.RemoteAddr)
	login:= r.Form.Get("login")
	password:= r.Form.Get("password")

	us, err := LoginUser(login, password)
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
		return
	}
	err = RegisterLog(0,us)
	if err!=nil{
		log.Println(err.Error())
	}
	fmt.Fprintf(w, CompilingMsgUser(us))
}
func RegisterUserHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	log.Println("path", r.URL.Path, r.RemoteAddr, r.URL.RawQuery)
	var us User
	us.login = r.Form.Get("login")
	us.password = r.Form.Get("password")
	us.fio = r.Form.Get("fio")
	us.email = r.Form.Get("email")
	us.phone = r.Form.Get("phone")
	us.perms,_ = strconv.Atoi(r.Form.Get("perms"))
	log.Println(us)
	err:=RegisterUser(us)
	if err != nil{
		log.Println(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Пользователь зарегистрирован!")
}
func TaskCompleteHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	log.Println("path", r.URL.Path, r.RemoteAddr)
	id:= r.Form.Get("id")
	_, err:= CompeteTask(id)
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
	}
	fmt.Fprintf(w, "Complete")
}
func GetUserListHandler(w http.ResponseWriter, r *http.Request){
	log.Println("path", r.URL.Path, r.RemoteAddr)
	msg, err:=GetUserList()
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
	}
	if msg!=""{
		fmt.Fprintf(w, msg)
	}
}
func CreateTaskHandler(w http.ResponseWriter, r *http.Request){
	var t Task
	r.ParseForm()
	log.Println("path", r.URL.Path, r.RemoteAddr)
	t.name = r.Form.Get("name")
	t.desc = r.Form.Get("des")
	t.time_cr = r.Form.Get("timeC")
	t.time_work = r.Form.Get("timeW")
	t.userF = r.Form.Get("userF")
	t.userT = r.Form.Get("userT")
	t.pr = r.Form.Get("pr")
	t.status = r.Form.Get("status")
	err:= RegisterTask(t)
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
	}
	i,_:= strconv.Atoi(t.userF)
	us, err:= GetUser(i)
	err = RegisterLog(3,us)
	if err!=nil{
		log.Println(err.Error())
	}
	fmt.Fprintf(w, "Задача успешно создана!")
}
func GetTaskListHandler(w http.ResponseWriter, r *http.Request){
	log.Println("path", r.URL.Path, r.RemoteAddr)
	msg, err:=GetTasksList()
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
	}
	if msg!=""{
		fmt.Fprintf(w, msg)
	}
}
func GetTaskListForUserHandler(w http.ResponseWriter, r *http.Request){
	log.Println("path", r.URL.Path, r.RemoteAddr)
	r.ParseForm()
	id:= r.Form.Get("user_id")
	msg, err:=GetTasksListForUser(id)
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
	}
	if msg!=""{
		fmt.Fprintf(w, msg)
	}
}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request){
	log.Println("path", r.URL.Path, r.RemoteAddr)
	r.ParseForm()
	id:= r.Form.Get("task_id")
	msg, err:=DeleteTask(id)
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
	}
	if msg!=""{
		fmt.Fprintf(w, "Задача успешно удалена!")
	}
}
func EditTaskHandler(w http.ResponseWriter, r *http.Request){
	var t Task
	r.ParseForm()
	log.Println("path", r.URL.Path, r.RemoteAddr)
	t.id = r.Form.Get("id")
	t.name = r.Form.Get("name")
	t.desc = r.Form.Get("des")
	t.time_cr = r.Form.Get("timeC")
	t.time_work = r.Form.Get("timeW")
	t.userF = r.Form.Get("userF")
	t.userT = r.Form.Get("userT")
	t.pr = r.Form.Get("pr")
	t.status = r.Form.Get("status")
	err:= EditTask(t)
	if err!=nil{
		log.Println(err.Error())
		fmt.Fprintf(w, "error:Ошибка данных -("+err.Error()+")")
	}
	//i,_:= strconv.Atoi(t.userF)
	//us, err:= GetUser(i)
	//err = RegisterLog(3,us)
	//if err!=nil{
	//	log.Println(err.Error())
	//}
	fmt.Fprintf(w, "Задача успешно изменена")
}

func main()  {
	go http.HandleFunc("/login", LoginHandler)
	go http.HandleFunc("/register", RegisterUserHandler)
	go http.HandleFunc("/log", GetLogHandler)
	go http.HandleFunc("/clear_log", ClearLogHandler)
	go http.HandleFunc("/user_list", GetUserListHandler)
	go http.HandleFunc("/create_task", CreateTaskHandler)
	go http.HandleFunc("/task_list", GetTaskListHandler)
	go http.HandleFunc("/tasks_forUS", GetTaskListForUserHandler)
	go http.HandleFunc("/complete_task", TaskCompleteHandler)
	go http.HandleFunc("/delete_task", DeleteTaskHandler)
	go http.HandleFunc("/edit_task", EditTaskHandler)
	log.Println("Server starting on",Server_adress)
	err := http.ListenAndServe(Server_adress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

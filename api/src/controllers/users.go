package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("criando usuario!"))
}
func GetUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("buscando usuario!"))
}
func GetUserById(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("buscando usuarios!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("atualizando usuario!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("deletando usuario!"))
}
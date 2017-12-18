package main

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/TestTask/model"
	"github.com/gorilla/mux"
	"log"
    "fmt"
	"math"
)

var (
	Auth map[string]string
	Work map[string]int32
)

func main() {
	r := mux.NewRouter()
    model.GormInit()
    defer model.GormClose()

	r.HandleFunc("/", mainPage)
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/login/pass", changePass).Methods("PUT")
	r.HandleFunc("/work", doWork).Methods("POST")

	//http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))

}

func init() {
    Auth = make(map[string]string)
	Work = make(map[string]int32, 0)
	Work["admin"] = 1000000
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta name="theme-color" content="#375EAB">
		
			<title>main page</title>
		</head>
		<body>
			Page body and some more content
		</body>
		</html>`))
}

func login(w http.ResponseWriter, r *http.Request) {
    //http://localhost:8000/login?login=test_user&pass=test_pass
	login := r.FormValue("login")
	pass := r.FormValue("pass")


    if login=="" && pass == "" {
		w.WriteHeader(http.StatusBadRequest)
        return
    }

    if Auth[login] == pass{
		w.WriteHeader(http.StatusOK)
		return
	}

	user := &model.User{}
	err := user.Get(login, pass)
	if err == nil {
		Auth[login] = pass
		Work[login] = user.WorkNumber
        w.WriteHeader(http.StatusOK)

	}else {
        w.WriteHeader(http.StatusBadRequest)

    }
}

func changePass(w http.ResponseWriter, r *http.Request) {
    //http://localhost:8000/login/pass?login=test_user&pass=test_pass&newPass=1
	login := r.FormValue("login")
	pass := r.FormValue("pass")

	newPass := r.FormValue("newPass")

	if Auth[login] != pass {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &model.User{}
	err := user.Get(login,pass)
	if err == nil {
        user.Pass = newPass
        err = user.Save()
    }
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}else {
		Auth[login] = newPass
        w.WriteHeader(http.StatusOK)
    }
}

type DTO struct {
	BigNumber int64  `json:"bigNumber"`
	Text      string `json:"text"`
}

func doWork(w http.ResponseWriter, r *http.Request) {
    //http://localhost:8000//work?login=test_user&value={"bigNumber":123123,"text":"string"}

	var value DTO
	login := r.FormValue("login")
	if _,ok:= Work[login]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

    st:= r.FormValue("value")
	json.Unmarshal([]byte(st), &value)

	v := reflect.ValueOf(value)
	for i := 0; i < v.NumField(); i++ {
		w.Write(append(reverse(v.Field(i)),[]byte("\n")...))
	}
}

func reverse(val reflect.Value) []byte {

	switch val.Kind() {
	case reflect.Int64:
		//result := make([]byte, 8)
		//binary.LittleEndian.PutUint64(result, uint64(math.MaxInt64-val.Interface().(int64)))
		return []byte(fmt.Sprint(math.MaxInt64-val.Interface().(int64)))

	case reflect.Int32:
		//result := make([]byte, 4)
		//binary.LittleEndian.PutUint32(result, uint32(math.MaxInt32-val.Interface().(int32)))
		//return result
		return []byte(fmt.Sprint(math.MaxInt32-val.Interface().(int32)))

	case reflect.String:
		var result string

		for i := len(val.Interface().(string))-1; i >= 0; i-- {
			result += string(val.Interface().(string)[i])
		}
		return []byte(result)
	}
	return nil
}

package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"io"
)

type user_struct struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)
	http.ListenAndServe(":8080", nil)
}



/* /user/createにおける処理を行う関数 */
func createUser(w http.ResponseWriter, req *http.Request) {
	// リクエストがPOSTか確認
	if req.Method != "POST" { //POST以外だったら...
		w.WriteHeader(http.StatusBadRequest) //400:bad requestにする
		return
	}

	// リクエストのコンテンツタイプがjson確認
	if req.Header.Get("Content-Type") != "application/json" { //json以外だったら
		w.WriteHeader(http.StatusBadRequest) //400:bad requestにする
		return
	}

	// コンテンツの長さをstringからintに変換してlengthに格納する
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// bodyという配列を作って、リクエストされたjsonのデータを格納する
	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// jsonBodyに、bodyに格納されているjsonデータをパースして格納する
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, jsonBody["name"].(string))

	w.WriteHeader(http.StatusOK)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to user/get, from Docker container!")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to user/update, from Docker container!")
}

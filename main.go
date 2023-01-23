package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
)

type post struct {
	iban          string `json:"iban"`
	bankname      string `json:"bankname"`
	routingnumber string `json:"routingnumber"`
}

func main() {

	resp, err := http.Get("https://random-data-api.com/api/v2/banks")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	/*sb := string(body)
	log.Println(sb)*/

	/*post := post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("iban no = %s", post.iban)
	log.Println("bankname = %s", post.bankname)
	log.Println("routing number = %s", post.routingnumber) */

	var responseObject post
	json.Unmarshal(body, &responseObject)

	fmt.Println(responseObject.iban)
	fmt.Println(responseObject.bankname)
	fmt.Println(responseObject.routingnumber)
}

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://feedbacks-api.wildberries.ru/api/v1/questions?isAnswered=true&take=100&skip=0", nil)
	req.Header.Set("Authorization", "eyJhbGciOiJFUzI1NiIsImtpZCI6IjIwMjQwNzE1djEiLCJ0eXAiOiJKV1QifQ.eyJlbnQiOjEsImV4cCI6MTczNjk5NTIzMCwiaWQiOiJiYmI4YTQyOS0yYjRiLTQ4ZTUtOTRlNi02YzFmNDBkZGE5ZDAiLCJpaWQiOjUzMjk4ODk0LCJvaWQiOjExNzI0NzQsInMiOjEyOCwic2lkIjoiMzFhMzNhZjMtN2JkNy00MzQ1LWFmYmItOWJlZDgwYTJiMzE1IiwidCI6ZmFsc2UsInVpZCI6NTMyOTg4OTR9.tn9yaSThziQ-AxluPhy-hGZhrlpnINCCQoaf2tKV3I-_FcRobHb2V0dZrBe0sEepHBbxSgSl2fTIEfSLa7oTew")
	res, _ := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

package main

import (
	"fmt"
	//"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	//"encoding/json"
	//"strings"
	"io/ioutil"
)


func StartRTM(token string) (error) {
	resp, err := http.PostForm("https://slack.com/api/rtm.start", url.Values{"token": {token}})
	if err != nil {
		return fmt.Errorf("no dice with rtm.start: %v", err)
	}
	defer resp.Body.Close()


	respout,_:=ioutil.ReadAll(resp.Body)
	fmt.Printf("%s",string(respout))

/*
	authResp:=new(AuthResponse)
	dec:=json.NewDecoder(resp.Body)
	err=dec.Decode(authResp)
	if err != nil {
		return fmt.Errorf("Couldn't decode json. ERR: %v", err)
	}
	splitUrl := strings.Split(authResp.Url, "/")
	splitUrl[2] = splitUrl[2] + ":443"
	authResp.Url = strings.Join(splitUrl, "/")

	fmt.Println(authResp)
*/
	return nil
}

func main(){
	
	err:=StartRTM(`xoxp-3057259082-3057259092-3247338018-9890ee`)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(`ok`)
	}
}

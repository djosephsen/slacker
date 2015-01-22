package handlers

import (
	sl "github.com/djosephsen/slacker/slackerlib"
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
)

type gifyout struct{
	Meta	interface{}
	Data	struct{
		Tags []string
		Caption string
		Username string
	 	Image_width string
    	Image_frames string
    	Image_mp4_url string
    	Image_url string
    	Image_original_url string
    	Url string
    	Id string
    	Type string
    	Image_height string
    	Fixed_height_downsampled_url string
    	Fixed_height_downsampled_width string
    	Fixed_height_downsampled_height string
    	Fixed_width_downsampled_url string
    	Fixed_width_downsampled_width string
    	Fixed_width_downsampled_height string
    	Rating string
	}
}


var Gifme = sl.MessageHandler{
	Name: `Gifme`, 
	Usage: `botname gif me freddie mercury: returns a random rated:PG gif of freddy mercury via the giphy API`,
	Method:  `RESPOND`,
	Pattern: `(?i)gif me (.*)`,
	Run: func(e *sl.Event, match []string){
	
		search:=match[1]
		q:=url.QueryEscape(search)
		myurl:=fmt.Sprintf("http://api.giphy.com/v1/gifs/random?rating=pg&api_key=dc6zaTOxFJmzC&tag=%s",q)
		g:=new(gifyout)
		resp,_:=http.Get(myurl)
		dec:= json.NewDecoder(resp.Body)
		dec.Decode(g)
		e.Respond(g.Data.Image_url)
	},
}

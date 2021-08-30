package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {

	data := url.Values{
		"user": {"****"},
		"pass": {"***"},
	}
	jar, _ := cookiejar.New(nil)

	var Lcookies []*http.Cookie

	c := &http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println("_______________________")
			var cookies []*http.Cookie
			for _, d := range req.Response.Cookies() {
				switch d.Name {
				case "PHPSESSID":
					cookies = append(cookies, &http.Cookie{
						Name:  "PHPSESSID",
						Value: d.Value,
					})
				case "XSRF-TOKEN":
					cookies = append(cookies, &http.Cookie{
						Name:  "XSRF-TOKEN",
						Value: d.Value,
					})
				case "shop_login":
					cookies = append(cookies, &http.Cookie{
						Name:  "shop_login",
						Value: d.Value,
					})
				case "shop_session":
					cookies = append(cookies, &http.Cookie{
						Name:  "shop_session",
						Value: d.Value,
					})
				case "uid":
					cookies = append(cookies, &http.Cookie{
						Name:  "uid",
						Value: d.Value,
					})
				case "rem":
					cookies = append(cookies, &http.Cookie{
						Name:  "rem",
						Value: d.Value,
					})
				case "key":
					cookies = append(cookies, &http.Cookie{
						Name:  "key",
						Value: d.Value,
					})
				}
			}

			Lcookies = append(Lcookies, cookies...)

			return nil
		},
		Jar:     jar,
		Timeout: 0,
	}

	resp, err := c.PostForm("https://admin.rankedgaming.com/accounts/game-logs.php?action=login", data)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
	if len(Lcookies) == 0 {
		log.Println("no paso nada aqui esto es una mierda")
	}
	url, _ := url.Parse("https://admin.rankedgaming.com/accounts/game-logs.php")
	c.Jar.SetCookies(url, Lcookies)
	req, err := http.NewRequest("POST", "https://admin.rankedgaming.com/accounts/game-logs.php", nil)
	if err != nil {
		log.Println("Error en el segundo post")
	}
	res, err := c.Do(req)
	if err != nil {
		log.Println("error en el segundo post x2")
	}
	_, err = ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error aqui")
	}
	log.Println(res.Header)
	//log.Println(string(body))
}

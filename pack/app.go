package pack

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/anaskhan96/soup"
)

func Tested() {
	resp, err := soup.Get("https://xkcd.com")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}
func Get_action() {
	resp, err := soup.Get("http://testphp.vulnweb.com/login.php")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "story").FindAll("form")
	for _, link := range links {
		fmt.Println("Link :", link.Attrs()["action"])
	}
}

func Run(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//sb := string(body)
	log.Print(body)

}

/*POST METHOD*/
func Post() {
	data := url.Values{
		"uname": {"test"},
		"pass":  {"test"},
	}

	resp, err := http.PostForm("http://testphp.vulnweb.com/userinfo.php", data)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Println(sb)
	/*  TESTED    */
	log.Println(resp.Header)
	log.Println(resp.Cookies())
	log.Println(resp.Location())

}

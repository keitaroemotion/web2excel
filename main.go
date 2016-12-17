package main
import "fmt"
import "net/http"

func main(){
  url := "https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html"
  response, _ := http.Get(url)
  fmt.Println(response)
}

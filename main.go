package main
import "fmt"
import "io/ioutil"
import "net/http"

func main(){
  fmt.Println(GetResponseBody("https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html"))
}

func GetResponseBody(url string) string{
  response, _ := http.Get(url)
  bytes, _    := ioutil.ReadAll(response.Body)
  response.Body.Close()
  return string(bytes)
}

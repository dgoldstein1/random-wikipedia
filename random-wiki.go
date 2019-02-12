package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/davecgh/go-spew/spew"
    "io/ioutil"
)

type RArticleResp struct {
    Query       RQuery      `json:"query"`
}

type RQuery struct {
    Pages       map[string]interface{}      `json:"pages"`
}


func randomArticle(w http.ResponseWriter, r *http.Request) {

    url := "https://www.mediawiki.org/w/api.php?action=query&generator=random&grnnamespace=0&grnlimit=1&prop=info&format=json"
    res, err := http.Get(url)
    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }

    rArticle := &RArticleResp{} // or &Foo{}
    err = json.Unmarshal(body, &rArticle)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    spew.Dump(rArticle)

    fmt.Fprintf(w, "temp")
}

func main() {
    http.HandleFunc("/randomArticle", randomArticle) // set router
    http.Handle("/metrics", promhttp.Handler())
	log.Println("Serving on port 8080")
    err := http.ListenAndServe(":8080", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
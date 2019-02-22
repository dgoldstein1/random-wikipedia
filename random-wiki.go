package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "io/ioutil"
)

type RArticleResp struct {
    Query RQuery        `json:"query"`
}

type RQuery struct {
    Pages  map[string]Page     `json:"pages"`
}

type Page struct {
    Extract interface{} `json:"extract"`
}


func randomArticle(w http.ResponseWriter, r *http.Request) {
    log.Printf("Method: '%s',url: '%s'",  r.Method, r.URL)
    // make request
    url := "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exlimit=max&explaintext&exintro&generator=random&grnnamespace=0&grnlimit=1ts="
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
    
    text := ""
    for _, v := range rArticle.Query.Pages {
        text = fmt.Sprint(v.Extract)
        break
    }
    
    fmt.Fprintf(w, text)
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
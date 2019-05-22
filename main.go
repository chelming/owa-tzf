package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "regexp"
    )

func handler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get(os.Getenv("URL"))
    if err != nil {
        // handle error
        log.Fatal(err)
    }
    defer resp.Body.Close()
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    re := regexp.MustCompile("TZID=UTC")
    calendar := re.ReplaceAllString(bodyString, "TZID="+os.Getenv("TZID"))
    fmt.Fprintf(w, "%s", calendar)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

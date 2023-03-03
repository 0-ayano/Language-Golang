package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

// test server
const URL = "http://localhost:8080/somePost" 

// sample structure
//// ポイント1
type Sample struct {
    Id   int    `json: "id"`
    Name string `json: "name"`
}

func main() {
    // build sample structure
    sample := new(Sample)
    sample.Id = 0
    sample.Name = "hoge"

    // encode json
    sample_json, _ := json.Marshal(sample)
    fmt.Printf("[+] %s\n", string(sample_json))

    // send json
    //// ポイント2, 3
    res, err := http.Post(URL, "application/json", bytes.NewBuffer(sample_json))
    defer res.Body.Close()

    if err != nil {
        fmt.Println("[!] " + err.Error())
    } else {
        fmt.Println("[*] " + res.Status)
    }
}

/* 
参考
- [Golangで構造体をJSONに変換してPOSTする](https://qiita.com/howmuch515/items/65094d18292ff1c4408c)


package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "flag"
    //"log"
)

type Page struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Snip   string `json:"snip"`
}

func (p Page) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func main() {
	


    //var ip = flag.Int("flagname", 1234, "help message for flagname")
    
    pages := getPages()
    for _, p := range pages {
        fmt.Println(p.Snip)
        out, err := exec.Command("powershell", p.Snip).Output()
        if err != nil {
          //log.Fatal(err)
        }
        fmt.Printf("The date is %s\n", out)
        
     }
   

    fmt.Println(toJson(pages))
}

func getPages() []Page {
	wordPtr := flag.String("filename", "./snips.json", "The name to run")
	flag.Parse()
	 fmt.Println("word:", *wordPtr)
    raw, err := ioutil.ReadFile(*wordPtr)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Page
    json.Unmarshal(raw, &c)
    return c
}

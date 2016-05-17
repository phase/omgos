package main

import (
    "os"
    "fmt"
    "strings"
    "os/exec"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

var config map[string]interface{}

func readConfig() {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        fmt.Println("  Error:", err);
    } else {
        var c map[string]interface{}
        if err = json.Unmarshal(data, &c); err != nil {
            fmt.Println("  Error:", err);
            return
        }
        config = c
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    fmt.Printf("GET %s\n", path)
    
    //Get file extension
    s := strings.Split(path, ".")
    extension := s[len(s) - 1]
    
    //TODO: Clean up code
    if(config[extension] != nil) {
        command := config[extension].(string)
        pwd, _ := os.Getwd()
        command = strings.Replace(command, "$file", pwd + "/" + path, -1)
        fmt.Println(" Running command:", command)
        splitCommand := strings.Split(command, " ")
        data, err := exec.Command(splitCommand[0], splitCommand[1:]...).Output()
        if err != nil {
            fmt.Fprintf(w, "Error: %q\n", err)
        } else {
            w.Write(data)
        }
    } else {
        data, err := ioutil.ReadFile(path)
        if err != nil {
            fmt.Fprintf(w, "Error: %q\n", err)
        } else {
            w.Write(data)
        }
    }
}

func main() {
    fmt.Println("Reading ./config.json")
    readConfig()
    fmt.Println("Starting HTTP Server")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
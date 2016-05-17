package main

import (
    "os"
    "fmt"
    "log"
    "strings"
    "os/exec"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

var info *log.Logger
var config map[string]interface{}

func readConfig() {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        info.Println("Error:", err);
    } else {
        var c map[string]interface{}
        if err = json.Unmarshal(data, &c); err != nil {
            info.Println("Error:", err);
            return
        }
        config = c
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    info.Printf("GET %s\n", path)
    
    //Get file extension
    s := strings.Split(path, ".")
    extension := s[len(s) - 1]
    
    var data []u8
    var err Error
    if config[extension] != nil {
        command := config[extension].(string)
        command = strings.Replace(command, "$file", path, -1)
        info.Println("Running command:", command)
        splitCommand := strings.Split(command, " ")
        data, err := exec.Command(splitCommand[0], splitCommand[1:]...).Output()
    } else {
        data, err := ioutil.ReadFile(path)
    }
    
    if err != nil {
        fmt.Fprintf(w, "Error: %q\n", err)
    } else {
        w.Write(data)
    }
}

func main() {
    info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    info.Println("Reading ./config.json")
    readConfig()
    info.Println("Starting HTTP Server")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
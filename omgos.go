package main

import (
    "os"
    "fmt"
    "log"
    "regexp"
    "strings"
    "os/exec"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

var info *log.Logger
var commands map[string]interface{}
var config map[string]interface{}

func readJson(file string, obj *map[string]interface{}) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        info.Println("Error:", err)
    } else {
        if err = json.Unmarshal(data, obj); err != nil {
            info.Println("Error:", err)
            return
        }
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    info.Printf("GET %s\n", path)
    
    //Check blocked files
    blocked := config["blocked"].([]interface{})
    for _, block := range blocked {
        match, _ := regexp.MatchString(block.(string), path)
        if match {
            //Blocked!
            fmt.Fprintf(w, "403 Forbidden: %s", path)
            return
        }
    }
    
    //Get file extension
    s := strings.Split(path, ".")
    extension := s[len(s) - 1]
    
    var data []byte
    var err error
    if commands[extension] != nil {
        command := commands[extension].(string)
        command = strings.Replace(command, "$file", path, -1)
        info.Println("Running command:", command)
        splitCommand := strings.Split(command, " ")
        data, err = exec.Command(splitCommand[0], splitCommand[1:]...).Output()
    } else {
        data, err = ioutil.ReadFile(path)
    }
    
    if err != nil {
        fmt.Fprintf(w, "Error: %q\n", err)
    } else {
        w.Write(data)
    }
}

func main() {
    info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    
    info.Println("Reading ./commands.json")
    readJson("commands.json", &commands)
    info.Println("Reading ./config.json")
    readJson("config.json", &config)
    
    info.Println("Starting HTTP Server")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
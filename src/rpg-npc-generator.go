package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "time"
    "math/rand"
)

type Config struct {
    Picks      []map[string]int            `yaml:"picks"`
    Attributes map[string][]map[string]int `yaml:"attributes"`
}

func main() {
    var result []string
    var arg_file string = os.Args[1]
    var config Config
    seed := rand.NewSource(time.Now().UnixNano())
    random := rand.New(seed)
    data, _ := ioutil.ReadFile(arg_file)
    err := yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    for p := range config.Picks {
        result = result[:0]
        for pk, pv := range config.Picks[p] {
            for pn := 0; pn < pv; pn++ {
                for ak, _ := range config.Attributes[pk][random.Intn(len(config.Attributes[pk]))] {
                    if stringNotInSlice(ak, result) {
                        result = append(result, ak)
                    } else {
                        pn -= 1
                    }
                }
            }
        }
        for i := range result {
            fmt.Printf("%s\n", result[i])
        }
    }
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func stringNotInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return false
        }
    }
    return true
}

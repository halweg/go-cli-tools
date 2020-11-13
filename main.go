package main

import (
    "flag"
    "log"
)

func main() {
    var name string
    //var n string

    flag.StringVar(&name, "name", "go-cli-tools", "help info")
    flag.StringVar(&name, "n", "go-cli-tools", "help info")
    flag.Parse()


    log.Printf("name : %s\n", name)
    //log.Printf("n : %s", n)
}

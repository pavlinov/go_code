package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

type ISTRINGTABLE struct {
    XMLName  xml.Name `xml:"ISTRINGTABLE"`
    ID       string   `xml:"ID,attr"`
    LANG     string   `xml:"LANG,attr"`
    ISTRINGS []ISTRING `xml:"ISTRING"`
}

type ISTRING struct {
    KEY   string   `xml:"KEY"`
    VALUE MyValue `xml:"VALUE"`
}

type MyValue struct {
    XML string `xml:",innerxml"`
}

func main() {
    xmlFile, err := os.Open("example.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    byteValue, _ := ioutil.ReadAll(xmlFile)

    var iStringTable ISTRINGTABLE
    err = xml.Unmarshal(byteValue, &iStringTable)
    if err != nil {
        fmt.Println("Error parsing XML:", err)
        return
    }

    for _, iString := range iStringTable.ISTRINGS {
        fmt.Println("KEY:", iString.KEY)
        fmt.Println("VALUE:", iString.VALUE.XML)
    }
}
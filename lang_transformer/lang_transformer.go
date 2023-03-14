package main

import "os"
import "fmt"
import "log"
import "io/ioutil"
import "encoding/xml"
import "strings"
import "gopkg.in/yaml.v2"

type Istringtable struct {
    XMLName  xml.Name  `xml:"ISTRINGTABLE"`
    Istrings []Istring `xml:"ISTRING"`
    Id       string   `xml:"ID,attr"`
    Lang     string   `xml:"LANG,attr"`
}

type Istring struct {
    Key   string `xml:"KEY"`
    Value MyHtml `xml:"VALUE"`
}

type MyHtml struct {
    XML string `xml:",innerxml"`
}

func main() {
    map_files := []string{
        "al.strings: application_app_name_1_0.al_AL.yml",
        "br.strings: application_app_name_1_0.pt_BR.yml",
        "de.strings: application_app_name_1_0.de_DE.yml",
        "el.strings: application_app_name_1_0.el_GR.yml",
        "en.strings: application_app_name_1_0.en_US.yml",
        "esMX.strings: application_app_name_1_0.es_MX.yml",
        "fr.strings: application_app_name_1_0.fr_FR.yml",
        "frCA.strings: application_app_name_1_0.fr_CA.yml",
        "frFR.strings: application_app_name_1_0.fr_FR.yml",
        "hr.strings: application_app_name_1_0.hr_HR.yml",
        "hu.strings: application_app_name_1_0.hu_HU.yml",
        "it.strings: application_app_name_1_0.it_IT.yml",
        "pt.strings: application_app_name_1_0.pt_PT.yml",
        "ro.strings: application_app_name_1_0.ro_RO.yml",
        "sp.strings: application_app_name_1_0.es_ES.yml",
    }


    en_langs, err := getLangs("new_en.strings.xml")
    if err != nil {
        fmt.Println("Error parsing us_EN stings: ", err)
        log.Fatal(err)
        return
    }

    fmt.Println("Parsed EN strings: ", en_langs)

    enLangMap := langsToDict(en_langs.Istrings)
    fmt.Println(strings.Repeat("Alexey\n", 8))
    fmt.Println(enLangMap["LANG_help_results_overview"])
    fmt.Println(strings.Repeat("===\n", 8))

/*
    for _, en_string := range en_langs.Istrings {
            fmt.Printf("en_strings: %+v\n",en_string)
            //os.Exit(0)
    }
*/

    // Create a map from the file names
    m := make(map[string]string)
    for _, f := range map_files {
        parts := strings.Split(f, ": ")
        if len(parts) == 2 {
            m[parts[0]] = parts[1]

            process_file := parts[0]
            output_file := parts[1]



            br_langs, err := getLangs(process_file)
            if err != nil {
                log.Fatal(fmt.Sprintf("Error parsing us_EN stings: ", err))
                return
            }

            yamlLangsMap := make(map[string]string)
            for _, item := range br_langs.Istrings {
                fmt.Printf("KEY: %s\nVALUE: %s\n\n", item.Key, string(item.Value.XML))
                fmt.Printf("EN VALUE: %s\nVALUE: %s\n\n", enLangMap[item.Key], string(item.Value.XML))
                yamlLangsMap[string(enLangMap[item.Key])] = string(item.Value.XML)
            }

            fmt.Printf("yamlLangsMap > %+v\n", yamlLangsMap)

            err = saveToYaml(output_file, yamlLangsMap)
            if err != nil {
                fmt.Println(err)
                return
            }
/*
            fmt.Printf("Parsed BR strings: %+v\n", br_langs)
            brLangMap := langsToDict(br_langs.Istrings)
            fmt.Println(strings.Repeat("pt_BR\n", 9))
            fmt.Println(brLangMap["LANG_formtool_help"])
            fmt.Println(strings.Repeat("===\n", 9))
            */

        }
    }

    // Print the map
    fmt.Println(m)

    fmt.Fprintf(os.Stdout, "STDOUT Finiched OK")
}

func getLangs(filename string) (*Istringtable, error) {
    fmt.Println("Parsing XML:", filename)
    xml_data, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return nil, err
    }
    defer xml_data.Close()

    //fmt.Fprintf(os.Stdout, "length xml_data: %v\n",len(xml_data))

    var lang_strings Istringtable
    decoder := xml.NewDecoder(xml_data)
    err = decoder.Decode(&lang_strings)
    if err != nil {
        fmt.Println("Error decoding XML:", err)
        fmt.Fprintf(os.Stderr, "Error decoding en_xml_data: %v\n", err)
    }

    for _, en_string := range lang_strings.Istrings {
            fmt.Printf("strings > %+v\n",en_string)
    }

    return &lang_strings, nil
}

func langsToDict(lang_strings []Istring) map[string]string {
    langMap := make(map[string]string)
    for _, lang := range lang_strings {
        value := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(string(lang.Value.XML), "\n", ""), "\t", ""), "\\s+", "")
        //value := strings.ReplaceAll(string(lang.Value.XML), "\\s+", "")
        langMap[lang.Key] = value
    }

    return langMap
}


func saveToXml(filename string, langs *Istringtable) error {
    xmlBytes, err := xml.MarshalIndent(langs, "", "    ")
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(filename, xmlBytes, 0644)
    if err != nil {
        return err
    }

    return nil
}


func saveToYaml(filename string, langs map[string]string) error {
    yamlBytes, err := yaml.Marshal(langs)
    if err != nil {
        return fmt.Errorf("Error encoding YAML: %v", err)
    }

    err = ioutil.WriteFile(filename, yamlBytes, 0644)
    if err != nil {
        return fmt.Errorf("Error writing file: %v", err)
    }

    fmt.Println("Saved langs to file: ", filename)
    return nil
}

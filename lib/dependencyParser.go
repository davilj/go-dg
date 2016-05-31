package dg

import (
  "regexp"
  "fmt"
  "io/ioutil"
  "os"
  "path/filepath"
  "encoding/xml"
)

type node struct {
  name string
  //version to list of dependencies
  version2Dependencies map[string][]node
}

type contentHandler func (content string) XMLivyModule

type fileHandler func (file string, contentParser contentHandler) XMLivyModule


func extractDepencies(content string) XMLivyModule {
  d:=XMLivyModule{}
  xml.Unmarshal([]byte(content), &d)
  return d
}

func BuildDB(baseDir string) DB {
  allIvyModules := WalkDir(baseDir, fileParser, extractDepencies)
  fmt.Println(len(allIvyModules), " ivy modules to sort")
  return Load(allIvyModules)
}

func fileParser(file string, aContentHandler contentHandler) XMLivyModule {
  fmt.Println("Handling file: " + file)
  dat, err := ioutil.ReadFile(file)
  if err!=nil {
    panic(err)
  }

  content := aContentHandler(string(dat))
  return content
}

func WalkDir(baseDir string, aFilehandler fileHandler, aContentHandler contentHandler) []XMLivyModule {
  //ivy-1.0.0.beta.xml
  fileRegx := regexp.MustCompile("^ivy-[0-9]\\.[0-9]\\.[0-9]\\.(beta|final)\\.xml$")

  files, err := ioutil.ReadDir(baseDir)
  if err != nil {
    panic("Could not read base dir")
  }

  ivyModules := make([]XMLivyModule,0)

  for _,file := range(files) {
    fileName := file.Name()
    fileLocation := baseDir + string(filepath.Separator) + fileName
    if fileRegx.MatchString(fileName)==true {
      ivyModules = append(ivyModules, aFilehandler(fileLocation, aContentHandler))
    }

    info, err := os.Stat(fileLocation)
    if (err!=nil) {
      panic(err)
    }

    if info.IsDir() {
      newList:= WalkDir(fileLocation, aFilehandler, aContentHandler)
      ivyModules = append(ivyModules, newList...)
    }
  }
  return ivyModules
}

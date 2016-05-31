package dg

import (
  "encoding/xml"
)

//package version, map to pacakage, version
type DB struct {
  DbLink  map[string][]string
  DbInfo  map[string]XMLivyModule
}

type XMLInfo struct {
  XMLName xml.Name  `xml:"info"`
  Org string        `xml:"organisation,attr"`
  Module string     `xml:"module,attr"`
  Rev string        `xml:"revision,attr"`
  Status string     `xml:"status,attr"`
  Pub string        `xml:"publication,attr"`
}

type XMLDependency struct {
  XMLName xml.Name      `xml:"dependency"`
  Org string            `xml:"org,attr"`
  Name string           `xml:"name,attr"`
  Rev string            `xml:"rev,attr"`
  RevConstraint string  `xml:"revConstraint,attr"`
  Conf string           `xml:"conf,attr"`
}

type XMLDependencies struct {
  XMLName xml.Name          `xml:"dependencies"`
  Dependency []XMLDependency  `xml:"dependency"`
}

type XMLivyModule struct {
  XMLName xml.Name              `xml:"ivy-module"`
  Info XMLInfo                  `xml:"info"`
  Dependencies XMLDependencies  `xml:"dependencies"`
}

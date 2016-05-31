// Package stringutil contains utility functions for working with strings.
package main


import (
	"os"
	"fmt"
	"bytes"
	"bufio"
	dg "github.com/davilj/dg/lib"
)

func main() {
	p:=fmt.Println
	//the base location for package information, usually ivy dir
	dir := os.Args[0]
	//package to build dependency tree off
	package2Search := os.Args[1]

	p("Ready to parse: ", dir)
	fileContent:=Parse2Map(dir, package2Search)
	Write2File(fileContent)
}

func Write2File(graph string) {
	fmt.Println("Writint file")
	fileHandle, _ := os.Create("data.js")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, graph)
	writer.Flush()
}

// location of main osgi application
func Parse2Map(dir string, package2Search string) string{
	var db = dg.BuildDB(dir)
	var nodesBuffer bytes.Buffer
	var linksBuffer bytes.Buffer
	for node,links := range(db.DbLink) {
		nodesBuffer.WriteString(buildNode(node))
		for _,link := range(links) {
			linksBuffer.WriteString(buildLink(node, link))
			nodesBuffer.WriteString(buildNode(link))
		}
	}

	var allDB bytes.Buffer;
	allDB.WriteString("var elements2 = [")
	allDB.WriteString(nodesBuffer.String())
	allDB.WriteString(linksBuffer.String())
	var finalStr = allDB.String()
	var test = finalStr[:len(finalStr)-2]

	var endStr bytes.Buffer;
	endStr.WriteString(test)
	endStr.WriteString("]")
	return endStr.String()

}

func buildNode(nodeName string) string{
	var buffer bytes.Buffer
	buffer.WriteString("{ group: 'nodes', data: { id: '")
	buffer.WriteString(nodeName)
	buffer.WriteString("' }  },\n")

	s1 := buffer.String()
	return s1
}

func buildLink(from string, to string) string {
	var buffer bytes.Buffer
	buffer.WriteString("{ group: 'edges', data: { id: '")
	buffer.WriteString(from)
	buffer.WriteString("-")
	buffer.WriteString(to)
	buffer.WriteString("', source: '")
	buffer.WriteString(from)
	buffer.WriteString("', target: '")
	buffer.WriteString(to)
	buffer.WriteString("' } },\n")

	s1 := buffer.String()
	return s1
}

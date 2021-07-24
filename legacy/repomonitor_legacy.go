package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"path/filepath"
	//"regexp"
)

func main() {
	file, _ := ioutil.ReadFile("repoqueue.dat")

	s := string(file)

	data := strings.Replace(s, ",", "", -1)
	//data := strings.Fields(s)

	datatl := strings.Replace(data, "[", "", -1)
	datatr := strings.Replace(datatl, "]", "", -1)

	datar := strings.Replace(datatr, "Url:", "", -1)
	datas := strings.Fields(datar)
	//t := strings.Replace(data, ",", "", -1)

	for _, k := range datas {
		d := fmt.Sprint(strings.Replace(k, `https://github.com/`, "", -1))
		c := string(fmt.Sprint(strings.Split(d, "/")))
		t := strings.Trim(c, "[[")
		r := strings.Trim(t, "]]")
		//g := fmt.Sprintln(strings.Fields(r))
		e := strings.Fields(r)
		//fmt.Print("UserName: ", g[0])
		//fmt.Print("RepoName: ", g[1])
		//username := e[0]
		//reponame := e[1]
		//userpath := filepath.Join(".", "repos/e[0]")
		//err := os.MkdirAll(userpath, os.ModePerm)
		//if err == nil {
        	    //return nil
		//}
		repopath := filepath.Join(".", "repos", e[0], e[1])
		err := os.MkdirAll(repopath, os.ModePerm)
		if err != nil {
        	    panic(err)
		}

		fmt.Println(e[0], e[1], "success")

		//fmt.Println(regexp.MustCompile(`"([^"]*)" *github.com/`), k)
		//fmt.Println(regexp.MatchString
		//fmt.Println("github.com", k)
		}

	//fmt.Println(strings.Trim(fmt.Sprint(datas), "[]"))
	//fmt.Println("1st Field: ", datas[:])
	//fmt.Println(k)
}

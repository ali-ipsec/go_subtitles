package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Episode struct {
	mkvName string
	strName string
}

//var regex = regex.MustCompile(`S\d+E\d+`)
func (r Episode) rename(path string) {
	mkvName := strings.TrimSuffix(r.mkvName, ".mkv")
	fmt.Println(mkvName)
	// fmt.Println(path+"/"+r.strName, path+"/"+mkvName+".srt")
	os.Rename(path+"/"+r.strName, path+"/"+mkvName+".srt")
}
func Renamify(path string) {
	episods := make(map[int]Episode)
	var err error
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	l := len(files)
	for _, file := range files {

		name := file.Name()
		ext := filepath.Ext(name)

		re := regexp.MustCompile(`S\d+E\d+`)
		SE := re.FindStringSubmatch(name)

		if len(SE) != 0 {
			fmt.Println(ext, name)
			var season int
			var epi int
			fmt.Sscanf(SE[0], "S%dE%d", &season, &epi)
			fmt.Println(SE[0], season, epi)
			entry, ok := episods[epi]
			if !ok {
				entry = Episode{}
			}

			if ext == ".mkv" {
				entry.mkvName = name
			}
			if ext == ".srt" {
				entry.strName = name
			}
			episods[epi] = entry
			// fmt.Println(SE[0])
		}
	}
	for i := 0; i < l; i++ {
		fmt.Println(episods[i].mkvName, episods[i].strName)
		if episods[i].mkvName != "" && episods[i].strName != "" {
			episods[i].rename(path)
		}
	}

}

func main() {

	mydir, err := os.Getwd()

	if err != nil {
		log.Println(err)

	}
	Renamify(mydir)

}

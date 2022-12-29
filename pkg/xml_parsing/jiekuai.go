package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type title1 struct {
	XMLName xml.Name `xml:"dict"`
	Key     []string `xml:"key"`
	Array   []struct {
		XMLName xml.Name `xml:"array"`
		Dict    []struct {
			XMLName xml.Name `xml:"dict"`
			Key     []string `xml:"key"`
			Integer []int    `xml:"integer"`
			String  []string `xml:"string"`
			Array   []struct {
				XMLName xml.Name `xml:"array"`
				Dict    []struct {
					XMLName xml.Name `xml:"dict"`
					Key     []string `xml:"key"`
					Integer []int    `xml:"integer"`
					String  []string `xml:"string"`
				} `xml:"dict"`
			} `xml:"array"`
		} `xml:"dict"`
	} `xml:"array"`
}

func main() {
	var (
		url1            = "D:\\ActiveFile\\工作文件\\temp1.txt"
		url2            = "D:\\ActiveFile\\工作文件\\temp2.txt"
		section, piece string
		temp           = "D:\\ActiveFile\\工作文件\\3879967"
	)
	/*peo := new(title)
	//subject_1150_chapter_462055.plist
	//subject_71_chapter_26919.plist
	var fileName = "subject_71_chapter_26919_section_26927_node_26954.plist"
	file, err := os.Open(temp + "\\" + fileName)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	str := string(data)
	str = strings.ReplaceAll(str, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "")
	str = strings.ReplaceAll(str, "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">", "")
	str = strings.ReplaceAll(str, "<plist version=\"1.0\">", "")
	err = xml.Unmarshal([]byte(str), peo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v1 := range peo.Array {
		for _, v2 := range v1.Dict {
			fmt.Println(v2.Key[0], v2.Integer[0], v2.Key[2], v2.Integer[1], v2.Key[3], v2.Integer[2], v2.Key[4], v2.String[0], v2.Key[5], v2.String[1])
			fmt.Println(v2.Key[1])
			for _, v3 := range v2.Array {
				for _, v4 := range v3.Dict {
					for x, y := range v4.Key {
						switch y {
						case "lastLevelNodeId":
							fmt.Println(y, v4.Integer[x])
						case "questionCount":
							fmt.Println(y, v4.Integer[x-1])
						case "lastLevelNodeName":
							fmt.Println(y, v4.String[0])
						}
					}
					//fmt.Println(v4.Key[0], v4.Integer[0], v4.Key[1], v4.Integer[1], v4.Key[2], v4.Integer[2], v4.Key[3], v4.String[0], v4.Key[4], v4.Integer[3], v4.Key[5], v4.Integer[4], v4.Key[6], v4.Integer[5])
				}
			}
			fmt.Println("---------------------------------------")
		}
	}*/
	dir, err := ioutil.ReadDir(temp)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _, c := range dir {
		match, err := regexp.MatchString("subject_([0-9]+)_chapter_([0-9]+).plist", c.Name())
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		if match {
			peo := new(title1)
			file, err := os.Open(temp + "\\" + c.Name())
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
			str := string(data)
			str = strings.ReplaceAll(str, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "")
			str = strings.ReplaceAll(str, "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">", "")
			str = strings.ReplaceAll(str, "<plist version=\"1.0\">", "")
			err = xml.Unmarshal([]byte(str), peo)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			i := strings.Split(c.Name(), "_")
			split := strings.Split(i[len(i)-1], ".")
			//fmt.Println(c.Name())
			for _, v1 := range peo.Array {
				for _, v2 := range v1.Dict {
					//fmt.Println(v2.Key[0], v2.Integer[0], v2.Key[2], v2.Integer[1], v2.Key[3], v2.Integer[2], v2.Key[4], v2.String[0], v2.Key[5], v2.String[1])
					//fmt.Println(v2.Key[1])
					for _, v3 := range v2.Array {
						for _, v4 := range v3.Dict {
							var lastLevelNodeId, questionCount int
							var lastLevelNodeName string
							for x, y := range v4.Key {
								switch y {
								case "lastLevelNodeId":
									//fmt.Println(y, v4.Integer[x])
									lastLevelNodeId = v4.Integer[x]
								case "questionCount":
									//fmt.Println(y, v4.Integer[x-1])
									questionCount = v4.Integer[x-1]
								case "lastLevelNodeName":
									//fmt.Println(y, v4.String[0])
									lastLevelNodeName = v4.String[0]
								}
							}
							pieceSql := fmt.Sprintf(`insert into t_piece (id, name, question_count, section_id) 
								values (%d, '%s', %d, %d);`, lastLevelNodeId, lastLevelNodeName, questionCount, v2.Integer[2])
							piece += pieceSql
						}
					}
					sectionSql := fmt.Sprintf(`insert into t_section (id, sequence, name, question_count, chapter_id) 
						values (%d, '%s', '%s', %d, %s);`, v2.Integer[2], v2.String[1], v2.String[0], v2.Integer[1], split[0])
					section += sectionSql
				}
			}
			//fmt.Println("---------------------------------------")
		}
	}

	fmt.Println("开始写入section：", section)
	err = ioutil.WriteFile(url1, []byte(section), 0666)
	if err != nil {
		fmt.Println("section写入失败")
	} else {
		fmt.Println("section写入成功")
	}
	fmt.Println("开始写入piece：", piece)
	err = ioutil.WriteFile(url2, []byte(piece), 0666)
	if err != nil {
		fmt.Println("piece写入失败")
	} else {
		fmt.Println("piece写入成功")
	}
}

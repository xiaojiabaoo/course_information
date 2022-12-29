package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type title2 struct {
	XMLName xml.Name `xml:"dict"`
	Key     []string `xml:"key"`
	Integer []int    `xml:"integer"`
	Array   []struct {
		XMLName xml.Name `xml:"array"`
		Dict    []struct {
			XMLName xml.Name  `xml:"dict"`
			Key     []string  `xml:"key"`
			Integer []int     `xml:"integer"`
			String  []string  `xml:"string"`
			Real    []float64 `xml:"real"`
			Array   [][]struct {
				XMLName xml.Name `xml:"array"`
				Dict    []struct {
					XMLName xml.Name  `xml:"dict"`
					Key     []string  `xml:"key"`
					Integer []int     `xml:"integer"`
					String  []string  `xml:"string"`
					Real    []float64 `xml:"real"`
				} `xml:"dict"`
			} `xml:"array"`
		} `xml:"dict"`
	} `xml:"array"`
}

func main() {
	var (
		url1             = "D:\\ActiveFile\\工作文件\\temp1.txt"
		url2             = "D:\\ActiveFile\\工作文件\\temp2.txt"
		question, option string
		temp             = "D:\\ActiveFile\\工作文件\\3879967"
		//err              error
	)
	dir, err := ioutil.ReadDir(temp)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _, c := range dir {
		match, err := regexp.MatchString("subject_([0-9]+)_chapter_([0-9]+)_section_([0-9]+)_node_([0-9]+).plist", c.Name())
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		if match {
			question, option, err = queryData(temp, question, option, c.Name())
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}
		}
	}
	//fmt.Println(url1, url2)
	//fmt.Println(question)
	//fmt.Println(option)
	fmt.Println("开始写入section：", question)
	err = ioutil.WriteFile(url1, []byte(question), 0666)
	if err != nil {
		fmt.Println("question写入失败")
	} else {
		fmt.Println("question写入成功")
	}
	fmt.Println("开始写入option：", option)
	err = ioutil.WriteFile(url2, []byte(option), 0666)
	if err != nil {
		fmt.Println("option写入失败")
	} else {
		fmt.Println("option写入成功")
	}
	//test(temp, question, option)
}

func queryData(temp, question, option, fileName string) (string, string, error) {
	peo := new(title2)
	file, err := os.Open(temp + "\\" + fileName)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", "", err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", "", err
	}
	str := string(data)
	str = strings.ReplaceAll(str, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "")
	str = strings.ReplaceAll(str, "<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">", "")
	str = strings.ReplaceAll(str, "<plist version=\"1.0\">", "")
	err = xml.Unmarshal([]byte(str), peo)
	if err != nil {
		fmt.Println(err.Error())
		return "", "", err
	}
	i := strings.Split(fileName, "_")
	split := strings.Split(i[len(i)-1], ".")
	for _, v1 := range peo.Array {
		for _, v2 := range v1.Dict {
			var questionType string
			for k3, v3 := range v2.Key {
				if v3 == "questionType" {
					switch k3 {
					case 10:
						questionType = v2.String[5]
					case 13:
						questionType = v2.String[7]
					case 12:
						if len(v2.String) <= 7 {
							questionType = v2.String[6]
						} else {
							questionType = v2.String[7]
						}
					}
				}
			}
			switch questionType {
			//单项选择题、多项选择题、判断题、填空题、简答题
			case "SINGLE_CHOICE", "MULTI_CHOICE", "JUDGE_CHOICE", "ORDER_FILL_BLANK", "DISORDER_FILL_BLANK", "ESSAY":
				var (
					analysis          = v2.String[0]
					analysisType      = v2.Integer[0]
					avgCorrectRate    = v2.String[1]
					canPhoto          = v2.Integer[1]
					favorite          = v2.Integer[2]
					mainNodeFrequency = v2.String[2]
					mainNodeId        = v2.Integer[3]
					mainNodeName      = v2.String[3]
					questionAnswer    = v2.String[4]
					questionContent   = v2.String[5]
					questionId        = v2.Integer[4]
					questionSource    = v2.String[6]
					//questionsType     = v2.String[7]
					score    = v2.Real[0]
					sequence = v2.Integer[5]
				)
				questionContent = escape(questionContent)
				analysis = escape(analysis)
				questionSql := fmt.Sprintf(`insert into t_question (id, piece_id, sequence, questionType,
								questionSource, questionContent, questionAnswer, score, mainNodeId, mainNodeName,
								mainNodeFrequency, favorite, canPhoto, avgCorrectRate, analysisType, analysis)
								values (%d,%s,%d,'%s','%s','%s','%s','%v',%d,'%s','%s',%d,%d,'%s',%d,'%s');`,
					questionId, split[0], sequence, questionType, questionSource, questionContent,
					questionAnswer, score, mainNodeId, mainNodeName, mainNodeFrequency, favorite, canPhoto,
					avgCorrectRate, analysisType, analysis)
				question += questionSql
				//如果是填空题 就不需要下面的array了
				if questionType == "ORDER_FILL_BLANK" || questionType == "DISORDER_FILL_BLANK" || questionType == "ESSAY" {
					break
				}
				for _, v3 := range v2.Array[0] {
					for _, v4 := range v3.Dict {
						var (
							correct       = v4.Integer[0]
							optionContent = v4.String[0]
							optionTitle   = v4.String[1]
							sequences     = v4.Integer[1]
						)
						optionContent = escape(optionContent)
						optionSql := fmt.Sprintf(`insert into t_option (question_id, sequence, title, 
									content, correct) values (%d, %d, '%s', '%s',  %d);`,
							questionId, sequences, optionTitle, optionContent, correct)
						option += optionSql
					}
				}
			//情景选择题
			case "MANY_TO_MANY":
				var (
					analysis          = v2.String[0]
					analysisType      = v2.Integer[0]
					avgCorrectRate    = v2.String[1]
					canPhoto          = v2.Integer[1]
					favorite          = v2.Integer[2]
					mainNodeFrequency = v2.String[2]
					mainNodeId        = v2.Integer[3]
					questionContent   = v2.String[4]
					questionId        = v2.Integer[4]
					questionSource    = v2.String[5]
					score             = v2.Real[0]
					sequence          = v2.Integer[5]
					questionAnswer    = make([]string, 0, 0)
				)
				//获取答案
				for _, v3 := range v2.Array[1] {
					for k4, v4 := range v3.Dict {
						parsing := escape(v4.String[0])
						questionAnswer = append(questionAnswer, "第"+fmt.Sprintf(`%d`, k4+1)+"处答案："+v4.String[4]+"，解析："+parsing)
					}
				}
				analysis = escape(analysis)
				questionContent = escape(questionContent)
				questionSql := fmt.Sprintf(`insert into t_question (id, piece_id, sequence, questionType,
								questionSource, questionContent, questionAnswer, score, mainNodeId, mainNodeName,
								mainNodeFrequency, favorite, canPhoto, avgCorrectRate, analysisType, analysis)
								values (%d,%s,%d,'%s','%s','%s','%s','%v',%d,'%s','%s',%d,%d,'%s',%d,'%s');`,
					questionId, split[0], sequence, questionType, questionSource, questionContent,
					questionAnswer, score, mainNodeId, "", mainNodeFrequency, favorite, canPhoto,
					avgCorrectRate, analysisType, analysis)
				question += questionSql
				for _, v3 := range v2.Array[0] {
					for _, v4 := range v3.Dict {
						var (
							correct       = v4.Integer[0]
							optionContent = v4.String[0]
							optionTitle   = v4.String[1]
							sequences     = v4.Integer[1]
						)
						optionContent = escape(optionContent)
						optionTitle = escape(optionTitle)
						optionSql := fmt.Sprintf(`insert into t_option (question_id, sequence, title, 
									content, correct) values (%d, %d, '%s', '%s',  %d);`,
							questionId, sequences, optionTitle, optionContent, correct)
						option += optionSql
					}
				}
			// 综合题
			case "COMPREHENSIVE":
				var (
					content, questionAnswer string
					mainNodeName, analysis  string
					mainNodeId              int
					analysisType            = v2.Integer[0]
					avgCorrectRate          = v2.String[1]
					canPhoto                = v2.Integer[1]
					favorite                = v2.Integer[2]
					mainNodeFrequency       = v2.String[2]
					questionContent         = v2.String[3]
					questionId              = v2.Integer[4]
					questionSource          = v2.String[4]
					score                   = v2.Real[0]
					sequence                = v2.Integer[5]
				)
				for _, v3 := range v2.Array[0] {
					for _, v4 := range v3.Dict {
						analysis = v4.String[0]
						mainNodeId = v4.Integer[5]
						mainNodeName = v4.String[3]
						content = v4.String[5]
						questionAnswer = v4.String[4]
					}
				}
				content = questionContent + "<br/>题目：" + content
				content = escape(content)
				questionAnswer = escape(questionAnswer)
				analysis = escape(analysis)
				questionSql := fmt.Sprintf(`insert into t_question (id, piece_id, sequence, questionType,
								questionSource, questionContent, questionAnswer, score, mainNodeId, mainNodeName,
								mainNodeFrequency, favorite, canPhoto, avgCorrectRate, analysisType, analysis)
								values (%d,%s,%d,'%s','%s','%s','%s','%v',%d,'%s','%s',%d,%d,'%s',%d,'%s');`,
					questionId, split[0], sequence, questionType, questionSource, content,
					questionAnswer, score, mainNodeId, mainNodeName, mainNodeFrequency, favorite, canPhoto,
					avgCorrectRate, analysisType, analysis)
				question += questionSql
			}
		}
	}
	return question, option, nil
}

func escape(str string) string {
	str = strings.ReplaceAll(str, "&gt;", ">")
	str = strings.ReplaceAll(str, "&lt;", "<")
	str = strings.ReplaceAll(str, "&amp;", "&")
	str = strings.ReplaceAll(str, "'", "\\'")
	return str
}

func test(temp, question, option string) {
	var fileName = "测试.plist"
	question, option, err := queryData(temp, question, option, fileName)
	if err != nil {
		return
	}
	fmt.Println(question)
	/*fmt.Println("开始写入section：", section)
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
	}*/
}

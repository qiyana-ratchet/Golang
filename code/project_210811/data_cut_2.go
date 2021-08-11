package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type measDataFile struct {
	XMLName    xml.Name   `xml:"measDataFile"`
	FileHeader fileHeader `xml:"fileHeader"`
	measData   measData   `xml:"measData"`
	FileFooter fileFooter `xml:"fileFooter"`
}

type fileHeader struct {
	XMLName    xml.Name   `xml:"fileHeader"`
	FileFormatVersion     string   `xml:"fileFormatVersion,attr"`
	VendorName    string   `xml:"vendorName,attr"`
	DnPrefix    string   `xml:"dnPrefix,attr"`
	FileSender    fileSender   `xml:"fileSender"`
	MeasData	measData	`xml:"measData"`

}

type fileSender struct {
	XMLName    xml.Name   `xml:"fileSender"`
	SenderName string `xml:"senderName,attr"`
	SenderType string `xml:"senderType,attr"`

}

type fileFooter struct {
	XMLName    xml.Name   `xml:"fileFooter"`
	MeasData	measData	`xml:"measData,attr"`
}

type measData struct {
	XMLName    xml.Name   `xml:"measData"`
	MeAsEntity measEntity `xml:"measEntity"`
	MeAsInfo   []measInfo `xml:"measInfo"`
	BeginTime string `xml:"beginTime,attr"`
	EndTime string `xml:"endTime,attr"`
}

//type beginTime struct {
//	XMLName xml.Name `xml:"beginTime"`
//	Key     string   `xml:"beginTime,attr"`
//}

//type endTime struct {
//	XMLName xml.Name `xml:"endTime"`
//	Key     string   `xml:"endTime,attr"`
//}

type measEntity struct {
	XMLName xml.Name `xml:"measEntity"`
	Key     string   `xml:"localDn,attr"`
	Key2    string   `xml:"swVersion,attr"`
}

type measInfo struct {
	XMLName    xml.Name   `xml:"measInfo"`
	MeAsInfoID string     `xml:"measInfoId,attr"`
	Job        job        `xml:"job"`
	GranPeriod granPeriod `xml:"granPeriod"`
	RepPeriod  repPeriod  `xml:"repPeriod"`
	MeAsType   []measType `xml:"measType"`
	MeAsValue  measValue  `xml:"measValue"`
}

type job struct {
	XMLName xml.Name `xml:"job"`
	//XMLAttr xml.Attr `xml:"jobId,attr"`
	Key string `xml:"jobId,attr"`
}

type granPeriod struct {
	XMLName xml.Name `xml:"granPeriod"`
	Key     string   `xml:"duration,attr"`
	Key2    string   `xml:"endTime,attr"`
}

type repPeriod struct {
	XMLName xml.Name `xml:"repPeriod"`
	Key     string   `xml:"duration,attr"`
}

type measType struct {
	XMLName xml.Name `xml:"measType"`
	Key     string   `xml:"p,attr"`
	Value   string   `xml:",chardata"`
}

type measValue struct {
	XMLName xml.Name `xml:"measValue"`
	Key     string   `xml:"measObjLdn,attr"`
	R       []r      `xml:"r"`
}

type r struct {
	XMLName xml.Name `xml:"r"`
	Key     string   `xml:"p,attr"`
	Value   string   `xml:",chardata"`
}

func main() {
	// xml 파일 오픈
	fp, err := os.Open("/home/thkim/GolandProjects/project_210810/data_cut_2.xml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// xml 파일 읽기
	data, err := ioutil.ReadAll(fp)

	// xml 디코딩
	var measDataFile measDataFile
	err = xml.Unmarshal(data, &measDataFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(measDataFile)
}

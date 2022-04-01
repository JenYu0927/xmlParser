package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
)

func getLatestFileName(performanceFolder string) string {
	files, err := ioutil.ReadDir(performanceFolder)
	if err != nil {
		log.Fatal(err)
	}
	var modTime time.Time
	var latestFileName []string
	for _, file := range files {
		if file.Mode().IsRegular() {
			if !file.ModTime().Before(modTime) {
				if file.ModTime().After(modTime) {
					modTime = file.ModTime()
					latestFileName = latestFileName[:0]
				}
				latestFileName = append(latestFileName, file.Name())
			}
		}
	}
	return latestFileName[0]
}

func readGzipFileIntoString(gzipFile *os.File) string {
	gzFile, _ := gzip.NewReader(gzipFile)
	defer gzFile.Close()

	performanceByte, _ := ioutil.ReadAll(gzFile)
	return string(performanceByte)
}

func findMetricIdInMeasType(xmlNode *xmlquery.Node, metricsName string) string {
	XPathExpr := "//measType[text()='" + metricsName + "']/@p"
	return xmlquery.FindOne(xmlNode, XPathExpr).InnerText()
}

func findValueWithMetricId(xmlNode *xmlquery.Node, Id string) string {
	XPathExpr := "//r[@p=" + Id + "]"
	return xmlquery.FindOne(xmlNode, XPathExpr).InnerText()
}

func main() {

	performanceFilesFolder := "/home/jenyu/pm" // going to be imported by telegraf config
	latestFileName := getLatestFileName(performanceFilesFolder)

	performanceFile, err := os.Open(performanceFilesFolder + "/" + latestFileName)
	defer performanceFile.Close()
	if err != nil {
		log.Fatal(err) // need to check how telegraf plugin handle error
	}

	performanceString := readGzipFileIntoString(performanceFile)
	performanceParseToNode, err := xmlquery.Parse(strings.NewReader(performanceString))
	if err != nil {
		log.Fatal(err)
	}

	println("ID:")
	RRC_ConnSuccRate_Id := findMetricIdInMeasType(performanceParseToNode, "RRC.ConnSucc.Rate")
	fmt.Println("RRC_ConnSuccRate_Id:", RRC_ConnSuccRate_Id)

	RRC_ConnDropRate_Id := findMetricIdInMeasType(performanceParseToNode, "RRC.ConnDrop.Rate")
	fmt.Println("RRC_ConnDropRate_Id:", RRC_ConnDropRate_Id)

	NGSetup_SuccessRate_Id := findMetricIdInMeasType(performanceParseToNode, "NG.SetupSuccess.Rate")
	fmt.Println("NGSetup_SuccessRate_Id:", NGSetup_SuccessRate_Id)

	MMHoIntra_SuccRate_Id := findMetricIdInMeasType(performanceParseToNode, "MM.HoIntraSuccRate")
	fmt.Println("MMHoIntra_SuccRate_Id:", MMHoIntra_SuccRate_Id)

	MMHoInter_SuccRate_Id := findMetricIdInMeasType(performanceParseToNode, "MM.HoInterSuccRate")
	fmt.Println("MMHoInter_SuccRate_Id:", MMHoInter_SuccRate_Id)

	MM_HoResInter_SuccRate_Id := findMetricIdInMeasType(performanceParseToNode, "MM.HoResInterSuccRate")
	fmt.Println("MM_HoResInter_SuccRate_Id:", MM_HoResInter_SuccRate_Id)

	NRCellDU_Dl_Throughput_Id := findMetricIdInMeasType(performanceParseToNode, "NRCellDU.DlThroughput")
	fmt.Println("NRCellDU_Dl_Throughput_Id:", NRCellDU_Dl_Throughput_Id)

	NRCellDU_Ul_Throughput_Id := findMetricIdInMeasType(performanceParseToNode, "NRCellDU.UlThroughput")
	fmt.Println("NRCellDU_Ul_Throughput_Id:", NRCellDU_Ul_Throughput_Id)

	NRCell_Ul_Avg_Latency_Id := findMetricIdInMeasType(performanceParseToNode, "NRCell.UlAvgLatency")
	fmt.Println("NRCell_Ul_Avg_Latency_Id:", NRCell_Ul_Avg_Latency_Id)

	NRCell_Dl_Avg_Latency_Id := findMetricIdInMeasType(performanceParseToNode, "NRCell.DlAvgLatency")
	fmt.Println("NRCell_Dl_Avg_Latency_Id:", NRCell_Dl_Avg_Latency_Id)

	NRCellDU_Ul_PrbUtilRate_Id := findMetricIdInMeasType(performanceParseToNode, "NRCellDU.UlPrbUtilRate")
	fmt.Println("NRCellDU_Ul_PrbUtilRate_Id:", NRCellDU_Ul_PrbUtilRate_Id)

	NRCellDU_Dl_PrbUtilRate_Id := findMetricIdInMeasType(performanceParseToNode, "NRCellDU.DlPrbUtilRate")
	fmt.Println("NRCellDU_Dl_PrbUtilRate_Id:", NRCellDU_Dl_PrbUtilRate_Id)

	RRC_Ue_Aver_Number_Id := findMetricIdInMeasType(performanceParseToNode, "RRC.UeAverNumber")
	fmt.Println("RRC_Ue_Aver_Number_Id:", RRC_Ue_Aver_Number_Id)

	RRC_UeLargest_Number_Id := findMetricIdInMeasType(performanceParseToNode, "RRC.UeLargestNumber")
	fmt.Println("RRC_UeLargest_Number_Id:", RRC_UeLargest_Number_Id)

	//UE_Dl_Throughput_Crnti_Ids := xmlquery.Find(performanceParseToNode, "//measType[UE.DlThroughput.Crnti]/@p") //array
	//UE_Ul_Throughput_Crnti_Ids := xmlquery.Find(performanceParseToNode, "//measType[UE.UlThroughput.Crnti]/@p") //array
	//UE_Average_Cqi_Crnti_Ids := xmlquery.Find(performanceParseToNode, "//measType[UE.AverageCqi.Crnti]/@p")     //array

	RRC_ConnSucc_Rate := findValueWithMetricId(performanceParseToNode, RRC_ConnSuccRate_Id)
	RRC_ConnDrop_Rate := findValueWithMetricId(performanceParseToNode, RRC_ConnDropRate_Id)
	NG_Setup_SuccRate := findValueWithMetricId(performanceParseToNode, NGSetup_SuccessRate_Id)
	MM_HoIntra_SuccRate := findValueWithMetricId(performanceParseToNode, MMHoIntra_SuccRate_Id)
	MM_HoInter_SuccRate := findValueWithMetricId(performanceParseToNode, MMHoInter_SuccRate_Id)
	MM_HoResInter_SuccRate := findValueWithMetricId(performanceParseToNode, MM_HoResInter_SuccRate_Id)
	NRCellDU_Dl_Throughput := findValueWithMetricId(performanceParseToNode, NRCellDU_Dl_Throughput_Id)
	NRCellDU_Ul_Throughput := findValueWithMetricId(performanceParseToNode, NRCellDU_Ul_Throughput_Id)
	NRCell_Ul_AvgLatency := findValueWithMetricId(performanceParseToNode, NRCell_Ul_Avg_Latency_Id)
	NRCell_Dl_AvgLatency := findValueWithMetricId(performanceParseToNode, NRCell_Dl_Avg_Latency_Id)
	NRCellDU_Ul_PrbUtilRate := findValueWithMetricId(performanceParseToNode, NRCellDU_Ul_PrbUtilRate_Id)
	NRCellDU_Dl_PrbUtilRate := findValueWithMetricId(performanceParseToNode, NRCellDU_Dl_PrbUtilRate_Id)
	RRC_Ue_AverNumber := findValueWithMetricId(performanceParseToNode, RRC_Ue_Aver_Number_Id)
	RRC_Ue_LargestNumber := findValueWithMetricId(performanceParseToNode, RRC_UeLargest_Number_Id)

	println("Value:")
	println("RRC_Conn_SuccRate", RRC_ConnSucc_Rate)
	println("RRC_Conn_DropRate", RRC_ConnDrop_Rate)
	println("NG_Setup_SuccessRate", NG_Setup_SuccRate)
	println("MM_HoIntra_SuccRate", MM_HoIntra_SuccRate)
	println("MM_HoInter_SuccRate", MM_HoInter_SuccRate)
	println("MM_HoResInter_SuccRate", MM_HoResInter_SuccRate)
	println("NRCellDU_Dl_Throughput", NRCellDU_Dl_Throughput)
	println("NRCellDU_Ul_Throughput", NRCellDU_Ul_Throughput)
	println("NRCell_Ul_AvgLatency", NRCell_Ul_AvgLatency)
	println("NRCell_Dl_AvgLatency", NRCell_Dl_AvgLatency)
	println("RRC_Ue_AverNumber", RRC_Ue_AverNumber)
	println("RRC_Ue_LargestNumber", RRC_Ue_LargestNumber)
	println("NRCellDU_Ul_PrbUtilRate", NRCellDU_Ul_PrbUtilRate)
	println("NRCellDU_Dl_PrbUtilRate", NRCellDU_Dl_PrbUtilRate)
}

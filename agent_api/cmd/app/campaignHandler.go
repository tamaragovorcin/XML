package main

import (
	"AgentApp/pkg/dtos"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)



func (app *application) getCampaignMonitoring(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]

	bestCampaigns := getBestCampainsFromNistagram(userId)

	imagesMarshaled, err := json.Marshal(bestCampaigns)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}

func getBestCampainsFromNistagram(id string) []dtos.CampaignDTO {

	err, campaigns := getCampaignsFromInstagram(id)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err, campaignsXMLFile := getOriginalXmlFiles()

	createNewXmlFile(campaigns,campaignsXMLFile)

	err2 := GeneratePdf("bestCampaigns.pdf",campaigns)
	if err2 != nil {
		panic(err2)
	}
	return campaigns
}

func createNewXmlFile(campaigns []dtos.CampaignDTO, bestCampaigns dtos.BestCampaigns) {

	lengthOfList := len(campaigns)
	if lengthOfList>0 {
		campaign := campaigns[0]
		defineXmlCampaign(bestCampaigns.FirstCampaign, campaign)
		if lengthOfList>1 {
			campaign2 := campaigns[1]
			defineXmlCampaign(bestCampaigns.SecondCampaign, campaign2)
			if lengthOfList>2 {
				campaign3 := campaigns[2]
				defineXmlCampaign(bestCampaigns.ThirdCampaign, campaign3)
			}
		}

	}
	output, err := xml.MarshalIndent(&bestCampaigns, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Kreiramo xml fajl na osnovu xml sadrzaja generisanog putem MarshalIndent metode na osnovu rada
	err2 := ioutil.WriteFile("agentCampaigns.xml", output, 0644)
	if err2 != nil {
		fmt.Println(err)
	}
}

func defineXmlCampaign(bestCampaigns dtos.Campaign, campaign dtos.CampaignDTO) {
	bestCampaigns.Id = campaign.Id
	bestCampaigns.User = campaign.User
	bestCampaigns.TargetGroup = campaign.TargetGroup
	bestCampaigns.Link = campaign.Link
	bestCampaigns.Date = campaign.Date

	bestCampaigns.Time = campaign.Time
	bestCampaigns.Description = campaign.Description
	bestCampaigns.ContentType = campaign.ContentType
	bestCampaigns.AgentId = campaign.AgentId.Hex()

	bestCampaigns.DesiredNumber = strconv.Itoa(campaign.DesiredNumber)
	bestCampaigns.CampaignType = campaign.CampaignType
	bestCampaigns.StartTime = campaign.StartTime
	bestCampaigns.EndTime = campaign.EndTime

	bestCampaigns.NumberOfLikes = strconv.Itoa(campaign.NumberOfLikes)
	bestCampaigns.NumberOfDislikes = strconv.Itoa(campaign.NumberOfDislikes)
	bestCampaigns.NumberOfComments = strconv.Itoa(campaign.NumberOfComments)
	bestCampaigns.Likes = campaign.Likes
	bestCampaigns.Dislikes = campaign.Dislikes
	bestCampaigns.Comments = campaign.Comments

	bestCampaigns.TimesShownTotal = campaign.TimesShownTotal
	bestCampaigns.BestInfluencer = campaign.BestInfluencer
	bestCampaigns.HiredInfluencers = campaign.HiredInfluencers
}

func getOriginalXmlFiles() (error, dtos.BestCampaigns) {
	xmlFile, err := os.Open("campaigns.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var bestCampaigns dtos.BestCampaigns

	xml.Unmarshal(byteValue, &bestCampaigns)
	return err, bestCampaigns
}

func getCampaignsFromInstagram(id string) (error, []dtos.CampaignDTO) {
	resp, err := http.Get("http://localhost:80/api/campaign/bestCampaigns/" + id)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var campaigns []dtos.CampaignDTO
	json.Unmarshal(body, &campaigns)

	return err, campaigns
}
func GeneratePdf(filename string,campaigns []dtos.CampaignDTO) error {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 7, "My best campaigns", "0", 0, "CM", false, 0, "")

	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	/*
	pdf.ImageOptions(
		"avatar.jpg",
		80, 20,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)
*/
	return pdf.OutputFileAndClose(filename)
}
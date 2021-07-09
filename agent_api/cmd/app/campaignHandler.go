package main

import (
	"AgentApp/pkg/dtos"
	"AgentApp/pkg/models"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/beevik/guid"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)



func (app *application) getCampaignMonitoring(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	token := vars["token"]

	err, bestCampaigns := getCampaignsFromInstagram(token)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err, campaignsXMLFile := getOriginalXmlFiles()

	createNewXmlFile(bestCampaigns,campaignsXMLFile)

	err2 := GeneratePdf("cmd/app/files/bestCampaigns.pdf",bestCampaigns)
	if err2 != nil {
		panic(err2)
	}
	_,eer := os.Open("cmd/app/files/bestCampaigns.pdf")
	if eer != nil {
		fmt.Println(err)
		return
	}

	imagesMarshaled, err := json.Marshal(bestCampaigns)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}


func createNewXmlFile(campaigns []dtos.CampaignDTO, bestCampaigns dtos.BestCampaigns) {
	list := models.CampaignStatisticReport{
		XMLName:   xml.Name{},
		FileId:    guid.New().String(),
		Campaigns: nil,
		DateCreating: time.Now(),
	}

	campa := []dtos.Campaign{}

	lengthOfList := len(campaigns)
	if lengthOfList>0 {
		campaign := campaigns[0]
		first := defineXmlCampaign(bestCampaigns.FirstCampaign, campaign)
		bestCampaigns.FirstCampaign =first
		campa = append(campa, first)
		if lengthOfList>1 {
			campaign2 := campaigns[1]
			second := defineXmlCampaign(bestCampaigns.SecondCampaign, campaign2)
			bestCampaigns.SecondCampaign = second
			campa = append(campa, second)
			if lengthOfList>2 {
				campaign3 := campaigns[2]
				third := defineXmlCampaign(bestCampaigns.ThirdCampaign, campaign3)
				bestCampaigns.ThirdCampaign = third
				campa = append(campa, third)
			}
		}

	}

	list.Campaigns = campa
	file, _ := xml.MarshalIndent(list, "", "	")


	_ = CreateDocument(file, list.FileId)



	err2 := ioutil.WriteFile("cmd/app/files/agentCampaigns.xml", file, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}


}

func CreateDocument(xmlFile []byte, fileId string) error {
	client := &http.Client{}
	baseXmlDbUrl = fmt.Sprintf("%s%s:%s/exist/rest", "http://", "xml-db", "8081")
	fmt.Println(fmt.Sprintf("%s/report/%s", baseXmlDbUrl, fileId))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/report/%s", baseXmlDbUrl , fileId + ".xml"), bytes.NewBuffer(xmlFile))
	if err != nil {
		return err
	}
	//req.Header.Add("Content-Length", string(len(xmlFile)))
	req.Header.Add("username", "admin")
	req.Header.Add("password", "")
	req.Header.Add("Content-Type", "application/xml")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		if resp == nil {
			return err
		}

		fmt.Println(resp.StatusCode)
		return errors.New("error while creating document")
	}

	return nil
}

func defineXmlCampaign(bestCampaigns dtos.Campaign, campaign dtos.CampaignDTO) dtos.Campaign{
	bestCampaigns.Id = campaign.Id
	bestCampaigns.User = campaign.User
	bestCampaigns.TargetGroup.Gender = campaign.TargetGroup.Gender
	bestCampaigns.TargetGroup.DateOne = campaign.TargetGroup.DateOne
	bestCampaigns.TargetGroup.DateTwo = campaign.TargetGroup.DateTwo
	bestCampaigns.TargetGroup.Location.Country = campaign.TargetGroup.Location.Country
	bestCampaigns.TargetGroup.Location.Town = campaign.TargetGroup.Location.Town
	bestCampaigns.TargetGroup.Location.Street = campaign.TargetGroup.Location.Street

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
	bestCampaigns.Media = campaign.Media
	return bestCampaigns
}

func getOriginalXmlFiles() (error, dtos.BestCampaigns) {
	xmlFile, err := os.Open("cmd/app/files/campaigns.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var bestCampaigns dtos.BestCampaigns

	xml.Unmarshal(byteValue, &bestCampaigns)
	return err, bestCampaigns
}

func getCampaignsFromInstagram(token string) (error, []dtos.CampaignDTO) {
	resp, err := http.Get("http://localhost:80/api/campaign/bestCampaigns/" + token)
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

	pdf.Ln(12)

	for i, campaign := range campaigns {
		pdf.Cell(190, 7, "***************         TOP "+strconv.Itoa(i+1)+" campaign          ***************")
		pdf.Ln(8)

		pdf.Cell(190, 5, "Campaign type:  " + campaign.CampaignType)
		pdf.Ln(8)

		pdf.Cell(190, 5, "Target group : ")
		pdf.Ln(8)
		pdf.Cell(190, 5, "                Gender: " + campaign.TargetGroup.Gender)
		pdf.Ln(8)
		pdf.Cell(190, 5, "                Date of birth between: " + campaign.TargetGroup.DateOne + "-" + campaign.TargetGroup.DateTwo)
		pdf.Ln(8)


		pdf.Cell(190, 5, "                Location: " +getLocationString(campaign.TargetGroup.Location))
		pdf.Ln(8)

		pdf.Cell(190, 5, "Link : " + campaign.Link)
		pdf.Ln(8)

		pdf.Cell(190, 5, "Description:" + campaign.Description)
		pdf.Ln(8)

		pdf.Cell(190, 5, "Number of likes: " +strconv.Itoa(campaign.NumberOfLikes))
		pdf.Ln(8)

		pdf.Cell(190, 5, "Number of dislikes:" +strconv.Itoa(campaign.NumberOfDislikes))
		pdf.Ln(8)

		pdf.Cell(80, 5, "Number of comments: "  +strconv.Itoa(campaign.NumberOfComments))
		pdf.Ln(8)

		pdf.Cell(190, 5, "Likes: " + campaign.Likes)
		pdf.Ln(8)

		pdf.Cell(190, 5, "Dislikes: " + campaign.Dislikes)
		pdf.Ln(8)

		pdf.Cell(190, 5, "Comments: ")
		pdf.Ln(8)
		allComments :=getCommentsSplitted(campaign.Comments)
		for _, comment := range allComments {
			pdf.Cell(190, 5, "                " + comment)
			pdf.Ln(8)
		}

		allHiredInfluencers := getHiredInfluencers(campaign.HiredInfluencers)
		pdf.Cell(190, 5, "Hired influencers: " )
		pdf.Ln(8)
		for _, influencer := range allHiredInfluencers {
			pdf.Cell(190, 5, "                " + influencer)
			pdf.Ln(8)
		}

		pdf.Cell(190, 5, "Best influencer: " +campaign.BestInfluencer)
		pdf.Ln(8)


		if campaign.CampaignType=="oneTime" {
			pdf.Cell(190, 5, "Date : " +campaign.Date)
			pdf.Ln(8)

			pdf.Cell(190, 5, "Time: " +campaign.Time)
			pdf.Ln(8)


		}else {
			pdf.Cell(190, 5, "Start Time: " +campaign.StartTime)
			pdf.Ln(8)

			pdf.Cell(190, 5, "End time: " +campaign.EndTime)
			pdf.Ln(8)

			pdf.Cell(190, 5, "Total publishments: " +strconv.Itoa(campaign.TimesShownTotal))
			pdf.Ln(8)

			pdf.Cell(190, 5, "Number of publishments in one day: " +strconv.Itoa(campaign.DesiredNumber))
			pdf.Ln(8)
		}
		img,_,err := image.Decode(bytes.NewReader(campaign.Media))
		if err != nil {
			fmt.Println(err.Error())
		}
		out, _ := os.Create("cmd/app/files/"+strconv.Itoa(i)+".jpg")
		defer out.Close()

		err = jpeg.Encode(out,img, nil)

			pdf.ImageOptions(
				"cmd/app/files/"+strconv.Itoa(i)+".jpg",
				70, 5,
				50, 50,
				true,
				gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
				0,
				"",
			)
		pdf.Ln(50)

	}

	return pdf.OutputFileAndClose(filename)
}

func getHiredInfluencers(influencers string) []string {
	splitted := strings.Split(influencers, ",")
	return splitted
}

func getCommentsSplitted(comments string) []string {
	splitted := strings.Split(comments, ",")
	return splitted
}

func getLocationString(location dtos.LocationTarget) string {
	locationString :=""

	if location.Country !="" {
		if location.Town=="" {
			locationString += location.Country
		}else {
			if location.Street=="" {
				locationString += location.Country + ", " + location.Town
			}else {
				locationString += location.Country + ", " + location.Town + ", " + location.Street
			}
		}
	}
	return locationString
}

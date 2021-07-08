package main

import (
	"AgentApp/pkg/models"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type XmlDbClient interface {
	CreateDocument(xmlFile []byte, fileId string) error
	GetAllDocument() (*models.XmlDatabaseResponse, error)
	GetDocumentById(name string) (*models.CampaignStatisticReport, error)
}

type xmlDbClient struct {

}

var (
	baseXmlDbUrl = ""
)

func NewXmlDbClient() XmlDbClient {
	baseXmlDbUrl = fmt.Sprintf("%s%s:%s/exist/rest", "http://", "xml-db", "8081")
	return &xmlDbClient{}
}


func (x xmlDbClient) CreateDocument(xmlFile []byte, fileId string) error {
	client := &http.Client{}

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

func (x xmlDbClient) GetAllDocument() (*models.XmlDatabaseResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/report/", baseXmlDbUrl), nil)
	if err != nil {
		return nil, err
	}
	//req.Header.Add("Content-Length", string(len(xmlFile)))
	req.Header.Add("username", "admin")
	req.Header.Add("password", "")
	//req.Header.Add("Content-Type", "application/xml")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return nil, err
		}

		fmt.Println(resp.StatusCode)
		return nil, errors.New("error while getting documents")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	var statistic *models.XmlDatabaseResponse
	err = xml.Unmarshal(bodyBytes, &statistic)
	if err != nil {
		return nil, err
	}

	return statistic, nil
}

func (x xmlDbClient) GetDocumentById(name string) (*models.CampaignStatisticReport, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/report/%s", baseXmlDbUrl, name), nil)
	if err != nil {
		return nil, err
	}
	//req.Header.Add("Content-Length", string(len(xmlFile)))
	req.Header.Add("username", "admin")
	req.Header.Add("password", "")
	//req.Header.Add("Content-Type", "application/xml")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if resp == nil {
			return nil, err
		}

		fmt.Println(resp.StatusCode)
		return nil, errors.New("error while getting documents")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

	var statistic *models.CampaignStatisticReport
	err = xml.Unmarshal(bodyBytes, &statistic)
	if err != nil {
		return nil, err
	}

	return statistic, nil
}
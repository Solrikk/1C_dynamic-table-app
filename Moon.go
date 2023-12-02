package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Offer struct {
	XMLName    xml.Name `xml:"offer"`
	ID         string   `xml:"id,attr"`
	Name       string   `xml:"name"`
	CategoryID string   `xml:"categoryId"`
	Price      string   `xml:"price"`
	OldPrice   string   `xml:"oldprice"`
	URL        string   `xml:"url"`
	Picture    string   `xml:"picture"`
	Size       Size     `xml:"param[name='Размер']"`
}

type Size struct {
	Value string `xml:",chardata"`
}

type Category struct {
	XMLName  xml.Name `xml:"category"`
	ID       string   `xml:"id,attr"`
	ParentID string   `xml:"parentId,attr"`
	Name     string   `xml:",chardata"`
}

func main() {
	url := "https://seal-parket.ru/yml/yandex.yml?category%5B%5D=20&limit=5000&price_type=piece"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	responseString := string(body)

	offers, categories := extractOffersAndCategories(responseString)

	directory := "C:/Users/Solrikk/Downloads/Doki"
	filename := "offers.csv"
	filepath := filepath.Join(directory, filename)

	err = os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writeOffersToCSV(offers, categories, filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Идентификаторы offers, имена, категории, идентификаторы категорий, цены, старые цены, URL, картинки, размеры успешно записаны в файл.")
}

func extractOffersAndCategories(responseString string) ([]Offer, map[string]string) {
	var offers []Offer
	categories := make(map[string]string)

	type OfferWrapper struct {
		Offers []Offer `xml:"shop>offers>offer"`
	}

	type CategoryWrapper struct {
		Categories []Category `xml:"shop>categories>category"`
	}

	var offerWrapper OfferWrapper
	var categoryWrapper CategoryWrapper
	xml.Unmarshal([]byte(responseString), &offerWrapper)
	xml.Unmarshal([]byte(responseString), &categoryWrapper)

	offers = offerWrapper.Offers

	for _, category := range categoryWrapper.Categories {
		categories[category.ID] = category.Name
	}

	return offers, categories
}

func writeOffersToCSV(offers []Offer, categories map[string]string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("\xEF\xBB\xBF") // Добавляем BOM для указания кодировки UTF-8

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Comma = ';'

	writer.Write([]string{"ID", "Name", "CategoryID", "Category", "Price", "OldPrice", "URL", "Picture", "Размер | мм"})

	for _, offer := range offers {
		category := categories[offer.CategoryID]

		pictures := strings.Split(offer.Picture, "///")
		size := ""

		if offer.Size.Value != "" {
			size = strings.ReplaceAll(offer.Size.Value, "х", "x")
		}

		if offer.OldPrice != "" {
			for _, picture := range pictures {
				writer.Write([]string{offer.ID, offer.Name, offer.CategoryID, category, offer.Price, offer.OldPrice, offer.URL, picture, size})
			}
		} else {
			for _, picture := range pictures {
				writer.Write([]string{offer.ID, offer.Name, offer.CategoryID, category, offer.Price, "", offer.URL, picture, size})
			}
		}
	}

	return nil
}

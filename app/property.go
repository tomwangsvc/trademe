package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Property struct {
	Id            string
	StreetAddress string
	Town          string
	ValuationDate string
	Value         float64
}

type Properties []Property

func (properties Properties) Print() string {
	tmp := struct {
		Length     int        `json:"length,omitempty"`
		Properties Properties `json:"properties,omitempty"`
	}{
		len(properties),
		properties,
	}

	bytes, _ := json.MarshalIndent(tmp, "\t\t", "\t")
	return string(bytes)
}

func (p Property) Valid() bool {
	return p.Id != "" && p.StreetAddress != "" && p.Town != "" && p.ValuationDate != "" && p.Value > 0
}

func InitProperties(filePath string) Properties {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var properties []Property
	var line = 0
	for scanner.Scan() {
		line++
		if line == 1 {
			continue
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		var newProperty Property
		// find id from text
		if index := strings.Index(text, "\t"); index != -1 {
			newProperty.Id = strings.TrimSpace(text[:index])
			text = strings.TrimSpace(text[index+1:])
		}

		// find valuation
		if index := strings.LastIndex(text, "\t"); index != -1 {
			value, err := strconv.ParseFloat(text[index+1:], 64)
			if err != nil {
				LogInfo("failed to parse expected float with error (%v) but it will continue to next row", err)
				continue
			}
			newProperty.Value = value
			text = strings.TrimSpace(text[:index])
		}

		if index := strings.LastIndex(text, "\t"); index != -1 {
			newProperty.ValuationDate = text[index+1:]
			text = strings.TrimSpace(text[:index])
		}

		if index := strings.LastIndex(text, "\t"); index != -1 {
			newProperty.Town = text[index+1:]
			text = strings.TrimSpace(text[:index])
		}

		newProperty.StreetAddress = text

		if !newProperty.Valid() {
			LogInfo("property is not valid %+v and it will be ignored", newProperty)
			continue
		}
		properties = append(properties, newProperty)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return properties
}

func FilterAndUseFirstItem(properties []Property) Properties {
	mapProperties := make(map[string]Property)
	for _, property := range properties {
		key := property.StreetAddress + property.Town + property.ValuationDate
		// we ignore property when it is found in map because we want to use the first one
		if _, ok := mapProperties[key]; !ok {
			mapProperties[key] = property
		}
	}

	var results []Property
	for _, value := range mapProperties {
		results = append(results, value)
	}

	return results
}

func FilterAndUseLastItem(properties []Property) Properties {
	mapProperties := make(map[string]Property)
	for _, property := range properties {
		key := property.StreetAddress + property.Town + property.ValuationDate
		// the last one will overrite the first one if they have the same key
		mapProperties[key] = property
	}

	var results []Property
	for _, value := range mapProperties {
		results = append(results, value)
	}

	return results
}

func FilterOutDuplicatedItem(properties []Property) Properties {
	duplicatedKeys := []string{}
	mapProperties := make(map[string]Property)
	for _, property := range properties {

		key := property.StreetAddress + property.Town + property.ValuationDate
		if _, ok := mapProperties[key]; ok {
			duplicatedKeys = append(duplicatedKeys, key)
			// remove it in map
			delete(mapProperties, key)

			// check if key is duplicated before
		} else if !ContainString(duplicatedKeys, key) {
			mapProperties[key] = property
		}
	}

	var results []Property
	for _, value := range mapProperties {
		results = append(results, value)
	}

	return results
}

// filter out under 400k
// filter out street in black list
// filter out 10th item
func FilterByMutipleConditions(properties []Property) Properties {
	var results []Property
	for index, property := range properties {
		if property.Value >= 400000 && !isStreetInBlackList(property.StreetAddress) && ((index+1)%10 != 0) {
			results = append(results, property)
		}
	}

	return results
}

var blackListWords = []string{"AVE", "CRES", "PL"}

func isStreetInBlackList(address string) bool {
	addr := strings.Split(address, " ")
	if len(addr) < 3 {
		LogInfo("address is wrong formatted with less than 3 words :%s", address)
		return false
	}

	lastAddress := addr[len(addr)-1]
	for _, blackListWord := range blackListWords {
		if lastAddress == blackListWord {
			return true
		}
	}

	return false
}

func ProcessByChunk(properties []Property, chunkSize int) Properties {
	if chunkSize < 1 {
		LogInfo("chunk size %d is less than 1", chunkSize)
		return nil
	}

	properties = FilterOutDuplicatedItem(properties)
	var numOfPropertiesForOneRoutine = len(properties) / chunkSize
	if chunkSize > len(properties) {
		numOfPropertiesForOneRoutine = 1
		chunkSize = len(properties)
	}

	var result []Property
	var wg sync.WaitGroup
	wg.Add(chunkSize)
	for i := 0; i < chunkSize; i++ {
		go func(i int) {
			defer wg.Done()
			beginIndex := i * numOfPropertiesForOneRoutine
			endIndex := (i + 1) * numOfPropertiesForOneRoutine
			if endIndex > len(properties) {
				endIndex = len(properties)
			}
			fmt.Println(beginIndex, endIndex, numOfPropertiesForOneRoutine)
			for property := range worker(properties[beginIndex:endIndex]) {
				result = append(result, property)
			}
		}(i)
	}

	wg.Wait()
	return result
}

func worker(properties []Property) (chanProperties chan Property) {
	chanProperties = make(chan Property)
	go func() {
		for _, property := range FilterByMutipleConditions(properties) {
			chanProperties <- property
		}
		close(chanProperties)
	}()
	return chanProperties
}

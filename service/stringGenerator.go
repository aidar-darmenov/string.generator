package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aidar-darmenov/string.generator/model"
	"github.com/aidar-darmenov/string.generator/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (s *Service) GenerateString(c *gin.Context) {
	numString := c.Param("number")

	numInt, err := strconv.Atoi(numString)
	if err != nil {
		//logger used by team has to be here
		c.JSON(400, fmt.Sprintf("%v", err)+numString)
	}

	client := &http.Client{}

	var encryptorResponseArray []model.EncryptorStringList

	var randomStringArray model.EncryptorStringList

	for i := 0; i < numInt; i++ {
		randomStringArray.StringArray = append(randomStringArray.StringArray, utils.RandStringRunes(10))
	}

	data, errMarshal := json.Marshal(randomStringArray)
	if errMarshal != nil {
		//logger used by team has to be here
		fmt.Println(errMarshal)
	}
	req, err := http.NewRequest("POST", s.Configuration.StringEncryptor, bytes.NewBuffer(data))
	if err != nil {
		//logger used by team has to be here
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		//logger used by team has to be here
		fmt.Println(err)
	}

	defer response.Body.Close()

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		//logger used by team has to be here
		fmt.Println(err)
	}

	var hashedArray model.EncryptorStringList

	err = json.Unmarshal(data, &hashedArray)
	if err != nil {
		//logger used by team has to be here
		fmt.Println("data:", string(data))
		fmt.Println(err)
	}

	encryptorResponseArray = append(encryptorResponseArray, hashedArray)

	c.JSON(200, encryptorResponseArray)
}

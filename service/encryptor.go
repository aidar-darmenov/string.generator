package service

import (
	"github.com/aidar-darmenov/string.generator/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) EncryptStringList(c *gin.Context) {
	var stringArray model.EncryptorStringList

	err := c.Bind(&stringArray)
	if err != nil {
		c.JSON(400, err)
	}

	var counter int = 0
	var stopFiller bool = false

	//Declaring and making newArray where we will assign new hashed values
	var newArray = make([][32]byte, len(stringArray.StringArray))

	//Starting listener for filling newArray
	go func() {
		for {
			if stopFiller {
				break
			}
			select {
			case newStringElem := <-s.ChannelFiller:
				newArray[newStringElem.Index] = newStringElem.Value
				counter++
				if counter == len(stringArray.StringArray) {
					c.JSON(200, newArray)
					stopFiller = true
				}
			}
		}
	}()

	//Looping stringArray and sending elements to channel
	for i := range stringArray.StringArray {
		s.ChannelString <- model.StringElement{
			Value: stringArray.StringArray[i],
			Index: i,
		}
	}
}

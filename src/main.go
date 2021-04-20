package main

import (
	"encoding/base64"
	"strconv"

	sdk "github.com/soracom/orbit-sdk-tinygo"
	"github.com/urawa72/jsonparser"
	"github.com/urawa72/orbit-tinygo/src/data"
)

// application entry point, but the orbit runtime never executes this.
func main() {
}

//export uplink
func uplink() sdk.ErrorCode {
	inputBuffer, err := sdk.GetInputBuffer()
	if err != nil {
		sdk.Log(err.Error())
		return -1
	}
	sdk.Log("Input Buffer: " + string(inputBuffer) + "\n")

	resourceId, err := sdk.GetSourceValue("resourceId")
	if err != nil {
		sdk.Log(err.Error())
		return -1
	}
	sdk.Log("Source Value: " + string(resourceId) + "\n")

	// name, err := sdk.GetTagValue("name")
	// if err != nil {
	// 	sdk.Log(err.Error())
	// 	return -1
	// }
	// sdk.Log("Tag Value: " + string(name) + "\n")

	decodedInputBuffer, err := base64.StdEncoding.DecodeString(string(inputBuffer))
	if err != nil {
		sdk.Log(err.Error())
		return -1
	}
	sdk.Log("Base64 Decoded: " + string(decodedInputBuffer) + "\n")

	output, err := convertInputToOutput(decodedInputBuffer)
	if err != nil {
		sdk.Log(err.Error())
		return - 1
	}

	serializedOuput, err := output.SerializeJSON()
	if err != nil {
		sdk.Log(err.Error())
		return - 1
	}

	sdk.SetOutputJSON(string(serializedOuput))
	sdk.Log("Serialize JSON: " + string(serializedOuput) + "\n")

	return sdk.ErrorCode(0)
}

func convertInputToOutput(input []byte) (*data.Output, error) {
	lat, err := jsonparser.GetFloat(input, "lat")
	if err != nil {
		return nil, err
	}
	slat := strconv.FormatFloat(lat, 'f', -1, 64)

	lon, err := jsonparser.GetFloat(input, "lon")
	if err != nil {
		return nil, err
	}
	slon := strconv.FormatFloat(lon, 'f', -1, 64)

	bat, err := jsonparser.GetInt(input, "bat")
	if err != nil {
		return nil, err
	}

	rs, err := jsonparser.GetInt(input, "rs")
	if err != nil {
		return nil, err
	}

	temp, err := jsonparser.GetFloat(input, "temp")
	if err != nil {
		return nil, err
	}
	stemp := strconv.FormatFloat(temp, 'f', -1, 64)

	humi, err := jsonparser.GetFloat(input, "humi")
	if err != nil {
		return nil, err
	}
	shumi := strconv.FormatFloat(humi, 'f', -1, 64)

	timestamp := sdk.GetTimestamp()

	return &data.Output{
		Lat: 	   	slat,
		Lon: 	   	slon,
		Bat: 	   	bat,
		Rs:	 	   	rs,
		Temp:		stemp,
		Humi:		shumi,
		Timestamp:	timestamp,
	}, nil
}

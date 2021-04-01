package main

import (
	"strconv"
 	"encoding/base64"

	sdk "github.com/soracom/orbit-sdk-tinygo"
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

	timestamp := sdk.GetTimestamp()
	if err != nil {
		sdk.Log(err.Error())
		return -1
	}
	sdk.Log("Timestamp: " + strconv.FormatInt(timestamp, 10) + "\n")

	dec, err := base64.StdEncoding.DecodeString(string(inputBuffer))
	if err != nil {
		sdk.Log(err.Error())
		return -1
	}
	sdk.Log("Base64 decode: " + string(dec) + "\n")

	sdk.SetOutputJSON("{\"message\": \"Hello from Orbit with TinyGo\"}")

	return sdk.ErrorCode(0)
}

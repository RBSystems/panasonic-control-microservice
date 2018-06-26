package handlers

import (
	"net/http"
	"strconv"

	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/panasonic-control-microservice/helpers"
	"github.com/labstack/echo"
)

////////////////////////////////////////
//Power Controls
////////////////////////////////////////

//PowerOn powers on the projector
func PowerOn(context echo.Context) error {
	log.L.Infof("Powering on %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                        //Get the address of the projector
	powerValue := "ON"

	//Send the Command to the helper to turn on power
	err := helpers.SetPower(address, powerValue)
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	//return the result turning on in a JSON format
	return context.JSON(http.StatusOK, se.PowerStatus{"on"})
}

//PowerStandby sets the projector in standby mode
func PowerStandby(context echo.Context) error {
	log.L.Infof("Going on standby %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                             //Get the address of the projector
	powerValue := "OFF"

	//Send the Command to the helper to turn off power
	err := helpers.SetPower(address, powerValue)
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	//return the result of going on standby in a JSON format
	return context.JSON(http.StatusOK, se.PowerStatus{"standby"})
}

//PowerStatus returns the power status, either on or standby
func PowerStatus(context echo.Context) error {
	log.L.Infof("Getting power status of %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                                    //Get the address of the projector

	//Send the Command to the helper to get power status
	status, err := helpers.GetPower(address)
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, status)
}

////////////////////////////////////////
//Volume Controls (Includse Muting)
////////////////////////////////////////

//SetVolume sets the volume of the projector
func SetVolume(context echo.Context) error {
	address := context.Param("address")                      //Get the address of the projector
	volumeLevel, err := strconv.Atoi(context.Param("level")) //Gets the volume level to be set.

	err = helpers.SetVolume(address, uint8(volumeLevel)) //Use SetVolume to change the volume level
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, se.Volume{volumeLevel})
}

//VolumeLevel returns the volume Level, giving the volume level number
func VolumeLevel(context echo.Context) error {
	log.L.Infof("Getting voulme status of %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                                     //Get the address of the projector

	level, err := helpers.GetVolume(address) //Calls the GetVolume fuction that parses the XML
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, level)
}

//Mute mutes the AUDIO of a projector
func Mute(context echo.Context) error {
	log.L.Infof("Muting %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "ON"

	err := helpers.SetAudioMute(address, muteValue) //Set the value to 'ON'
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, se.MuteStatus{true})
}

//UnMute umutes the AUDIO projector
func UnMute(context echo.Context) error {
	log.L.Infof("Unmuting %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "OFF"

	err := helpers.SetAudioMute(address, muteValue) //Sets the value to 'OFF'
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, se.MuteStatus{false})

}

//MuteStatus returns the Mute status, stating if mute is on or off
func MuteStatus(context echo.Context) error {

	address := context.Param("address") //Get the address of the projector

	status, err := helpers.GetMute(address) //Calls the GetMute fuction that parses the XML
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, status)
}

////////////////////////////////////////
//Input Controls
////////////////////////////////////////

//SetInputPort changes the input port
func SetInputPort(context echo.Context) error {
	address := context.Param("address") //Get the address of the projector
	port := context.Param("port")       //Get the input port of the projector

	err := helpers.SetInputPort(address, port) //Use SetInput to change to the selected port
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, se.Input{port})
}

//DisplayBlank sets the display screen to blank
func DisplayBlank(context echo.Context) error {
	log.L.Infof("Displaying blank screen on %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "ON"                   //Value is on to turn on the mute

	err := helpers.SetAVMute(address, muteValue) //Sets the value to 'ON'
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//DisplayUnBlank unblanks the display screen
func DisplayUnBlank(context echo.Context) error {
	log.L.Infof("Unblanking the blank screen on %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "OFF"

	err := helpers.SetAVMute(address, muteValue) //Sets the value to 'OFF'
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//InputStatus returns the Input status, giving the current input port
func InputStatus(context echo.Context) error {
	address := context.Param("address") //Get the address of the projector

	//Send the Command to the helper get the status of the input
	status, err := helpers.GetInput(address)
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, status)
}

//BlankedStatus returns the Input status, giving the current input port
func BlankedStatus(context echo.Context) error {

	address := context.Param("address") //Get the address of the projector

	//Send the Command to the helper to get the blanked status
	status, err := helpers.GetBlankedStatus(address)
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, status)
}

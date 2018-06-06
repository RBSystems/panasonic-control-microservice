package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/panasonic-microservice/helpers"
	"github.com/labstack/echo"
)

////////////////////////////////////////
//Power Controls
////////////////////////////////////////

//Function that powers on the projector
func PowerOn(context echo.Context) error {
	log.L.Info("Powering on %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                       //Get the address of the projector
	powerValue := "ON"

	err := helpers.SetPower(address, powerValue)

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
}

//Function that sets the projector in standby mode
func PowerStandby(context echo.Context) error {
	log.L.Info("Going on standby %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                            //Get the address of the projector
	powerValue := "OFF"

	err := helpers.SetPower(address, powerValue)

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//Function that returns the power status, either on or standby
func PowerStatus(context echo.Context) error {
	log.L.Infof("Getting power status of %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                                    //Get the address of the projector

	err := helpers.GetPower(address)

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
}

////////////////////////////////////////
//Volume Controls
////////////////////////////////////////

//Function that sets the volume of the projector
func SetVolume(context echo.Context) error {
	address := context.Param("address")                      //Get the address of the projector
	volumeLevel, err := strconv.Atoi(context.Param("level")) //Gets the volume level to be set.
	log.L.Info("Setting Volume to %v on %s", volumeLevel, address)

	err = helpers.SetVolume(address, uint8(volumeLevel)) //Use SetVolume to change the volume level

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//Function that changes the input port
func SetInputPort(context echo.Context) error {
	log.L.Info("Switching input for %s to %s...", context.Param("address"), context.Param("port"))
	address := context.Param("address") //Get the address of the projector
	port := context.Param("port")       //Get the input port of the projector

	err := helpers.SetInputPort(address, port) //Use SetInput to change to the selected port

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//Function that mutes the AUDIO of a projector
func Mute(context echo.Context) error {
	log.L.Info("Muting %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "ON"

	err := helpers.SetAudioMute(address, muteValue) //Set the value to 'ON'

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//Function that umutes the AUDIO projector
func UnMute(context echo.Context) error {
	log.L.Info("Unmuting %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "OFF"

	err := helpers.SetAudioMute(address, muteValue) //Sets the value to 'OFF'

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)

}

//Function that sets the display screen to blank
func DisplayBlank(context echo.Context) error {
	log.L.Info("Displaying blank screen on %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "OFF"

	err := helpers.SetAVMute(address, muteValue) //Sets the value to 'OFF'

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//Function that sets the display screen to blank
func DisplayUnBlank(context echo.Context) error {
	log.L.Info("Unblanking the blank screen on %s...", context.Param("address"))
	address := context.Param("address") //Get the address of the projector
	muteValue := "OFF"

	err := helpers.SetAVMute(address, muteValue) //Sets the value to 'OFF'

	if err != nil {
		log.L.Info("Error: %v", err.Error())                             //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

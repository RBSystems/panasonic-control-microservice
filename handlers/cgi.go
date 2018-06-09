package handlers

import (
	"net/http"
	"strconv"

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

	err := helpers.SetPower(address, powerValue)

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
}

//PowerStandby sets the projector in standby mode
func PowerStandby(context echo.Context) error {
	log.L.Info("Going on standby %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                            //Get the address of the projector
	powerValue := "OFF"

	err := helpers.SetPower(address, powerValue)

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//PowerStatus returns the power status, either on or standby
func PowerStatus(context echo.Context) error {
	log.L.Infof("Getting power status of %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                                    //Get the address of the projector

	err := helpers.GetPower(address)

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
}

////////////////////////////////////////
//Volume Controls (Includse Muting)
////////////////////////////////////////

//SetVolume sets the volume of the projector
func SetVolume(context echo.Context) error {
	address := context.Param("address")                      //Get the address of the projector
	volumeLevel, err := strconv.Atoi(context.Param("level")) //Gets the volume level to be set.
	log.L.Infof("Setting Volume to %v on %s", volumeLevel, address)

	err = helpers.SetVolume(address, uint8(volumeLevel)) //Use SetVolume to change the volume level

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
}

//VolumeLevel returns the volume Level, giving the volume level number
func VolumeLevel(context echo.Context) error {
	log.L.Infof("Getting voulme status of %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                                     //Get the address of the projector

	err := helpers.GetVolume(address) //Calls the GetVolume fuction that parses the XML

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
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
	return context.JSON(http.StatusOK, err)
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
	return context.JSON(http.StatusOK, err)

}

//MuteStatus returns the Mute status, stating if mute is on or off
func MuteStatus(context echo.Context) error {
	log.L.Infof("Getting mute status of %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")                                   //Get the address of the projector

	err := helpers.GetMute(address) //Calls the GetMute fuction that parses the XML

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
}

////////////////////////////////////////
//Input Controls
////////////////////////////////////////

//SetInputPort changes the input port
func SetInputPort(context echo.Context) error {
	log.L.Infof("Switching input for %s to %s...", context.Param("address"), context.Param("port"))
	address := context.Param("address") //Get the address of the projector
	port := context.Param("port")       //Get the input port of the projector

	err := helpers.SetInputPort(address, port) //Use SetInput to change to the selected port

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, err)
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
	log.L.Infof("Getting Input status of %s...", context.Param("address")) //Print the messafe of getting status
	address := context.Param("address")                                    //Get the address of the projector

	err := helpers.GetInput(address)

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
}

//InputStatus returns the Input status, giving the current input port
func BlankedStatus(context echo.Context) error {
	log.L.Infof("Getting blanked status of %s...", context.Param("address")) //Print the messafe of getting status
	address := context.Param("address")                                      //Get the address of the projector

	err := helpers.GetBlankedStatus(address)

	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return nil
}

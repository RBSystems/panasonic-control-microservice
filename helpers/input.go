//Used for switching the inputs of the projector

package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bobziuchkovski/digest"
	"github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
)

//PanasonicInputResponse is a struct for parsing the XML responses
type PanasonicInputResponse struct {
	Result xml.Name `xml:"RESULT"`
	Input  string   `xml:"INPUT"`
	AVMute string   `xml:"AVMUTE"`
}

//SetInputPort sends the CGI command to switch inputs
func SetInputPort(address string, inputValue string) error {
	log.L.Infof("Switching input for %s to %s...", address, inputValue)

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=INPUT&value=%s", address, inputValue) //CGI command based on the documentation provided
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return err
	}
	_, err = t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return err
	}
	return nil
}

//SetAVMute mutes Audio and Video(AV), thus blanking the screen
func SetAVMute(address string, muteValue string) error {

	command := fmt.Sprintf("http://%s/cgi-bin/controlCmd.cgi?param=AVMUTE&value=%s", address, muteValue) //CGI command based on the documentation provided
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return err
	}
	_, err = t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return err
	}
	return nil
}

//GetInput returns the current input of the projector
func GetInput(address string) (statusevaluators.Input, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=INPUT", address) //CGI command based on the documentation provided
	log.L.Infof("Getting Input status of %s...", address)                         //Print the messafe of getting status

	//This is for the digest auth for the projector
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Debugf("Nope Didn't work! - %v", err.Error())
		return statusevaluators.Input{}, err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return statusevaluators.Input{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return statusevaluators.Input{}, err
	}
	log.L.Debugf("%s", b)
	var response PanasonicInputResponse
	err = xml.Unmarshal(b, &response) //Unmarshal the XML
	if err != nil {
		log.L.Errorf("Error:", err)
		return statusevaluators.Input{}, err
	}

	log.L.Infof("Current Input Port: %s", response.Input) //Print out the input, whatever it be

	status := statusevaluators.Input{
		Input: string(response.Input),
	}
	return status, nil
}

//GetBlankedStatus returns the blanked status of the projector
func GetBlankedStatus(address string) (statusevaluators.BlankedStatus, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=INPUT", address) //CGI command based on the documentation provided
	log.L.Infof("Getting blanked status of %s...", address)                       //Print the messafe of getting status

	//This is for the digest auth for the projector
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return statusevaluators.BlankedStatus{}, err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return statusevaluators.BlankedStatus{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return statusevaluators.BlankedStatus{}, err
	}
	log.L.Debugf("%s", b)
	var response PanasonicInputResponse
	err = xml.Unmarshal(b, &response) //Unmarshal the XML
	if err != nil {
		log.L.Errorf("Error:", err)
		return statusevaluators.BlankedStatus{}, err
	}

	log.L.Infof("Current Blanked Stauts: %s", response.AVMute) //Print out the input, whatever it be on or off

	var status statusevaluators.BlankedStatus

	if response.AVMute == "ON" {
		status.Blanked = true
	} else if response.AVMute == "OFF" {
		status.Blanked = false
	}

	return status, nil
}

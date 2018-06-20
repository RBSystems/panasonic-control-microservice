//Used for switching the inputs of the projector

package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bobziuchkovski/digest"
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
func GetInput(address string) (string, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=INPUT", address) //CGI command based on the documentation provided

	//This is for the digest auth for the projector
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Debugf("Nope Didn't work! - %v", err.Error())
		return "", err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return "", err
	}
	log.L.Debugf("%s", b)
	var status PanasonicInputResponse
	err = xml.Unmarshal(b, &status) //Unmarshal the XML
	if err != nil {
		log.L.Errorf("Error:", err)
		return "", err
	}

	log.L.Infof("Current Input Port: %s", status.Input) //Print out the input, whatever it be
	return status.Input, nil
}

//GetBlankedStatus returns the blanked status of the projector
func GetBlankedStatus(address string) (string, error) {
	command := fmt.Sprintf("http://%s/cgi-bin/queryCmd.cgi?param=INPUT", address) //CGI command based on the documentation provided

	//This is for the digest auth for the projector
	t := digest.NewTransport("byuav", "test")
	req, err := http.NewRequest("GET", command, nil)
	if err != nil {
		log.L.Errorf("Nope Didn't work! - %v", err.Error())
		return "", err
	}
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.L.Errorf("Nope still didn't work! - %v", err.Error())
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.L.Errorf("Error retreiving Body. Error:", err)
		return "", err
	}
	log.L.Debugf("%s", b)
	var status PanasonicInputResponse
	err = xml.Unmarshal(b, &status) //Unmarshal the XML
	if err != nil {
		log.L.Errorf("Error:", err)
		return "", err
	}

	log.L.Infof("Current Blanked Stauts: %s", status.AVMute) //Print out the input, whatever it be on or off
	return status.AVMute, nil
}

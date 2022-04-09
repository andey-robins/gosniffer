package sniffer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Sniffer is a struct that logically groups all of the information about a single sniffer agent
type Sniffer struct {
	uid      string
	homeAddr string
}

// NewSniffer performs any initialization and returns a reference to a valid sniffer
func NewSniffer() *Sniffer {
	// TODO: Do any initialization for the sniffer right here
	return &Sniffer{}
}

// RegisterHomeServer associates a server address with the sniffer. If this has been called, reporting
// will be sent to the server instead of to STDOUT
func (s *Sniffer) RegisterHomeServer(addr string) {
	s.homeAddr = addr
}

// StartSniffing will launch the sniffer into it's working state. At this point, the only way to stop
// the sniffer is to kill the process
func (s *Sniffer) StartSniffing() {
	// TODO: Maybe rewrite the sniffer to send back data over channels?
	// TODO: Maybe hang here and provide an interupt to shut everything down gracefully?
	go s.sniff()
}

// Report will properly log and send out the data to the logical output. If a homeAddr is supplied,
// it will post to the api endpoint at that server. Otherwise, it logs out to STDOUT
func (s *Sniffer) Report(data map[string]string) {
	// TODO: Determine if data should be sent as json or something else
	// TODO: Prepare data for POST
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	// Phone data home if there is a registered home server
	if s.homeAddr != "" {
		_, err = http.Post(fmt.Sprintf("%s/register", s.homeAddr), "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println(err)
			return
		}
	}

	// Locally log the data if the sniffer is running in "detached" mode
	log.Println(data)

}

// sniff is the worker for sniffing. logic to run the sniffer will be in here. this internally calls
// the report method. The StartSniffing method is publicly exposed and so will handle verifying the
// state of the sniffer is set up properly to run. This function assumes that s is a correctly setup
// sniffer and just goes for it
func (s *Sniffer) sniff() {
	// TODO: Implement sniffer listening here
	data := make(map[string]string)
	for {
		data["test"] = "true"
		if false {
			s.Report(data)
		}
	}
}

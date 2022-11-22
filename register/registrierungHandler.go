package registrierung

import (
	"log"
	"net/http"
	"strconv"
)

type RegistrierungsHandler struct{}

func (rh *RegistrierungsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Printf("Could not parse form because of %v", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	registrierung := &Registrierung{}
	registrierung.Firstname = req.Form.Get("Firstname")

	b, err := strconv.ParseBool(req.Form.Get("DatenschutzAkzeptiert"))
	if err != nil {
		log.Printf("Could not parse value for DatenschutzAkzeptiert becauseof %v", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	registrierung.DatenschutzAkzeptiert = b
	log.Printf("new registration %v", registrierung)
	rw.WriteHeader(http.StatusCreated)
}

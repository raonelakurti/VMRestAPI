package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
	"os"
	"github.com/google/uuid"
)

func (conf *Config) UpdateVm(w http.ResponseWriter, r *http.Request) {
	var vm Vm
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &vm)
	if err != nil {
		fmt.Println("Error when unmarshalling json", err)
		writeError(w, http.StatusInternalServerError, "Error unMarshalling payload json")
		return
	}

	vms, err := readDataFile();

	if err != nil {
		fmt.Println(err)
		writeError(w, http.StatusBadRequest, "Error occured opening file")
		return
	}
    for i := 0; i <= len(vms); i++ {
		if vms[i].ID == vm.ID {
			vms[i] = vm
			break
		}
	}
	err = updateDataFile(vms)
	if err != nil {
		fmt.Println("Error when deleting given vm", err)
		writeError(w, http.StatusInternalServerError, "Error when deleting given vm")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(vms)

}

func (conf *Config) DeleteVm(w http.ResponseWriter, r *http.Request){
	var vm Vm
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &vm)
	if err != nil {
		fmt.Println("Error when unmarshalling json", err)
		writeError(w, http.StatusInternalServerError, "Error unMarshalling payload json")
		return
	}

	vms, err := readDataFile();

	if err != nil {
		fmt.Println(err)
		writeError(w, http.StatusBadRequest, "Error occured opening file")
		return
	}

	var index int
    for i := 0; i <= len(vms); i++ {
		if vms[i].ID == vm.ID {
			index = i
        	fmt.Println(i)
			break
		}
	}
	vms = append(vms[:index], vms[index+1:]...)

	err = updateDataFile(vms)
	if err != nil {
		fmt.Println("Error when deleting given vm", err)
		writeError(w, http.StatusInternalServerError, "Error when deleting given vm")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
}

func (conf *Config) CreateVm(w http.ResponseWriter, r *http.Request){

	var vm Vm
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &vm)


	if err != nil {
		fmt.Println("Error when unmarshalling json", err)
		writeError(w, http.StatusInternalServerError, "Error unMarshalling payload json")
		return
	}
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
    uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
	vm.ID = uuid

	vms, err := readDataFile();

	if(err != nil) {
		fmt.Println(err)
		writeError(w, http.StatusBadRequest, "Error occured opening file")
		return
	}
	vms = append(vms, vm)

	err = updateDataFile(vms)

	if(err != nil){
		fmt.Println("Error when adding new vm", err)
		writeError(w, http.StatusInternalServerError, "Error when adding new vm")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(vms)
}

func updateDataFile(vms []Vm) (error){
	byteValue, err := json.Marshal(vms)
	if err != nil {
		fmt.Println("Error when marshalling data", err)
		return err
	}

	// Write back to file
	err = ioutil.WriteFile("data.json", byteValue, 0644)

	if(err != nil){
		fmt.Println("Error when updating data file", err)
		return err
	}
	return nil
}

func (conf *Config) GetVmsList(w http.ResponseWriter, r *http.Request) {

	vms, err := readDataFile();

	if(err != nil) {
		fmt.Println(err)
		writeError(w, http.StatusBadRequest, "Error occured opening file")
		return
	}

	//if you have one perticular arg send that one as resp
	keys, ok := r.URL.Query()["vm"]
    
    if !ok || len(keys[0]) < 1 {
     
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(vms)
    }else {
		// Query()["key"] will return an array of items, 
		// we only want the single item.
		vmid := keys[0]

		fmt.Println("Url Param 'key' is: " + string(vmid))
		for i := 0; i <= len(vms); i++ {
			if vms[i].ID == string(vmid) {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.Header().Set("X-Content-Type-Options", "nosniff")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(vms[i])
				return
			}
		}

	}
}

func readDataFile() ([]Vm,error ){
	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	fmt.Println("Successfully Opened data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var vms []Vm

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'vms' which we defined above
	json.Unmarshal(byteValue, &vms)
	return vms,nil
}

func writeError(w http.ResponseWriter, errorCode int, errorMessage string) {
	err := APIError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(err)
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}

// User struct which contains a name
// a type and a list of social links
type Vm struct {
	ID  string `json:"id"`
    AdminUsername   string `json:"adminUsername"`
    Password   string `json:"password"`
    VmName    string    `json:"vmName"`
	VmSize string `json:"vmSize"`
    Region string `json:"region"`
    OsImage string `json:"osImage"`
}

type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

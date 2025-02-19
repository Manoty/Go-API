package main

import(
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		GetClientProfile(w, r)
	case http.MethodPatch:
		UpdateClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)		
	}
}

func GetClientProfile(w http.ResponseWriter, r *http.Request){
	clientProfile := r.Context().Value("clientProfile").(ClientProfile)
	response := ClientProfile{
		Email: clientProfile.Email,
		Name: clientProfile.Name,
		Id: clientProfile.Id,
	}
	json.NewEncoder(w).Encode(response)

}
func UpdateClientProfile(w http.ResponseWriter, r *http.Request){
	clientProfile := r.Context().Value("clientProfile").(ClientProfile)


	//Decode the json paylaod directly into the struct
	var payLoadData ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payLoadData); err != nil{
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// update profile
	clientProfile.Email = payLoadData.Email
	clientProfile.Name  = payLoadData.Name
	clientProfile.Id    = payLoadData.Id

	w.WriteHeader(http.StatusOK)
}
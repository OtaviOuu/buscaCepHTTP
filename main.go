package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type CepInfo struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", getCepData)
	http.ListenAndServe(":8080", nil)

}

func getCepData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Applicaion/json")

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cepData, error := SearchCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cepData)

}

func SearchCep(cep string) (*CepInfo, error) {
	req, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer req.Body.Close()

	resp, error := io.ReadAll(req.Body)
	if error != nil {
		return nil, error
	}

	var cepData CepInfo
	error = json.Unmarshal(resp, &cepData)
	if error != nil {
		return nil, error
	}
	return &cepData, nil
}

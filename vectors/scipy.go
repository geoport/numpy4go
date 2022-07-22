package vectors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

func sendRequest(app string, body any) map[string]interface{} {
	url := fmt.Sprintf("http://116.203.210.177:8005/%s", app)
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, _ := http.NewRequest("POST", url, payloadBuf)
	client := &http.Client{}
	res, _ := client.Do(req)
	jsonDataFromHttp, _ := ioutil.ReadAll(res.Body)

	return map[string]interface{}{
		"status":  res.Status,
		"respond": jsonDataFromHttp,
	}
}

// Cumtrapz cumulatively integrates f using the composite trapezoidal rule.
func Cumtrapz(f []float64, dx float64, initial float64) []float64 {

	integral := []float64{0}
	for i := 0; i < len(f)-1; i++ {
		n := 0.5 * dx * (f[i+1] + f[i])
		integral = append(integral, integral[i]+n)
	}
	integral[0] = initial
	return integral
}

func Interp1D(x, y, xi []float64) []float64 {
	switch {
	case len(x) != len(y):
		panic("interp1d: slice length mismatch")
	case len(x) < 2:
		panic("interp1d: input data too small")
	case !sort.Float64sAreSorted(x):
		panic("interp1d: input must be sorted")
	}
	type requestData struct {
		X  []float64 `json:"x"`
		Xi []float64 `json:"xi"`
		Y  []float64 `json:"y"`
	}
	type responseData struct {
		Y []float64 `json:"y"`
	}
	var responseBody responseData
	requestBody := &requestData{
		X:  x,
		Xi: xi,
		Y:  y,
	}
	response := sendRequest("interpolate", requestBody)
	if response["status"] == "200 OK" {
		json.Unmarshal(response["respond"].([]byte), &responseBody)
	}
	return responseBody.Y
}

func CurveFit(x, y, p0 []float64, bounds [][]float64) ([]float64, [][]float64) {
	switch {
	case len(x) != len(y):
		panic("interp1d: slice length mismatch")
	case len(x) < 2:
		panic("interp1d: input data too small")
	}
	type requestData struct {
		X      []float64   `json:"x"`
		Y      []float64   `json:"y"`
		P0     []float64   `json:"p0"`
		Bounds [][]float64 `json:"bounds"`
	}
	type responseData struct {
		Popt  []float64   `json:"popt"`
		Pconv [][]float64 `json:"pconv"`
	}
	var responseBody responseData
	requestBody := &requestData{
		X:      x,
		Y:      y,
		P0:     p0,
		Bounds: bounds,
	}
	response := sendRequest("curve_fit", requestBody)
	if response["status"] == "200 OK" {
		json.Unmarshal(response["respond"].([]byte), &responseBody)
	}
	return responseBody.Popt, responseBody.Pconv
}

func FSolve(x0 []float64) ([]float64, float64) {
	type requestData struct {
		X0 []float64 `json:"x0"`
	}
	type responseData struct {
		X   []float64 `json:"X"`
		Ier float64   `json:"ier"`
	}
	var responseBody responseData
	requestBody := &requestData{
		X0: x0,
	}
	response := sendRequest("fsolve", requestBody)
	if response["status"] == "200 OK" {
		json.Unmarshal(response["respond"].([]byte), &responseBody)
	}
	return responseBody.X, responseBody.Ier
}

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

// Cumtrapz cumulatively integrates f(x) using the composite trapezoidal rule.
func Cumtrapz(x, f []float64) []float64 {

	switch {
	case len(x) != len(f):
		panic("integrate: slice length mismatch")
	case len(x) < 2:
		panic("integrate: input data too small")
	case !sort.Float64sAreSorted(x):
		panic("integrate: input must be sorted")
	}
	integral := []float64{0}
	for i := 0; i < len(x)-1; i++ {
		n := 0.5 * (x[i+1] - x[i]) * (f[i+1] + f[i])
		integral = append(integral, integral[i]+n)
	}

	return integral[1:]
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

func Butterworth(Wn []float64, filterOrder int, filterType string) ([]float64, []float64) {
	type requestData struct {
		Wn          []float64 `json:"Wn"`
		FilterOrder int       `json:"N"`
		FilterType  string    `json:"btype"`
	}
	type responseData struct {
		B []float64 `json:"b"`
		A []float64 `json:"a"`
	}
	var responseBody responseData
	requestBody := &requestData{
		Wn:          Wn,
		FilterOrder: filterOrder,
		FilterType:  filterType,
	}
	response := sendRequest("butter", requestBody)
	if response["status"] == "200 OK" {
		json.Unmarshal(response["respond"].([]byte), &responseBody)
	}
	return responseBody.B, responseBody.A

}

func Cheby1(Wn []float64, filterOrder int, filterType string) ([]float64, []float64) {
	type requestData struct {
		Wn          []float64 `json:"Wn"`
		FilterOrder int       `json:"N"`
		FilterType  string    `json:"btype"`
	}
	type responseData struct {
		B []float64 `json:"b"`
		A []float64 `json:"a"`
	}
	var responseBody responseData
	requestBody := &requestData{
		Wn:          Wn,
		FilterOrder: filterOrder,
		FilterType:  filterType,
	}
	response := sendRequest("cheby1", requestBody)
	if response["status"] == "200 OK" {
		json.Unmarshal(response["respond"].([]byte), &responseBody)
	}
	return responseBody.B, responseBody.A

}

//Bessel filtering
func Bessel(Wn []float64, filterOrder int, filterType string) ([]float64, []float64) {
	type requestData struct {
		Wn          []float64 `json:"Wn"`
		FilterOrder int       `json:"N"`
		FilterType  string    `json:"btype"`
	}
	type responseData struct {
		B []float64 `json:"b"`
		A []float64 `json:"a"`
	}
	var responseBody responseData
	requestBody := &requestData{
		Wn:          Wn,
		FilterOrder: filterOrder,
		FilterType:  filterType,
	}
	response := sendRequest("bessel", requestBody)
	if response["status"] == "200 OK" {
		json.Unmarshal(response["respond"].([]byte), &responseBody)
	}
	return responseBody.B, responseBody.A

}

func Filter(b []float64, a []float64, data []float64) []float64 {
	type requestData struct {
		B []float64 `json:"b"`
		A []float64 `json:"a"`
		X []float64 `json:"x"`
	}
	type responseData struct {
		Y []float64 `json:"y"`
	}
	var responseBody responseData
	requestBody := &requestData{
		B: b,
		A: a,
		X: data,
	}
	response := sendRequest("lfilter", requestBody)
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

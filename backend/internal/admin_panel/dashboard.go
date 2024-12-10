package adminpanel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dashboard struct {
	Meta struct {
		Slug string `json:"slug"`
	} `json:"meta"`
	Dashboard json.RawMessage `json:"dashboard"`
}

func main() {
	apiKey := "YOUR_API_KEY"
	dashboardUID := "YOUR_DASHBOARD_UID"
	url := fmt.Sprintf("http://localhost:3000/api/dashboards/uid/%s", dashboardUID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var dashboard Dashboard
	err = json.Unmarshal(body, &dashboard)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Dashboard Slug:", dashboard.Meta.Slug)
	fmt.Println("Dashboard JSON:", string(dashboard.Dashboard))
}

package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// GETCmd represents the GET command
var GETCmd = &cobra.Command{
	Use:   "GET",
	Short: "It recieves all the albums. (You can use also flags to  recieve specific albums)",
	Long: `It recieves albums through flags or all albums without flags. 
	To get album through ID enter "GET --id='2'". To get album through title enter "GET --title='Jeru'".`,
	RunE: getCommand,
}

func getCommand(cmd *cobra.Command, args []string) error {
	/***************************************/

	ID, _ := cmd.Flags().GetString("id")
	Title, _ := cmd.Flags().GetString("title")
	IP, _ := cmd.Flags().GetString("ip")
	All, _ := cmd.Flags().GetString("all")

	if len(args) != 0 {
		println("<!-----------------------------------!>")
		println("Please enter only the requiered fields")
		println("<!-----------------------------------!>")
		println()
	}

	ip := "localhost"

	url := "http://" + ip + ":8070/albums/"

	if IP != "" {
		ip = IP
		url = "http://" + ip + ":8070/albums/"
	}

	if ID != "" {
		return getID(url, ID)

	} else if Title != "" {
		getTitle(url, Title)
	} else if All != "" {
		get(url)
	}
	return nil

}

func get(url string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("user-agent", "golang client")
	req.Header.Set("accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response body")
	}
	println(string(body))

	return nil
}

func init() {
	rootCmd.AddCommand(GETCmd)

	GETCmd.PersistentFlags().String("id", "", "Search for ids of albums. ID must be an integer. If it isn't you get error message")
	GETCmd.PersistentFlags().String("title", "", "Search for title of albums.")
	GETCmd.PersistentFlags().String("ip", "", "Changes the ip, which the host is located")
	GETCmd.PersistentFlags().String("all", "all", "Gets all Albums")
}

func getID(url, id string) error {
	return get(url + id)
}

func getTitle(url, title string) error {
	return get(url + "titles/" + title)
}

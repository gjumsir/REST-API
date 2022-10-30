package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example/web-service-gin/model"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// POSTCmd represents the POST command
var POSTCmd = &cobra.Command{
	Use:   "POST",
	Short: "Post an album",
	Long:  `Posts an album`,
	RunE:  getCommand1,
}

func getCommand1(cmd *cobra.Command, args []string) error {

	IP, _ := cmd.Flags().GetString("ip")

	ip := "localhost"

	url := "http://" + ip + ":8070/albums/"

	if IP != "" {
		ip = IP
		url = "http://" + ip + ":8070/albums/"
	}

	if len(args) == 3 {
		postAlbum(url, args[0], args[1], args[2])
	} else {
		println("Please enter 'title' 'artist' 'price' of the album that you want to post")
	}

	return nil
}

func init() {
	rootCmd.AddCommand(POSTCmd)
	POSTCmd.PersistentFlags().String("ip", "", "Changes the ip, which the host is located")
}

func postAlbum(url, title, artist, price string) error {
	fl, err := strconv.ParseFloat(price, 64)
	post(url, title, artist, fl)
	if err != nil {
		return errors.Wrapf(err, "could not parse integer from %s", fl)
	}
	return nil
}

func post(url, title, artists string, price float64) error {
	jsonData := model.Album{
		ID:     "",
		Title:  title,
		Artist: artists,
		Price:  price,
	}

	raw, err := json.Marshal(jsonData)
	if err != nil {
		return errors.Wrap(err, "could not encode to byte")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(raw))
	if err != nil {
		return errors.Wrap(err, "could not complete request")
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "could not complete request")
	}

	defer response.Body.Close()
	return get(fmt.Sprintf(url + "last"))
}

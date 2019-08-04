package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Middleware/app"
	"net/http"
	"net/http/httptest"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "booker",
	Short: "tool used to manage movie theater",
	Long: `tool used to configure SAP CP landscapes`,
}


func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("first_name", "f", ".", "Define your name")
	registerCmd.Flags().StringP("last_name", "l", ".", "Set your last name")
	registerCmd.Flags().StringP("email", "e", ".", "Set your email address")
	registerCmd.Flags().StringP("birthday", "b", ".", "Set your birthday")




}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	RunE: func(cmd *cobra.Command, args []string) error {

		firstName, err := cmd.Flags().GetString("first_name")
		if err != nil {
			return err
		}

		lastName, err := cmd.Flags().GetString("last_name")
		if err != nil {
			return err
		}

		birthday, err := cmd.Flags().GetString("birthday")
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}

		err = CreateUser(firstName,lastName,birthday,email)


		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		return nil
	},
}


func CreateUser(first_name, last_name, birthday, email string) error {

	testRouter := Middleware.SetupRouter()

	m := app.User{FirstName:first_name,  LastName:last_name, Birthday:"10-16-2020", Email: email}
	b, err := json.Marshal(m)



	//jsonString :="{\"first_name\": \"83\", \"last_name\": \"100\", \"birthday\": \"10-16-2020\", \"email\": \"500\"}"

	req, err := http.NewRequest("POST", "/api/v1/create_user", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Post hearteat failed with error %d.", err)
		return err
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		fmt.Println("/api/v1/instructions failed with error code %d and response", resp.Code, resp.Body)
	}else{
		fmt.Println(resp.Body)
	}
	return nil
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
}


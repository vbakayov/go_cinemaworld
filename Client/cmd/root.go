package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Client/functions"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "booker",
	Short: "tool used to manage movie theater",
	Long: `tool used to configure SAP CP landscapes`,
}


func init() {
	//add register command
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("first_name", "f", ".", "Define your name")
	registerCmd.Flags().StringP("last_name", "l", ".", "Set your last name")
	registerCmd.Flags().StringP("email", "e", ".", "Set your email address")
	registerCmd.Flags().StringP("birthday", "b", ".", "Set your birthday")

	//list the available movies
	RootCmd.AddCommand(listMoviesCmd)

	//register new movie theater
	RootCmd.AddCommand(registerTheater)
	registerTheater.Flags().StringP("name", "n", ".", "Set the name of the theater")
	registerTheater.Flags().StringP("rows", "r", ".", "Set how many rows does it have")
	registerTheater.Flags().StringP("floor", "f", ".", "Set on which floor it is")
}


var registerTheater = &cobra.Command{
	Use:  "register_theater",
	Short:"add new theater to the cinema",
	Long: "add new theater to the cinema",
	RunE: func(cmd *cobra.Command, args []string) error {


		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		rows, err := cmd.Flags().GetString("rows")
		if err != nil {
			return err
		}

		floor, err := cmd.Flags().GetString("floor")
		if err != nil {
			return err
		}


		err = functions.AddNewTheater(name,rows,floor)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		return nil
	},


}


var listMoviesCmd = &cobra.Command{
	Use:  "list",
	Short:"list movies in the cinema",
	Long: "list all screening movies in the cinema",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := functions.GetAvailableMovies()

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		return nil

	},


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

		err = functions.CreateUser(firstName,lastName,birthday,email)


		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		return nil
	},
}


func Execute() {

	functions.InitRouter()
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
}


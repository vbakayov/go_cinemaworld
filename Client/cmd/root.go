package cmd

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.infra.hana.ondemand.com/cloudfoundry/go_cinemaworld/Client/functions"
	"os"
	"regexp"
	"strings"
)

var RootCmd = &cobra.Command{
	Use:   "booker",
	Short: "tool used to manage movie theater",
	Long: `tool used to configure SAP CP landscapes`,
}


func init() {
	//add register command
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("first_name", "f", "", "Define your name")
	registerCmd.Flags().StringP("last_name", "l", "", "Set your last name")
	registerCmd.Flags().StringP("email", "e", "", "Set your email address")
	registerCmd.Flags().StringP("birthday", "b", "", "Set your birthday")

	//list the available movies
	RootCmd.AddCommand(listMoviesCmd)

	//register new movie theater
	RootCmd.AddCommand(registerTheater)
	registerTheater.Flags().StringP("name", "n", "", "Set the name of the theater")
	registerTheater.Flags().StringP("rows", "r", "", "Set how many rows does it have")
	registerTheater.Flags().StringP("floor", "f", "", "Set on which floor it is")

	//register new movie for theater
	RootCmd.AddCommand(registerMovie)
	registerMovie.Flags().StringP("movie_title", "t", "", "Set the title of the movie")
	registerMovie.Flags().StringP("movie_year", "y", "", "Set the year released")
	registerMovie.Flags().StringP("pg_type", "p", "", "Set the pg type classification")
	registerMovie.Flags().StringP("runtime", "r", "", "Set the runtime duration")

}

var registerMovie = &cobra.Command{
	Use:  "register_movie",
	Short:"add new movie to the cinema",
	Long: "add new movie to the cinema",
	RunE: func(cmd *cobra.Command, args []string) error {

		name, err := cmd.Flags().GetString("movie_title")
		if err != nil {
			return err
		}

		movieYear, err := cmd.Flags().GetString("movie_year")
		if movieYear == "" {
			fmt.Println("movie year was not set... Skipping")
		}

		pgType, _ := cmd.Flags().GetString("pg_type")
		if pgType == "" {
			fmt.Println("pg type was not set... Skipping")
		}

		runtime, _ := cmd.Flags().GetString("runtime")
		if runtime == "" {
			fmt.Println("runtime was not set... Skipping")
		}

		fmt.Println("Printing name, movieYear,pgType, runtime")
		fmt.Println(name, movieYear, pgType, runtime)
		data,err := functions.GetAvailableTheaters()


		theater := SelecTheather(data)

		schedule := make(map[string][]string)

		//loop to ask for new entries
		for {
			date := inputDate()
			inputTimes(schedule, date)

			if !IfMoreDates() {
				break
			}
		}

		functions.AddMovie(name, movieYear,pgType,runtime,theater,schedule)


		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		return nil
	},
}

func IfMoreDates()  bool  {
	validate := func(input string) error {
		if input == "y" || input == "n"{
			return nil
		}
		return errors.New("invalid input")
	}

	prompt := promptui.Prompt{
		Label:    "Do you want to add more dates to the movie schedule, input (y/n)?",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	if result == "y" {
		return true
	}else {
		return false
	}




}

func inputDate() string  {
	validate := func(input string) error {
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")
		if re.MatchString(input){
			return nil
		}
		return errors.New("invalid date or data format")
	}

	prompt := promptui.Prompt{
		Label:    "Schedule a date for your movie in the format: dd-mm-yyyy",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	fmt.Println(result)
	return  result



}

func inputTimes( schedule map[string][]string, date string)   {
	validate := func(input string) error {
		slice := strings.Split(input, ",")
		re := regexp.MustCompile("^(0[0-9]|1[0-9]|2[0-3]|[0-9]):[0-5][0-9]$")

		for _, time := range slice{
			if !re.MatchString(time){
				return fmt.Errorf("invalid time format or time for time entry: %s", time)
			}
		}
		return nil


	}

	prompt := promptui.Prompt{
		Label:    "Enter a coma separated list of times in HH:MM time format",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	schedule[date] = append(schedule[date], result)

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

func SelecTheather(data []string) string  {

	prompt := promptui.Select{
		Label: "Select a Theater you want to screen your movie",
		Items: data,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	chosenTheater := strings.Split(result, " PgType:")[0]

	fmt.Printf("You choose %q\n",chosenTheater) //will be incorrect if the name movie name contains the separator string :(

	return chosenTheater

}


func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
}


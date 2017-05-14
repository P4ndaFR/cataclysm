package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	yaml "gopkg.in/yaml.v1"

	"github.com/miton18/cataclysm/model"
	"github.com/miton18/cataclysm/util"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	c              model.Config
	password       []byte
	configPath     string
	configfilePath = "cataclysm.config.yaml"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise Cataclysm configuration",
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {

	fmt.Println("\nWelcome to the Cataclysm config generator")
	for {
		if ok := askUsername(); ok {
			break
		}
	}
	for {
		if ok := askPassword(); ok {
			break
		}
	}
	askConfigPath()

	fmt.Println("\nLet me crypt your password")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	c.Password = string(hashedPassword)

	filePath := path.Join(configPath, configfilePath)

	fmt.Println(fmt.Sprintf("\n%s, it's ok, I will write your configuration file at '%s'", c.Username, filePath))

	_, err = os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	} else {
		var res string
		fmt.Print("\nThis file alreadey exist, override it ? (y/n): ")
		fmt.Scanln(&res)
		if !util.ContainsString([]string{"y", "yes"}, strings.ToLower(res)) {
			return
		}
	}

	c, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filePath, c, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nYour starting configuration is ready, just launch Cataclysm with:")
	fmt.Println(fmt.Sprintf("cataclysm -c %s", filePath))
}

func askUsername() (ok bool) {
	ok = false
	fmt.Println("\nFirst of all, please choose a Username:")
	_, err := fmt.Scanln(&c.Username)
	if err != nil {
		fmt.Println("Failed to get your username", err)
		return
	}
	return true
}

func askPassword() (ok bool) {
	ok = false
	fmt.Println(fmt.Sprintf("\nOK %s, now choose a Password :", c.Username))

	password, err := terminal.ReadPassword(0)
	if err != nil {
		panic(err)
		return
	}
	tmp := string(password)

	if len(tmp) < 8 {
		fmt.Println("It's too weak, 8 characters minimum")
		return
	}

	fmt.Println("One more time ?")
	passTwo, err := terminal.ReadPassword(0)
	if err != nil {
		panic(err)
	}

	if tmp != string(passTwo) {
		fmt.Println("You don't write the same passwords")
		return
	}
	return true
}

func askConfigPath() {
	fmt.Println("\nLast question, where can I put your configutation file?")
	fmt.Println("ex: /home/me/.cataclysm")
	_, err := fmt.Scanln(&configPath)
	if err != nil {
		fmt.Println("Oh no, I can't do this here", err)
		askConfigPath()
	}
}

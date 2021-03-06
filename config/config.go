package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/BlueMedoraPublic/bpcli/util/uuid"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

// account stores users BindPlane account information
type account struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	Current bool   `json:"current"`
}

// AddAccount appends an account to the configuration file
func AddAccount(name string, key string) error {
	envWarning()

	accounts, err := read()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") == false {
			return err
		}

		if err := create(); err != nil {
			return err
		}
	}

	accounts, err = read()
	if err != nil {
		return err
	}

	if err := validateNewAccount(accounts, name, key); err != nil {
		return err
	}

	a := account{Name: name, Key: key, Current: false}
	newList := append(accounts, a)

	newListBytes, err := json.Marshal(newList)
	if err != nil {
		return err
	}

	return write(newListBytes)
}

/*
CurrentAPIKey returns the API key found in the environment,
or the 'current' API key found in the credentials file if
the environment is not set
*/
func CurrentAPIKey() (string, error) {
	apiKey, found, err := currentAPIKeyENV()

	// return an error if env is found but malformed
	if found == true && err != nil {
		return "", err
	}

	// return api key if found
	if found == true && err == nil {
		return apiKey, nil
	}

	apiKey, err = currentAccount()
	if err != nil {
		// return both ENV and File errors
		//return "", errors.Wrap(err, e.Error())
		return "", err
	}
	return apiKey, nil
}

// ListAccounts prints a formatted list of users read from the configuration file
func ListAccounts() error {
	envWarning()

	currentList, err := read()
	if err != nil {
		return errors.Wrap(err, fileNotFoundError().Error())
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	if len(currentList) == 0 {
		return errors.New(path + " is empty, add an account with 'bpcli account add'")
	}

	fmt.Println("List of Accounts and API Keys. * Denotes Current Account")

	// Print the list in a formatted way
	for _, acc := range currentList {
		if acc.Current == true {
			fmt.Println("* "+acc.Name, acc.Key)
		} else {
			fmt.Println(acc.Name, acc.Key)
		}
	}
	return nil
}

// Remove erases an account from the configuration file
func Remove(name string) error {

	currentList, err := read()
	if err != nil {
		return errors.Wrap(err, fileNotFoundError().Error())
	}

	newList := currentList

	if !(len(newList) > 0) {
		return errors.New("The account list exists, but it is empty")
	}

	for i := 0; i < (len(newList)); i++ {
		if name != newList[i].Name {
			continue
		} else {
			newList = append(newList[:i], newList[i+1:]...)
			break
		}
	}

	if cmp.Equal(newList, currentList) {
		os.Stderr.WriteString("No names match the given input" +
			"Name Given: " + name)
		return nil
	}

	newListBytes, err := json.Marshal(newList)
	if err != nil {
		return err
	}

	return write(newListBytes)
}

// SetCurrent sets a chosen account to be the current account being worked in
func SetCurrent(name string) error {
	envWarning()

	currentList, err := read()
	if err != nil {
		return errors.Wrap(err, fileNotFoundError().Error())
	}

	b, err := accountExists(name)
	if err != nil {
		return err
	}
	if b == false {
		return accountNotFoundError(name)
	}

	for i := range currentList {
		if name == currentList[i].Name {
			currentList[i].Current = true
		} else {
			currentList[i].Current = false
		}
	}

	updatedListBytes, err := json.Marshal(currentList)
	if err != nil {
		return err
	}

	return write(updatedListBytes)
}

// currentAPIKeyENV returns the API key, true, and nil if
// the API key is found in the environment and is a valid uuid
// returns false if the environment is empty
func currentAPIKeyENV() (string, bool, error) {
	a := os.Getenv("BINDPLANE_API_KEY")

	if len(strings.TrimSpace(a)) == 0 {
		return "", false, errors.New("ERROR: The BINDPLANE_API_KEY environment variable is not set")
	}

	if !uuid.IsUUID(a) {
		return "", true, errors.New("ERROR: The BINDPLANE_API_KEY environment variable is not a valid uuid")
	}

	return a, true, nil
}

func currentAccount() (string, error) {
	accounts, err := read()
	if err != nil {
		return "", err
	}

	for _, a := range accounts {
		if a.Current == true {
			if uuid.IsUUID(a.Key) {
				return a.Key, nil
			}
			//return a.Key, nil
			return "", errors.New("Found current account in config, '" + a.Name + "', however, the API key is not a valid UUID")
		}
	}
	return "", noCurrentAccountError()
}

func accountExists(name string) (bool, error) {
	currentList, err := read()
	if err != nil {
		return false, err
	}

	for i := range currentList {
		if name == currentList[i].Name {
			return true, nil
		}
	}
	return false, nil
}

func validateNewAccount(accounts []account, name string, key string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return errors.New("The name cannot be an empty string")
	}

	if !uuid.IsUUID(key) {
		return errors.New("The API Key given is not a valid UUID")
	}

	b, err := uniqueUUID(accounts, key)
	if err != nil {
		return err
	}
	if b == false {
		return errors.New("The API Key given already exists within the config file")
	}

	n, err := uniqueName(accounts, name)
	if err != nil {
		return err
	}
	if n == false {
		return errors.New("The name given already exists within the config file")
	}

	return nil
}

func envWarning() {
	x := os.Getenv("BINDPLANE_API_KEY")
	if len(x) > 0 {
		fmt.Fprintf(os.Stderr, "WARNING: BINDPLANE_API_KEY is set and will take precidence over the configuration file\n")
	}
}

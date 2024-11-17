package internal

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
	"os"
	"os/user"
	"strconv"
)

type ConfigPath struct {
	Path string
	File string
}

type Config struct {
	Accounts []Account `yaml:"accounts"`
}

type Account struct {
	Name   string `yaml:"name"`
	Number string `yaml:"number"`
}

func GetWhoiamConfigDir() (*ConfigPath, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	path := usr.HomeDir + "/.whoiam"
	return &ConfigPath{Path: path}, nil
}

func NewConfigPath() (*ConfigPath, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	path := usr.HomeDir + "/.whoiam"
	return &ConfigPath{Path: path, File: "whoiam.yaml"}, nil
}

func NewTemplateConfig() (*Config, error) {
	var accounts []Account
	account := Account{Name: "my-account", Number: "123456789012"}
	accounts = append(accounts, account)

	return &Config{
		Accounts: accounts,
	}, nil
}

func ValidateAccountNumber(number string) error {
	if len(number) != 12 {
		return fmt.Errorf("account number must be 12 digits")
	}

	if _, err := strconv.Atoi(number); err != nil {
		return fmt.Errorf("account number must only contain digits")
	}
	return nil
}

func (c *ConfigPath) Exists() bool {
	_, err := os.Stat(c.Path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (c *ConfigPath) FullPath() string {
	return c.Path + "/" + c.File
}

func (c *ConfigPath) Create() error {
	err := os.MkdirAll(c.Path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigPath) LoadConfig() (*Config, error) {
	data, err := os.ReadFile(c.FullPath())
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (c *ConfigPath) SaveConfig(config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.FullPath(), data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) AccountExists(number string) bool {
	for _, account := range c.Accounts {
		if account.Number == number {
			return true
		}
	}
	return false
}

func (c *Config) AddAccount(name, number string) error {
	if err := ValidateAccountNumber(number); err != nil {
		return err
	}
	c.Accounts = append(c.Accounts, Account{Name: name, Number: number})
	return nil
}

func (c *Config) DeleteAccount(number string) error {
	for i, account := range c.Accounts {
		if account.Number == number {
			c.Accounts = append(c.Accounts[:i], c.Accounts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("account not found")
}

func (c *Config) PrintConfigTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Account Name", "Account Number"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	for _, account := range c.Accounts {
		table.Append([]string{account.Name, account.Number})
	}
	table.Render()
}

func HandelError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

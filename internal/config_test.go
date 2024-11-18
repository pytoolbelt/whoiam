package internal

import (
	"os"
	"testing"
)

func TestConfigPathFullPath(t *testing.T) {
	configPath, _ := NewConfigPath()
	expected := configPath.Path + "/" + configPath.File
	result := configPath.FullPath()
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestConfigPathCreate(t *testing.T) {
	configPath, _ := NewConfigPath()
	err := configPath.Create()
	if err != nil {
		t.Errorf("expected nil, got error %v", err)
	}
	if !configPath.Exists() {
		t.Errorf("expected true, got false")
	}
	os.RemoveAll(configPath.Path)
}

func TestConfigPathLoadConfig(t *testing.T) {
	configPath, _ := NewConfigPath()
	configPath.Create()
	config, err := NewTemplateConfig()
	if err != nil {
		t.Errorf("expected nil, got error %v", err)
	}
	configPath.SaveConfig(config)
	loadedConfig, err := configPath.LoadConfig()
	if err != nil {
		t.Errorf("expected nil, got error %v", err)
	}
	if len(loadedConfig.Accounts) != 1 {
		t.Errorf("expected 1 account, got %v", len(loadedConfig.Accounts))
	}
	os.RemoveAll(configPath.Path)
}

func TestConfigPathSaveConfig(t *testing.T) {
	configPath, _ := NewConfigPath()
	configPath.Create()
	config, err := NewTemplateConfig()
	if err != nil {
		t.Errorf("expected nil, got error %v", err)
	}
	err = configPath.SaveConfig(config)
	if err != nil {
		t.Errorf("expected nil, got error %v", err)
	}
	os.RemoveAll(configPath.Path)
}

func TestConfigAccountExists(t *testing.T) {
	config, _ := NewTemplateConfig()
	if !config.AccountExists("account") {
		t.Errorf("expected true, got false")
	}
	if config.AccountExists("nonexistent") {
		t.Errorf("expected false, got true")
	}
}

func TestConfigAddAccount(t *testing.T) {
	config, _ := NewTemplateConfig()
	err := config.AddAccount("newaccount", "123456789012")
	if err != nil {
		t.Errorf("expected nil, got error %v", err)
	}
	if !config.AccountExists("newaccount") {
		t.Errorf("expected true, got false")
	}
}

func TestConfigDeleteAccount(t *testing.T) {
	config, _ := NewTemplateConfig()
	config.DeleteAccount("account")
	if config.AccountExists("account") {
		t.Errorf("expected false, got true")
	}
}

func TestConfigGetAccountByNumber(t *testing.T) {
	config, _ := NewTemplateConfig()
	result := config.GetAccountByNumber("123456789012")
	if result != "account" {
		t.Errorf("expected account, got %v", result)
	}
	result = config.GetAccountByNumber("nonexistent")
	if result != "" {
		t.Errorf("expected empty string, got %v", result)
	}
}

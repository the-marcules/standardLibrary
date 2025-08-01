package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Scenario struct {
	ScenarioName    string
	ScenarioKeyName string
}
type Scenarios struct {
	Sign    Scenario
	Verify  Scenario
	Encrypt Scenario
	Decrypt Scenario
}

type Config struct {
	ApiUserName         string
	AccountName         string
	ScenarioEncryptName string
	AccessKey           string
	SecretKey           string
	BaseUrl             string
	Scenarios           Scenarios
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file" + err.Error())

	}

	return &Config{
		ApiUserName:         os.Getenv("API_USER_NAME"),
		AccountName:         os.Getenv("ACCOUNT_NAME"),
		ScenarioEncryptName: os.Getenv("SCENARIO_ENCRYPT_NAME"),
		AccessKey:           os.Getenv("API_ACCESS_KEY"),
		SecretKey:           os.Getenv("API_SECRET_KEY"),
		BaseUrl:             os.Getenv("BASE_URL"),
		Scenarios: Scenarios{
			Sign: Scenario{
				ScenarioName:    os.Getenv("SCENARIO_SIGN_NAME"),
				ScenarioKeyName: os.Getenv("SCENARIO_SIGN_KEY_NAME"),
			},
			Verify: Scenario{
				ScenarioName:    os.Getenv("SCENARIO_VERIFY_NAME"),
				ScenarioKeyName: os.Getenv("SCENARIO_VERIFY_KEY_NAME"),
			},
			Encrypt: Scenario{
				ScenarioName:    os.Getenv("SCENARIO_ENCRYPT_NAME"),
				ScenarioKeyName: os.Getenv("SCENARIO_ENCRYPT_KEY_NAME"),
			},
			Decrypt: Scenario{
				ScenarioName:    os.Getenv("SCENARIO_DECRYPT_NAME"),
				ScenarioKeyName: os.Getenv("SCENARIO_DECRYPT_KEY_NAME"),
			},
		},
	}
}

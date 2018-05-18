package main

import (
	"reflect"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)
type DBConfig struct {
	Host string `json:"host" yaml:"host,omitempty"`
  Port string `yaml:"port"`
  Username string `yaml:"username"`
  Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type DBYamlConfig struct {
	Development DBConfig `yaml:"development"`
	Test DBConfig `yaml:"test"`
	Production DBConfig `yaml:"production"`
}

// database default value
// i think in development need default value, in test & production don't need default value
var defaultBaseDBConfig = DBConfig {
	Host: "loclhost", 
	Port: "5432", 
	Username: "postgres", 
	Password: "postgres",
	Database: "postgres",
}

func main(){
	// read yaml file
	yarmFile, err := ioutil.ReadFile("database.yml")
	checkError(err)
	// use os.ExpandEnv to read env data
	yarmFile = []byte(os.ExpandEnv(string(yarmFile)))
	var databaseConfig DBYamlConfig
	err = yaml.Unmarshal(yarmFile, &databaseConfig)
	checkError(err)
	// merge development default value
	databaseConfig.Development.Merge(defaultBaseDBConfig)
	// return yaml value
	fmt.Println(databaseConfig)
}

// check error
func checkError(err error){
	if err != nil {
		panic(err)
	}
}

func (dbConfig *DBConfig) Merge(defaultDbConfig DBConfig) {
	// defaultBaseDBConfig value object 
	DbElem := reflect.ValueOf(&defaultBaseDBConfig).Elem()
	DBConfigElem := reflect.ValueOf(dbConfig).Elem()

	// NumField: fields list
	for i := 0; i < DbElem.NumField(); i++ {
		fieldName := DbElem.Type().Field(i).Name
		DBConfigName := DBConfigElem.FieldByName(fieldName)
		// if value is "" or nil use default value
		if DBConfigName.Interface() == "" || DBConfigName.Interface() == nil {
			DBConfigName.SetString(DbElem.Field(i).String())
		}
	}
}
package configs

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/timdevlet/mp4/internal/helpers"

	env "github.com/caarlos0/env/v6"
)

func NewConfigsFromEnv() *Configs {
	o := Configs{}

	return o.loadFromEnv()
}

//nolint // uniq code style
type Configs struct {
	// ENV - dev, prod, test
	ENV string `env:"ENV" default:"prod"`

	// APP
	LOG_LEVEL  string `env:"LOG_LEVEL" envDefault:"debug"`
	LOG_FORMAT string `env:"LOG_FORMAT" envDefault:"plain"`
}

func (o *Configs) Debug() {
	for key, value := range o.GetFieldsWithValues() {
		logrus.Debug(key + ": " + value)
	}
}

func (o *Configs) GetSecuredFilds() []string {
	secureFields := []string{}

	e := reflect.ValueOf(o).Elem()

	for i := 0; i < e.NumField(); i++ {
		field := e.Type().Field(i)

		if getStructTag(field, "secured") != "" {
			secureFields = append(secureFields, field.Name)
		}
	}

	return secureFields
}

func (o *Configs) GetFieldsWithValues() map[string]string {
	result := make(map[string]string)

	secureFields := o.GetSecuredFilds()

	e := reflect.ValueOf(o).Elem()
	for i := 0; i < e.NumField(); i++ {
		field := e.Type().Field(i)

		if helpers.InArray(field.Name, secureFields) {
			value := e.Field(i).String()

			switch lvalue := len(value); {
			case lvalue == 0:
				result[field.Name] = ""
			case lvalue > 5:
				value = value[:len(value)/10*7]
				result[field.Name] = value + "..."
			case lvalue > 0:
				result[field.Name] = "xxxxxxx"
			}
		} else {
			result[field.Name] = fmt.Sprintf("%v", e.Field(i).Interface())
		}
	}

	return result
}

func getStructTag(f reflect.StructField, tagName string) string {
	return f.Tag.Get(tagName)
}

func (o *Configs) loadFromEnv() *Configs {
	options := Configs{}
	if err := env.Parse(&options); err != nil {
		panic(err)
	}

	return &options
}

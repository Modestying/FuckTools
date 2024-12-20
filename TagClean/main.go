package main

import (
	"fmt"
	"go/format"
	"log"
	"reflect"
	"strings"
)

type Config struct {
	ID   int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id" xx:"xxx"`
	Name string `gorm:"size:255" json:"name"`
	Age  int    `gorm:"column:age" json:"age"`
}

func generateStructWithTag(structure interface{}, tagNames []string) (string, error) {
	t := reflect.TypeOf(structure)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("input must be a struct or pointer to struct")
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", t.Name()))
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		var tagInfo string
		for _, tag := range tagNames {
			tagStr := field.Tag.Get(tag)
			// ignore tag not exist
			if len(tagStr) != 0 {
				tagInfo = fmt.Sprintf("%s %s:\"%s\"", tagInfo, tag, field.Tag.Get(tag))
			}
		}
		if tagInfo != "" {
			builder.WriteString(fmt.Sprintf("    %s %s `%s`\n", field.Name, field.Type, tagInfo))
		} else {
			builder.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
		}

	}
	builder.WriteString("}\n")

	code, err := format.Source([]byte(builder.String()))
	if err != nil {
		return "", err
	}
	return string(code), nil
}

func main() {
	result, err := generateStructWithTag(Config{}, []string{"json", "xx"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Generated struct:")
	fmt.Println(result)
}

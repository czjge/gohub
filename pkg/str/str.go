package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// user -> users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

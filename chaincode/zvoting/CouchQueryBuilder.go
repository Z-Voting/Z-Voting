package main

import "fmt"

type CouchQueryBuilder struct {
	start         string
	selectorStart string
	selectorBody  string
	selectorEnd   string
	end           string
}

func newCouchQueryBuilder() *CouchQueryBuilder {
	return &CouchQueryBuilder{start: "{", selectorStart: "\"selector\":{", selectorBody: "", selectorEnd: "}", end: "}"}
}

func (q *CouchQueryBuilder) addSelector(key string, value interface{}) *CouchQueryBuilder {
	if q.selectorBody != "" {
		q.selectorBody = q.selectorBody + ","
	}
	var addedString string
	switch v := value.(type) {
	case string:
		addedString = fmt.Sprintf("\"%s\":\"%v\"", key, value)
	case []byte:
		addedString = fmt.Sprintf("\"%s\":\"%v\"", key, value)
	default:
		addedString = fmt.Sprintf("\"%s\":%v", key, value)
		fmt.Printf("%q", v)
	}
	q.selectorBody = q.selectorBody + addedString
	return q
}

func (q *CouchQueryBuilder) addSelectorWithOperator(key string, operator string, value interface{}) *CouchQueryBuilder {
	if q.selectorBody != "" {
		q.selectorBody = q.selectorBody + ","
	}
	var addedString string
	switch v := value.(type) {
	case string:
		addedString = fmt.Sprintf("\"%s\":{\"%s\":\"%s\"}", key, operator, value)
	case []byte:
		addedString = fmt.Sprintf("\"%s\":{\"%s\":\"%s\"}", key, operator, value)
	default:
		addedString = fmt.Sprintf("\"%s\":{\"%s\":%v}", key, operator, value)
		fmt.Printf("%q", v)
	}
	q.selectorBody = q.selectorBody + addedString
	return q
}

func (q *CouchQueryBuilder) getQueryString() string {
	return q.start + q.selectorStart + q.selectorBody + q.selectorEnd + q.end
}

package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type QueryBuilder struct {
	table    string
	columns  []string
	selects  []string
	wheres   []string
	orderBys []string
	limit    int
	offset   int

	insertData interface{}
	updateData interface{}
	deleteData interface{}
}

func NewQueryBuilder(table string) *QueryBuilder {
	return &QueryBuilder{
		table:   table,
		columns: make([]string, 0),
	}
}

func (qb *QueryBuilder) Entity(entity interface{}) *QueryBuilder {
	// get fields from struct type if qb.entity is exist
	if entity != nil {
		entityTyp := reflect.TypeOf(entity)
		for i := 0; i < entityTyp.NumField(); i++ {
			qb.columns = append(qb.columns, StringPascalToSnake(entityTyp.Field(i).Name))
		}
	}

	return qb
}

func (qb *QueryBuilder) Select(columns ...string) *QueryBuilder {
	qb.selects = columns
	return qb
}

func (qb *QueryBuilder) Where(conditions ...string) *QueryBuilder {
	qb.wheres = append(qb.wheres, conditions...)
	return qb
}

func (qb *QueryBuilder) WhereStruct(data interface{}) *QueryBuilder {
	return qb.processStruct(data)
}

func (qb *QueryBuilder) processStruct(data interface{}) *QueryBuilder {
	// convert struct to map string interface
	var dataMap map[string]interface{}
	dataJson, _ := json.Marshal(data)
	json.Unmarshal(dataJson, &dataMap)

	// loop through map
	for key, value := range dataMap {
		// if type is map, recursively process it with updated prefix
		if reflect.TypeOf(value).Kind() == reflect.Map {
			qb.processStruct(value)
			continue
		}

		// if theres a start_date
		if key == "start_date" {
			condition := fmt.Sprintf("created_at >= '%v'", value)
			qb.wheres = append(qb.wheres, condition)
			continue
		}

		// if theres a end_date
		if key == "end_date" {
			condition := fmt.Sprintf("created_at <= '%v'", value)
			qb.wheres = append(qb.wheres, condition)
			continue
		}

		// if columns is not empty, check if column is in columns
		if len(qb.columns) > 0 {
			// if column is not in columns, skip
			if !StringInSlice(qb.columns, key) {
				continue
			}
		}

		// if type is string, wrap with single quote
		if reflect.TypeOf(value).Kind() == reflect.String {
			condition := fmt.Sprintf("%s = '%v'", key, value)
			qb.wheres = append(qb.wheres, condition)
			continue
		}

		// if type is int, float, or bool, no need to wrap with single quote
		condition := fmt.Sprintf("%s = %v", key, value)
		qb.wheres = append(qb.wheres, condition)
	}
	return qb
}

func (qb *QueryBuilder) In(column string, data []string) *QueryBuilder {
	values := make([]string, 0)

	for i := range data {
		// if type is string, wrap with single quote
		if reflect.TypeOf(data[i]).Kind() == reflect.String {
			values = append(values, fmt.Sprintf("'%v'", data[i]))
			continue
		}

		// if type is int, float, or bool, no need to wrap with single quote
		values = append(values, fmt.Sprintf("%v", data[i]))
	}

	condition := fmt.Sprintf("%s IN (%s)", column, strings.Join(values, ", "))
	qb.wheres = append(qb.wheres, condition)
	return qb
}

func (qb *QueryBuilder) OrderBy(columns ...string) *QueryBuilder {
	qb.orderBys = columns
	return qb
}

func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *QueryBuilder) Insert(data interface{}) *QueryBuilder {
	qb.insertData = data
	return qb
}

func (qb *QueryBuilder) Update(data interface{}) *QueryBuilder {
	qb.updateData = data
	return qb
}

func (qb *QueryBuilder) Delete() *QueryBuilder {
	qb.deleteData = true
	return qb
}

func (qb *QueryBuilder) buildInsert(args []interface{}) string {
	query := strings.Builder{}
	query.WriteString("INSERT INTO ")
	query.WriteString(qb.table)

	columns := make([]string, 0)
	values := make([]string, 0)

	// convert struct to map string interface
	var dataMap map[string]interface{}
	dataJson, _ := json.Marshal(qb.insertData)
	json.Unmarshal(dataJson, &dataMap)

	// loop through map
	for key, value := range dataMap {
		columns = append(columns, StringPascalToSnake(key))

		// if type is string, wrap with single quote
		if reflect.TypeOf(value).Kind() == reflect.String {
			values = append(values, fmt.Sprintf("'%v'", value))
			continue
		}

		// if type is int, float, or bool, no need to wrap with single quote
		values = append(values, fmt.Sprintf("%v", value))
	}

	query.WriteString(fmt.Sprintf("(%s) VALUES (%s)", strings.Join(columns, ", "), strings.Join(values, ", ")))

	return query.String()
}

func (qb *QueryBuilder) buildUpdate(args []interface{}) string {
	query := strings.Builder{}
	query.WriteString("UPDATE ")
	query.WriteString(qb.table)
	query.WriteString(" SET ")

	// convert struct to map string interface
	var dataMap map[string]interface{}
	dataJson, _ := json.Marshal(qb.updateData)
	json.Unmarshal(dataJson, &dataMap)

	updates := make([]string, 0)

	// loop through map
	for key, value := range dataMap {
		// if type is string, wrap with single quote
		if reflect.TypeOf(value).Kind() == reflect.String {
			updates = append(updates, fmt.Sprintf("%s = '%v'", key, value))
			continue
		}

		// if type is int, float, or bool, no need to wrap with single quote
		updates = append(updates, fmt.Sprintf("%s = %v", key, value))
	}

	query.WriteString(strings.Join(updates, ", "))

	if len(qb.wheres) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(qb.wheres, " AND "))
	}

	return query.String()
}

func (qb *QueryBuilder) buildDelete(args []interface{}) string {
	query := strings.Builder{}
	query.WriteString("DELETE FROM ")
	query.WriteString(qb.table)

	if len(qb.wheres) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(qb.wheres, " AND "))
	}

	return query.String()
}

func (qb *QueryBuilder) Build() string {
	query := strings.Builder{}
	var args []interface{}

	if qb.insertData != nil {
		return qb.buildInsert(args)
	} else if qb.updateData != nil {
		return qb.buildUpdate(args)
	} else if qb.deleteData != nil {
		return qb.buildDelete(args)
	}

	query.WriteString("SELECT ")

	if len(qb.selects) > 0 {
		query.WriteString(strings.Join(qb.selects, ", "))
	} else {
		query.WriteString("*")
	}

	query.WriteString(" FROM ")
	query.WriteString(qb.table)

	if len(qb.wheres) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(qb.wheres, " AND "))
	}

	if len(qb.orderBys) > 0 {
		query.WriteString(" ORDER BY ")
		query.WriteString(strings.Join(qb.orderBys, ", "))
	}

	if qb.limit > 0 {
		query.WriteString(fmt.Sprintf(" LIMIT %d", qb.limit))
	}

	if qb.offset > 0 {
		query.WriteString(fmt.Sprintf(" OFFSET %d", qb.offset))
	}

	return query.String()
}

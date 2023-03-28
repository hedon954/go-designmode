package main

import (
	"fmt"
)

// DB is the interface of abstract
type DB interface {
	Connect()
	SetExporter(Exporter)
	ExportData()
}

// Exporter is the interface of implemention
type Exporter interface {
	Export()
}

type JSONExporter struct{}

func (j *JSONExporter) Export() {
	fmt.Println("Export data as JSON")
}

type XMLExporter struct{}

func (x *XMLExporter) Export() {
	fmt.Println("Export data as XML")
}

type ExcelExporter struct{}

func (e *ExcelExporter) Export() {
	fmt.Println("Export data as Excel")
}

type MySQL struct {
	exp Exporter
}

func (m *MySQL) Connect() {
	fmt.Println("Connect to MySQL")
}

func (m *MySQL) SetExporter(exporter Exporter) {
	m.exp = exporter
}

func (m *MySQL) ExportData() {
	m.exp.Export()
}

type Oracle struct {
	exp Exporter
}

func (o *Oracle) Connect() {
	fmt.Println("Connect to Oracle")
}

func (o *Oracle) SetExporter(exp Exporter) {
	o.exp = exp
}

func (o *Oracle) ExportData() {
	o.exp.Export()
}

// ExportData is used to export datas from db
func ExportData(db DB) {
	db.Connect()
	db.ExportData()
}

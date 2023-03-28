package main

func main() {
	mysql := &MySQL{}
	oracle := &Oracle{}

	jsonExp := &JSONExporter{}
	xmlExp := &XMLExporter{}
	excelExp := &ExcelExporter{}

	mysql.SetExporter(jsonExp)
	ExportData(mysql)
	mysql.SetExporter(xmlExp)
	ExportData(mysql)

	oracle.SetExporter(jsonExp)
	ExportData(oracle)
	oracle.SetExporter(xmlExp)
	ExportData(oracle)
	oracle.SetExporter(excelExp)
	ExportData(oracle)
}

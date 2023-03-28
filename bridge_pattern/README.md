# 桥接模式

某业务系统, 现需要开发数据库导出工具, 根据SQL语句导出表数据到文件，数据库类型有多种, 目前需要支持MySQL, Orache 未来可能支持 SQLServer。导出格式可能有多种, 目前需要支持 CSV 和 JSON 格式

此场景下, 数据库类型是一种维度, 导出格式是另一种维度，组合可能性是乘法关系，即数据可以从 MySQL 读出后，导出成 CSV 或者JSON 格式，对于 Oracle 也是同样的情况。

如果我们用常规的继承来实现这个数据库导出模块，模块中首先要有一个类似抽象基础类的基类，然后再用继承分别实现：MySQL-CSV 导出类、MySQL- JSON 导出类、Oracle-CSV 导出类、Oracle-JSON 导出类，如果以后模块再加一种支持的数据库 SQLServer 和导出格式 XML，那么系统里实现类就更多了。

那么此时我们换一种思路，将"导出工具"分离出"**数据抓取**"和"**数据导出**"两个维度, 以便自由扩展、互相组合，从而减少类数目。这便是使用**桥接模式解决“需求多维度变化时系统会变臃肿的核心思想。**

## 1. 概念

桥接模式（Bridge Pattern）是一种将抽象与实现分离的设计模式，它的核心思想是通过一个抽象部分和一个实现部分的分离，使它们可以独立地变化，从而提高系统的可扩展性和可维护性。

在桥接模式中，抽象部分和实现部分分别定义了一个接口，并且通过一个桥接类来将它们联系起来。桥接类包含了一个抽象部分的引用和一个实现部分的引用，用来实现抽象部分和实现部分的关联。

- 抽象部分包含了客户端所需的接口和方法，它是客户端和具体实现部分之间的桥梁。
- 实现部分包含了实现抽象部分的具体代码，它负责实现抽象部分定义的接口和方法。

通过将抽象部分和实现部分分离开来，桥接模式可以使得它们可以独立地变化，从而可以轻松地添加新的抽象部分和实现部分，从而提高了系统的可扩展性和可维护性。

桥接模式的优点：

- 抽象和实现的分离：桥接模式可以将抽象和实现部分分离开来，使它们可以独立地变化，从而提高了系统的可扩展性和可维护性。
- 更好的扩展性：通过使用桥接模式，可以轻松地添加新的抽象部分和实现部分，从而使系统更具有扩展性。
- 透明性：桥接模式可以隐藏系统的实现细节，使得客户端可以专注于业务逻辑，而不必关心系统的具体实现。
- 代码复用：桥接模式可以将多个对象之间共享的实现部分提取出来，从而减少了代码的重复，提高了代码的复用性。

桥接模式的缺点：

- 增加了系统的理解与设计难度。
- 需要正确地识别系统中两个（或者多个）独立变化的维度，这一条也是桥接模式的难点。

## 2. 理解

桥接模式的核心就是抽离出影响业务的多个独立的维度，这样他们就可以进行自由组合，从而减少大量的重复代码和类。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/5b1792e20001370e12400705.jpg)

## 4. 实现

假设我们现在有一个需求：从不同类型的数据库中导出不同类型的数据格式，包括 JSON、XML 和 Excel。具体要求如下：

- 从 MySQL 中导出 JSON 和 XML 格式的数据
- 从 Oracle 中导出 JSON、XML 和 Excel 格式的数据

### 4.1 定义抽象部分和实现部分

```go
// DB is the interface of abstract component
type DB interface {
	Connect()
	SetExporter(Exporter)
	ExportData()
}

// Exporter is the interface of implemention component
type Exporter interface {
	Export()
}
```

### 4.2 实现 DB

```go
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
```

### 4.3 实现 Exporter

```go
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
```

### 4.4 实现数据导出方法

```go
// ExportData is used to export datas from db
func ExportData(db DB) {
	db.Connect()
	db.ExportData()
}
```

### 4.5 使用

```go
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
```

## 5. 场景

1. 需要将抽象和实现部分分离，并允许它们可以独立地变化。
2. 当一个类有多个变化的维度时，可以使用桥接模式来将它们分离，从而使得每个维度的变化都不会影响到其他维度。
3. 当需要在运行时动态地切换抽象和实现部分的实现时，可以使用桥接模式来实现这一需求。
4. 当不希望使用继承或者混入来扩展一个类的功能时，可以使用桥接模式。
5. 当需要将实现部分的代码独立出来，以便将来可以进行单独的维护和测试时，可以使用桥接模式。
6. 当需要对一个类的多个实现进行统一管理时，可以使用桥接模式来实现这一需求。



## 参考

- ChatGPT
- [Go设计模式--桥接模式，让代码既能多维度扩展又不会臃肿](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497649&idx=1&sn=81740f11f67e84c1aa21d701eed45ad9&chksm=fa832626cdf4af30df7fa606b9ba2f84b99b4d33e6a31d86d32ba704c424d970c4add80fcace&scene=178&cur_album_id=2531498848431669249#rd)
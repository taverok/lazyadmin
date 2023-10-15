package page

type Page struct {
	Name    string
	Content any
	Menu    []string
}

type Table struct {
	Fields []*Field
	Data   [][]any
}

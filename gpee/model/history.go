package model

type Table struct {
	Body body `xml:"tbody"`
}

type body struct {
	Rows []Row `xml:"tr"`
}

type Row struct {
	Cell []cell `xml:"td"`
}

type cell struct {
	Content string `xml:",innerxml"`
}

type HistoryListRow struct {
	Id     string
	Code   string
	Date   string
	Status string
}

type HistoryDataRow struct {
	Date string
	Hour string
	E    string
	P    string
}

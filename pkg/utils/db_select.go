package utils

import "strings"

type DbCond struct {
	Page   *Page
	Where  map[string]any
	Likes  []*Like
	Range  []*DbRangeCond
	LRange []*DbSingleRangeCond
	RRange []*DbSingleRangeCond
	Sorts  Sorts
}

type Like struct {
	Column string
	Value  string
}

func (l *Like) Sql() string {
	return l.Column + " like " + "'%" + l.Value + "%'"
}

type DbRangeCond struct {
	Column string
	Left   any
	Right  any
}

type DbSingleRangeCond struct {
	Column string
	Value  any
}

type Sort struct {
	Column string
	IsAsc  bool
}

type Page struct {
	Offset int
	Limit  int
}

type Sorts []*Sort

func (s *Sort) Sql() string {
	if s.IsAsc {
		return s.Column + " ASC"
	}
	return s.Column + " DESC"
}

func (s Sorts) Sql() string {
	sortSql := Slice[string]{}
	for _, sort := range s {
		sortSql = append(sortSql, sort.Sql())
	}
	return strings.Join(sortSql, ",")
}

func (db *DbCond) CleanPage() *DbCond {
	db.Page = nil
	return db
}

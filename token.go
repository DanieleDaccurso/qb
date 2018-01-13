package qb

type tokenInterface interface{}

type selectToken struct {
	name []string
}

type deleteToken struct {
}

type updateToken struct {
	name  string
	alias string
}

type setToken struct {
	key   string
	value string
}

type fromToken struct {
	from  string
	alias string
	index string
}

type joinToken struct {
	joinType      string
	join          string
	alias         string
	conditionType string
	condition     string
}

type whereToken struct {
	prefix string
	where  string
}

type groupByToken struct {
	group string
}

type havingToken struct {
	prefix string
	having string
}

type orderToken struct {
	sort  string
	order string
}

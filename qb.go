package qb

type QueryBuilder struct {
	master     tokenInterface
	tokens     []tokenInterface
	parameters map[string]interface{}
}

func (q *QueryBuilder) Select(selects ...string) *QueryBuilder {
	st := new(selectToken)
	for _, s := range selects {
		st.name = append(st.name, s)
	}
	q.master = st
	return q
}

func (q *QueryBuilder) AddSelect(selects ...string) *QueryBuilder {
	st := q.master.(*selectToken)
	for _, s := range selects {
		st.name = append(st.name, s)
	}
	q.master = st
	return q
}

func (q *QueryBuilder) Delete() *QueryBuilder {
	q.master = &deleteToken{}
	return q
}

func (q *QueryBuilder) Update(u string) *QueryBuilder {
	q.master = &updateToken{name: u}
	return q
}

func (q *QueryBuilder) UpdateWithAlias(u string, a string) *QueryBuilder {
	q.master = &updateToken{name: u, alias: a}
	return q
}

func (q *QueryBuilder) Set(k string, v string) *QueryBuilder {
	q.tokens = append(q.tokens, &setToken{key: k, value: v})
	return q
}

func (q *QueryBuilder) From(from string, alias string) *QueryBuilder {
	q.tokens = append(q.tokens, &fromToken{from: from, alias: alias})
	return q
}

func (q *QueryBuilder) FromWithIndex(from string, alias string, indexBy string) *QueryBuilder {
	q.tokens = append(q.tokens, &fromToken{from: from, alias: alias, index: indexBy})
	return q
}

func (q *QueryBuilder) Join(join string, alias string, conditionType string, condition string) *QueryBuilder {
	q.tokens = append(q.tokens, &joinToken{joinType: "inner", join: join, alias: alias, conditionType: conditionType, condition: condition})
	return q
}

func (q *QueryBuilder) InnerJoin(join string, alias string, conditionType string, condition string) *QueryBuilder {
	q.tokens = append(q.tokens, &joinToken{joinType: "inner", join: join, alias: alias, conditionType: conditionType, condition: condition})
	return q
}

func (q *QueryBuilder) LeftJoin(join string, alias string, conditionType string, condition string) *QueryBuilder {
	q.tokens = append(q.tokens, &joinToken{joinType: "left", join: join, alias: alias, conditionType: conditionType, condition: condition})
	return q
}

func (q *QueryBuilder) Where(where string) *QueryBuilder {
	q.tokens = append(q.tokens, &whereToken{where: where})
	return q
}

func (q *QueryBuilder) OrWhere(where string) *QueryBuilder {
	q.tokens = append(q.tokens, &whereToken{where: where, prefix: "or"})
	return q
}

func (q *QueryBuilder) AndWhere(where string) *QueryBuilder {
	q.tokens = append(q.tokens, &whereToken{where: where, prefix: "and"})
	return q
}

func (q *QueryBuilder) GroupBy(group string) *QueryBuilder {
	q.tokens = append(q.tokens, &groupByToken{group: group})
	return q
}

func (q *QueryBuilder) AddGroupBy(group string) *QueryBuilder {
	q.tokens = append(q.tokens, &groupByToken{group: group})
	return q
}

func (q *QueryBuilder) Having(having string) *QueryBuilder {
	q.tokens = append(q.tokens, &havingToken{having: having})
	return q
}

func (q *QueryBuilder) OrHaving(having string) *QueryBuilder {
	q.tokens = append(q.tokens, &havingToken{prefix: "or", having: having})
	return q
}

func (q *QueryBuilder) AndHaving(having string) *QueryBuilder {
	q.tokens = append(q.tokens, &havingToken{prefix: "and", having: having})
	return q
}

func (q *QueryBuilder) OrderBy(sort string, order string) *QueryBuilder {
	q.tokens = append(q.tokens, &orderToken{sort: sort, order: order})
	return q
}

func (q *QueryBuilder) AndOrderBy(sort string, order string) *QueryBuilder {
	q.tokens = append(q.tokens, &orderToken{sort: sort, order: order})
	return q
}

func (q *QueryBuilder) AddParameter(key string, value interface{}) *QueryBuilder {
	q.parameters[key] = value
	return q
}

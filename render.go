package qb

func Render(q *QueryBuilder) string {
	r := new(qbCompiler)
	r.compile()
	return r.q
}

type qbCompiler struct {
	q  string
	qb *QueryBuilder
}

func (r *qbCompiler) compile() {
	r.compileMaster()
	r.compileFrom()
	r.compileJoins()
	r.compileWhere()
	r.compileGroupBy()
	r.compileHaving()
	r.compileOrderBy()
}

func (r *qbCompiler) compileMaster() {
	switch r.qb.master.(type) {
	case *selectToken:
		master := r.qb.master.(*selectToken)
		r.q = "SELECT "
		for i, s := range master.name {
			r.q += s
			if i != len(master.name)-1 {
				r.q += ", "
			} else {
				r.q += " "
			}
		}
	case *deleteToken:
		r.q = "DELETE "
	case *updateToken:
		master := r.qb.master.(*updateToken)
		r.q = "UPDATE " + master.name + " "
		if master.alias != "" {
			r.q += "AS " + master.alias + " "
		}
	}
}

func (r *qbCompiler) compileFrom() {

}

func (r *qbCompiler) compileJoins() {

}

func (r *qbCompiler) compileWhere() {

}

func (r *qbCompiler) compileGroupBy() {

}

func (r *qbCompiler) compileHaving() {

}

func (r *qbCompiler) compileOrderBy() {

}

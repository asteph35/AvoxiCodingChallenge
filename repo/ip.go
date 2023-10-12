package repo

//func (r Repo) getCountry(ctx context.Context, name string) string {
//	query := `
//	SELECT * FROM avoxi.countries where country_name = $1
//`
//	queryArgs := []interface{}{name}
//
//	err := r.read.QueryRow(ctx, query, queryArgs).Scan()
//}
//
//func (r Repo) getCountry(ctx context.Context, name string) string {
//	query := `
//	SELECT * FROM avoxi.countries where country_name = $1
//`
//	queryArgs := []interface{}{name}
//
//	err := r.read.QueryRow(ctx, query, queryArgs).Scan()
//}

func (r Repo) GetNetwork(country string) (string, error) {
	//query := `SELECT * FROM avoxi.countries where country_name = $1`
	//	queryArgs := []interface{}{name}
	//
	//	err := r.read.QueryRow(ctx, query, queryArgs).Scan()
	return "", nil
}

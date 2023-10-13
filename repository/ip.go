package repository

func (r Repo) GetNetworks(country string) ([]string, error) {
	query := `
SELECT icb.network FROM avoxi.ip_country_blocks icb
JOIN avoxi.countries c on icb.geoname_id = c.geoname_id
	WHERE c.country_name like ? or c.country_iso_code = ?;
`

	networks := make([]string, 0)
	rows, err := r.db.Query(query, country, country)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var network string
		err := rows.Scan(&network)
		if err != nil {
			return nil, err
		}
		networks = append(networks, network)
	}

	return networks, nil
}

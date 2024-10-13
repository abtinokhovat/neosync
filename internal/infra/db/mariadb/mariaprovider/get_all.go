package mariaprovider

import (
	"context"
	"neosync/internal/domain/provider"
	"neosync/pkg/richerror"
)

func (d *DB) GetAll(ctx context.Context) (map[uint]provider.Provider, error) {
	const op = "mariaprovider.GetAll"
	const query = "select * from providers"

	rows, err := d.conn.Conn().QueryContext(ctx, query)
	if err != nil {
		// here we 100 have to have provider, so I skip to check no rows error
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	providers := make(map[uint]provider.Provider)
	for rows.Next() {
		p, sErr := scanProvider(rows)
		if sErr != nil {
			return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(sErr)
		}

		providers[p.ID] = p
	}

	return providers, nil
}

package infrastructure

import "openjyotish/internal/infrastructure/swiss"

type SwissAdapter struct{}

func (a SwissAdapter) Exec(opt *swiss.SwissOptions) (*swiss.Result, error) {
	res, err := swiss.ExecSwiss(opt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

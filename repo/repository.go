package repo

import (
	"github.com/denizcamalan/movie_app/operator"
)

type OperatorModel struct {
}

func NewOperatorModel() *OperatorModel {
	var model OperatorModel
	return &model
}
func (*OperatorModel) DB_Operator() operator.MovieModelMeneger {
	return operator.NewMovieModel()
}

func (*OperatorModel) Register_Operator() operator.RegisterManager {
	return operator.NewRegiterModel()
}

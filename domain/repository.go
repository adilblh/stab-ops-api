package domain

type OperatorRepository interface {
	InsertAction(operatorId string, action Action) error
	FindOperator(operatorId string, opts OperatorFilter) (*Operator, error)
	GetOperators(opts OperatorFilter) ([]*Operator, error)
	InsertOperators(ops []interface{})
}

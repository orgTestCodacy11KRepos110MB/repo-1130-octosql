package logical

import (
	"context"

	"github.com/cube2222/octosql/physical"
)

type UnionDistinct struct {
	first, second Node
}

func NewUnionDistinct(first, second Node) *UnionDistinct {
	return &UnionDistinct{first: first, second: second}
}

func (node *UnionDistinct) Typecheck(ctx context.Context, env physical.Environment, logicalEnv Environment) (physical.Node, map[string]string) {
	panic("implement me")
	// return NewDistinct(NewUnionAll(node.first, node.second)).Physical(ctx, physicalCreator)
}

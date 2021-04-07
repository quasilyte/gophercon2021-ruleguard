package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func nolintGocritic(m dsl.Matcher) {
	m.MatchComment(`//nolint:gocritic`).Report(`that is heartbreaking`)
}

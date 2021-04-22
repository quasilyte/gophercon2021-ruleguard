package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func valSwap(m dsl.Matcher) {
	m.Match(`$tmp := $y; $y = $x; $x = $tmp`).
		Suggest(`$x, $y = $y, $x`).
		Report(`could rewrite as '$x, $y = $y, $x'`)
}

func redundantCast(m dsl.Matcher) {
	m.Match(`string($x)`).Where(m[`x`].Type.Is(`string`)).
		Suggest(`$x`).
		Report(`$x is already string-typed`)
}

func assignOp(m dsl.Matcher) {
	m.Match(`$x = $x + $y`).Where(m["x"].Pure).
		Suggest(`$x += $y`).
		Report(`could rewrite as '$x += $y'`)
}

func yodaStyle(m dsl.Matcher) {
	m.Match(`$x > $y`).Where(m["x"].Const && !m["y"].Const).
		Suggest(`$y > $x`).
		Report(`avoid yoda-style expressions`)
}

func redundantBuffer(m dsl.Matcher) {
	m.Match(`bytes.NewBuffer($x).Bytes()`).
		Suggest(`$x`).
		Report(`can use $x byte slice directly`)
}

func returnElse(m dsl.Matcher) {
	m.Match(`if $*_; $_ { $*_; return $*_; } else { $*_ }`).
		Report(`if block ends with return, so drop this else`)
}

func suggestFprintf(m dsl.Matcher) {
	m.Match(`$w.Write([]byte(fmt.Sprintf($*args)))`).
		Suggest(`fmt.Fprintf($w, $args)`).
		Report(`could use fmt.Fprintf here: fmt.Fprintf($w, $args)`)
}

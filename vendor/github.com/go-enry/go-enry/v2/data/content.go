// Code generated by github.com/go-enry/go-enry/v2/internal/code-generator DO NOT EDIT.
// Extracted from github/linguist commit: 40992ba7f86889f80dfed3ba95e11e1082200bad

package data

import (
	"regexp"

	"github.com/go-enry/go-enry/v2/data/rule"
)

var ContentHeuristics = map[string]*Heuristics{
	".1": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".1in": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".1m": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".1x": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".2": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".3": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".3in": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".3m": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".3p": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".3pm": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".3qt": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".3x": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".4": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".5": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".6": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".7": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".8": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".9": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".as": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("ActionScript"),
			regexp.MustCompile(`(?m)^\s*(package\s+[a-z0-9_\.]+|import\s+[a-zA-Z0-9_\.]+;|class\s+[A-Za-z0-9_]+\s+extends\s+[A-Za-z0-9_]+)`),
		),
		rule.Always(
			rule.MatchingLanguages("AngelScript"),
		),
	},
	".asc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Public Key"),
			regexp.MustCompile(`(?m)^(----[- ]BEGIN|ssh-(rsa|dss)) `),
		),
		rule.Or(
			rule.MatchingLanguages("AsciiDoc"),
			regexp.MustCompile(`(?m)^[=-]+(\s|\n)|{{[A-Za-z]`),
		),
		rule.Or(
			rule.MatchingLanguages("AGS Script"),
			regexp.MustCompile(`(?m)^(\/\/.+|((import|export)\s+)?(function|int|float|char)\s+((room|repeatedly|on|game)_)?([A-Za-z]+[A-Za-z_0-9]+)\s*[;\(])`),
		),
	},
	".asm": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Motorola 68K Assembly"),
			regexp.MustCompile(`(?m)(?im)\bmoveq(?:\.l)?\s+#(?:\$-?[0-9a-f]{1,3}|%[0-1]{1,8}|-?[0-9]{1,3}),\s*d[0-7]\b|(?im)^\s*move(?:\.[bwl])?\s+(?:sr|usp),\s*[^\s]+|(?im)^\s*move\.[bwl]\s+.*\b[ad]\d|(?im)^\s*movem\.[bwl]\b|(?im)^\s*move[mp](?:\.[wl])?\b|(?im)^\s*btst\b|(?im)^\s*dbra\b`),
		),
	},
	".asy": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("LTspice Symbol"),
			regexp.MustCompile(`(?m)^SymbolType[ \t]`),
		),
		rule.Always(
			rule.MatchingLanguages("Asymptote"),
		),
	},
	".bb": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("BlitzBasic"),
			regexp.MustCompile(`(?m)(<^\s*; |End Function)`),
		),
		rule.Or(
			rule.MatchingLanguages("BitBake"),
			regexp.MustCompile(`(?m)^\s*(# |include|require)\b`),
		),
	},
	".builds": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)^(\s*)(?i:<Project|<Import|<Property|<?xml|xmlns)`),
		),
		rule.Always(
			rule.MatchingLanguages("Text"),
		),
	},
	".ch": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("xBase"),
			regexp.MustCompile(`(?m)^\s*#\s*(?i:if|ifdef|ifndef|define|command|xcommand|translate|xtranslate|include|pragma|undef)\b`),
		),
	},
	".cl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)^\s*\((?i:defun|in-package|defpackage) `),
		),
		rule.Or(
			rule.MatchingLanguages("Cool"),
			regexp.MustCompile(`(?m)^class`),
		),
		rule.Or(
			rule.MatchingLanguages("OpenCL"),
			regexp.MustCompile(`(?m)\/\* |\/\/ |^\}`),
		),
	},
	".cls": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("TeX"),
			regexp.MustCompile(`(?m)\\\w+{`),
		),
		rule.Or(
			rule.MatchingLanguages("ObjectScript"),
			regexp.MustCompile(`(?m)^Class\s`),
		),
	},
	".cs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Smalltalk"),
			regexp.MustCompile(`(?m)![\w\s]+methodsFor: `),
		),
		rule.Or(
			rule.MatchingLanguages("C#"),
			regexp.MustCompile(`(?m)^(\s*namespace\s*[\w\.]+\s*{|\s*\/\/)`),
		),
	},
	".d": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("D"),
			regexp.MustCompile(`(?m)^module\s+[\w.]*\s*;|import\s+[\w\s,.:]*;|\w+\s+\w+\s*\(.*\)(?:\(.*\))?\s*{[^}]*}|unittest\s*(?:\(.*\))?\s*{[^}]*}`),
		),
		rule.Or(
			rule.MatchingLanguages("DTrace"),
			regexp.MustCompile(`(?m)^(\w+:\w*:\w*:\w*|BEGIN|END|provider\s+|(tick|profile)-\w+\s+{[^}]*}|#pragma\s+D\s+(option|attributes|depends_on)\s|#pragma\s+ident\s)`),
		),
		rule.Or(
			rule.MatchingLanguages("Makefile"),
			regexp.MustCompile(`(?m)([\/\\].*:\s+.*\s\\$|: \\$|^[ %]:|^[\w\s\/\\.]+\w+\.\w+\s*:\s+[\w\s\/\\.]+\w+\.\w+)`),
		),
	},
	".dsp": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Microsoft Developer Studio Project"),
			regexp.MustCompile(`(?m)# Microsoft Developer Studio Generated Build File`),
		),
		rule.Or(
			rule.MatchingLanguages("Faust"),
			regexp.MustCompile(`(?m)\bprocess\s*[(=]|\b(library|import)\s*\(\s*"|\bdeclare\s+(name|version|author|copyright|license)\s+"`),
		),
	},
	".ecl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("ECLiPSe"),
			regexp.MustCompile(`(?m)^[^#]+:-`),
		),
		rule.Or(
			rule.MatchingLanguages("ECL"),
			regexp.MustCompile(`(?m):=`),
		),
	},
	".es": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Erlang"),
			regexp.MustCompile(`(?m)^\s*(?:%%|main\s*\(.*?\)\s*->)`),
		),
	},
	".f": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^: `),
		),
		rule.Or(
			rule.MatchingLanguages("Filebench WML"),
			regexp.MustCompile(`(?m)flowop`),
		),
		rule.Or(
			rule.MatchingLanguages("Fortran"),
			regexp.MustCompile(`(?m)^(?i:[c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)`),
		),
	},
	".for": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^: `),
		),
		rule.Or(
			rule.MatchingLanguages("Fortran"),
			regexp.MustCompile(`(?m)^(?i:[c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)`),
		),
	},
	".fr": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^(: |also |new-device|previous )`),
		),
		rule.Or(
			rule.MatchingLanguages("Frege"),
			regexp.MustCompile(`(?m)^\s*(import|module|package|data|type) `),
		),
		rule.Always(
			rule.MatchingLanguages("Text"),
		),
	},
	".fs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^(: |new-device)`),
		),
		rule.Or(
			rule.MatchingLanguages("F#"),
			regexp.MustCompile(`(?m)^\s*(#light|import|let|module|namespace|open|type)`),
		),
		rule.Or(
			rule.MatchingLanguages("GLSL"),
			regexp.MustCompile(`(?m)^\s*(#version|precision|uniform|varying|vec[234])`),
		),
		rule.Or(
			rule.MatchingLanguages("Filterscript"),
			regexp.MustCompile(`(?m)#include|#pragma\s+(rs|version)|__attribute__`),
		),
	},
	".gd": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("GAP"),
			regexp.MustCompile(`(?m)\s*(Declare|BindGlobal|KeyDependentOperation)`),
		),
		rule.Or(
			rule.MatchingLanguages("GDScript"),
			regexp.MustCompile(`(?m)\s*(extends|var|const|enum|func|class|signal|tool|yield|assert|onready)`),
		),
	},
	".gml": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)(?i:^\s*(\<\?xml|xmlns))`),
		),
		rule.Or(
			rule.MatchingLanguages("Graph Modeling Language"),
			regexp.MustCompile(`(?m)(?i:^\s*(graph|node)\s+\[$)`),
		),
		rule.Or(
			rule.MatchingLanguages("Gerber Image"),
			regexp.MustCompile(`(?m)\*\%$`),
		),
		rule.Always(
			rule.MatchingLanguages("Game Maker Language"),
		),
	},
	".gs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("GLSL"),
			regexp.MustCompile(`(?m)^#version\s+[0-9]+\b`),
		),
		rule.Or(
			rule.MatchingLanguages("Gosu"),
			regexp.MustCompile(`(?m)^uses java\.`),
		),
	},
	".h": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Objective-C"),
			regexp.MustCompile(`(?m)^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])`),
		),
		rule.Or(
			rule.MatchingLanguages("C++"),
			regexp.MustCompile(`(?m)^\s*#\s*include <(cstdint|string|vector|map|list|array|bitset|queue|stack|forward_list|unordered_map|unordered_set|(i|o|io)stream)>|^\s*template\s*<|^[ \t]*(try|constexpr)|^[ \t]*catch\s*\(|^[ \t]*(class|(using[ \t]+)?namespace)\s+\w+|^[ \t]*(private|public|protected):$|std::\w+`),
		),
	},
	".hh": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Hack"),
			regexp.MustCompile(`(?m)<\?hh`),
		),
	},
	".i": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Motorola 68K Assembly"),
			regexp.MustCompile(`(?m)(?im)\bmoveq(?:\.l)?\s+#(?:\$-?[0-9a-f]{1,3}|%[0-1]{1,8}|-?[0-9]{1,3}),\s*d[0-7]\b|(?im)^\s*move(?:\.[bwl])?\s+(?:sr|usp),\s*[^\s]+|(?im)^\s*move\.[bwl]\s+.*\b[ad]\d|(?im)^\s*movem\.[bwl]\b|(?im)^\s*move[mp](?:\.[wl])?\b|(?im)^\s*btst\b|(?im)^\s*dbra\b`),
		),
		rule.Or(
			rule.MatchingLanguages("SWIG"),
			regexp.MustCompile(`(?m)^[ \t]*%[a-z_]+\b|^%[{}]$`),
		),
	},
	".ice": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("JSON"),
			regexp.MustCompile(`(?m)\A\s*[{\[]`),
		),
		rule.Always(
			rule.MatchingLanguages("Slice"),
		),
	},
	".inc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Motorola 68K Assembly"),
			regexp.MustCompile(`(?m)(?im)\bmoveq(?:\.l)?\s+#(?:\$-?[0-9a-f]{1,3}|%[0-1]{1,8}|-?[0-9]{1,3}),\s*d[0-7]\b|(?im)^\s*move(?:\.[bwl])?\s+(?:sr|usp),\s*[^\s]+|(?im)^\s*move\.[bwl]\s+.*\b[ad]\d|(?im)^\s*movem\.[bwl]\b|(?im)^\s*move[mp](?:\.[wl])?\b|(?im)^\s*btst\b|(?im)^\s*dbra\b`),
		),
		rule.Or(
			rule.MatchingLanguages("PHP"),
			regexp.MustCompile(`(?m)^<\?(?:php)?`),
		),
		rule.Or(
			rule.MatchingLanguages("SourcePawn"),
			regexp.MustCompile(`(?m)^public\s+(?:SharedPlugin(?:\s+|:)__pl_\w+\s*=(?:\s*{)?|(?:void\s+)?__pl_\w+_SetNTVOptional\(\)(?:\s*{)?)`),
		),
		rule.Or(
			rule.MatchingLanguages("POV-Ray SDL"),
			regexp.MustCompile(`(?m)^\s*#(declare|local|macro|while)\s`),
		),
	},
	".l": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)\(def(un|macro)\s`),
		),
		rule.Or(
			rule.MatchingLanguages("Lex"),
			regexp.MustCompile(`(?m)^(%[%{}]xs|<.*>)`),
		),
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^\.[A-Za-z]{2}(\s|$)`),
		),
		rule.Or(
			rule.MatchingLanguages("PicoLisp"),
			regexp.MustCompile(`(?m)^\((de|class|rel|code|data|must)\s`),
		),
	},
	".lisp": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)^\s*\((?i:defun|in-package|defpackage) `),
		),
		rule.Or(
			rule.MatchingLanguages("NewLisp"),
			regexp.MustCompile(`(?m)^\s*\(define `),
		),
	},
	".ls": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("LoomScript"),
			regexp.MustCompile(`(?m)^\s*package\s*[\w\.\/\*\s]*\s*{`),
		),
		rule.Always(
			rule.MatchingLanguages("LiveScript"),
		),
	},
	".lsp": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)^\s*\((?i:defun|in-package|defpackage) `),
		),
		rule.Or(
			rule.MatchingLanguages("NewLisp"),
			regexp.MustCompile(`(?m)^\s*\(define `),
		),
	},
	".m": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Objective-C"),
			regexp.MustCompile(`(?m)^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])`),
		),
		rule.Or(
			rule.MatchingLanguages("Mercury"),
			regexp.MustCompile(`(?m):- module`),
		),
		rule.Or(
			rule.MatchingLanguages("MUF"),
			regexp.MustCompile(`(?m)^: `),
		),
		rule.Or(
			rule.MatchingLanguages("M"),
			regexp.MustCompile(`(?m)^\s*;`),
		),
		rule.And(
			rule.MatchingLanguages("Mathematica"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)\(\*`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)\*\)$`),
			),
		),
		rule.Or(
			rule.MatchingLanguages("MATLAB"),
			regexp.MustCompile(`(?m)^\s*%`),
		),
		rule.Or(
			rule.MatchingLanguages("Limbo"),
			regexp.MustCompile(`(?m)^\w+\s*:\s*module\s*{`),
		),
	},
	".man": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".mask": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Unity3d Asset"),
			regexp.MustCompile(`(?m)tag:unity3d.com`),
		),
	},
	".md": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Markdown"),
			regexp.MustCompile(`(?m)(^[-A-Za-z0-9=#!\*\[|>])|<\/|\A\z`),
		),
		rule.Or(
			rule.MatchingLanguages("GCC Machine Description"),
			regexp.MustCompile(`(?m)^(;;|\(define_)`),
		),
		rule.Always(
			rule.MatchingLanguages("Markdown"),
		),
	},
	".mdoc": &Heuristics{
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dd +(?:[^"\s]+|"[^"]+")`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Dt +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*Sh +(?:[^"\s]|"[^"]+")`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Roff Manpage"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*TH +(?:[^"\s]+|"[^"]+") +"?(?:[1-9]|@[^\s@]+@)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[.'][ \t]*SH +(?:[^"\s]+|"[^"\s]+)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("Roff"),
		),
	},
	".ml": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("OCaml"),
			regexp.MustCompile(`(?m)(^\s*module)|let rec |match\s+(\S+\s)+with`),
		),
		rule.Or(
			rule.MatchingLanguages("Standard ML"),
			regexp.MustCompile(`(?m)=> |case\s+(\S+\s)+of`),
		),
	},
	".mod": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)<!ENTITY `),
		),
		rule.Or(
			rule.MatchingLanguages("Modula-2"),
			regexp.MustCompile(`(?m)^\s*(?i:MODULE|END) [\w\.]+;`),
		),
		rule.Always(
			rule.MatchingLanguages("Linux Kernel Module", "AMPL"),
		),
	},
	".ms": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^[.'][A-Za-z]{2}(\s|$)`),
		),
		rule.And(
			rule.MatchingLanguages("Unix Assembly"),
			rule.Not(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)/\*`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^\s*\.(?:include\s|globa?l\s|[A-Za-z][_A-Za-z0-9]*:)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("MAXScript"),
		),
	},
	".n": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^[.']`),
		),
		rule.Or(
			rule.MatchingLanguages("Nemerle"),
			regexp.MustCompile(`(?m)^(module|namespace|using)\s`),
		),
	},
	".ncl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)^\s*<\?xml\s+version`),
		),
		rule.Or(
			rule.MatchingLanguages("Text"),
			regexp.MustCompile(`(?m)THE_TITLE`),
		),
	},
	".nl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("NL"),
			regexp.MustCompile(`(?m)^(b|g)[0-9]+ `),
		),
		rule.Always(
			rule.MatchingLanguages("NewLisp"),
		),
	},
	".odin": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Object Data Instance Notation"),
			regexp.MustCompile(`(?m)(?:^|<)\s*[A-Za-z0-9_]+\s*=\s*<`),
		),
		rule.Or(
			rule.MatchingLanguages("Odin"),
			regexp.MustCompile(`(?m)package\s+\w+|\b(?:im|ex)port\s*"[\w:./]+"|\w+\s*::\s*(?:proc|struct)\s*\(|^\s*//\s`),
		),
	},
	".p": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Gnuplot"),
			regexp.MustCompile(`(?m)^s?plot\b|^set\s+(term|terminal|out|output|[xy]tics|[xy]label|[xy]range|style)\b`),
		),
		rule.Always(
			rule.MatchingLanguages("OpenEdge ABL"),
		),
	},
	".php": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Hack"),
			regexp.MustCompile(`(?m)<\?hh`),
		),
		rule.Or(
			rule.MatchingLanguages("PHP"),
			regexp.MustCompile(`(?m)<\?[^h]`),
		),
	},
	".pl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Prolog"),
			regexp.MustCompile(`(?m)^[^#]*:-`),
		),
		rule.Or(
			rule.MatchingLanguages("Perl"),
			regexp.MustCompile(`(?m)\buse\s+(?:strict\b|v?5\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Raku"),
			regexp.MustCompile(`(?m)^\s*(?:use\s+v6\b|\bmodule\b|\b(?:my\s+)?class\b)`),
		),
	},
	".plist": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML Property List"),
			regexp.MustCompile(`(?m)<!DOCTYPE\s+plist`),
		),
		rule.Always(
			rule.MatchingLanguages("OpenStep Property List"),
		),
	},
	".pm": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Perl"),
			regexp.MustCompile(`(?m)\buse\s+(?:strict\b|v?5\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Raku"),
			regexp.MustCompile(`(?m)^\s*(?:use\s+v6\b|\bmodule\b|\b(?:my\s+)?class\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("X PixMap"),
			regexp.MustCompile(`(?m)^\s*\/\* XPM \*\/`),
		),
	},
	".pod": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Pod 6"),
			regexp.MustCompile(`(?m)^[\s&&[^\n]]*=(comment|begin pod|begin para|item\d+)`),
		),
		rule.Always(
			rule.MatchingLanguages("Pod"),
		),
	},
	".pp": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Pascal"),
			regexp.MustCompile(`(?m)^\s*end[.;]`),
		),
		rule.Or(
			rule.MatchingLanguages("Puppet"),
			regexp.MustCompile(`(?m)^\s+\w+\s+=>\s`),
		),
	},
	".pro": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Proguard"),
			regexp.MustCompile(`(?m)^-(include\b.*\.pro$|keep\b|keepclassmembers\b|keepattributes\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("Prolog"),
			regexp.MustCompile(`(?m)^[^\[#]+:-`),
		),
		rule.Or(
			rule.MatchingLanguages("INI"),
			regexp.MustCompile(`(?m)last_client=`),
		),
		rule.And(
			rule.MatchingLanguages("QMake"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)HEADERS`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)SOURCES`),
			),
		),
		rule.Or(
			rule.MatchingLanguages("IDL"),
			regexp.MustCompile(`(?m)^\s*function[ \w,]+$`),
		),
	},
	".properties": &Heuristics{
		rule.And(
			rule.MatchingLanguages("INI"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[^#!;][^=]*=`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[;\[]`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Java Properties"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[^#!;][^=]*=`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[#!]`),
			),
		),
		rule.Or(
			rule.MatchingLanguages("INI"),
			regexp.MustCompile(`(?m)^[^#!;][^=]*=`),
		),
		rule.Or(
			rule.MatchingLanguages("Java properties"),
			regexp.MustCompile(`(?m)^[^#!][^:]*:`),
		),
	},
	".props": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)^(\s*)(?i:<Project|<Import|<Property|<\?xml|xmlns)`),
		),
		rule.Or(
			rule.MatchingLanguages("INI"),
			regexp.MustCompile(`(?m)(?i:\w+\s*=\s*)`),
		),
	},
	".q": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("q"),
			regexp.MustCompile(`(?m)((?i:[A-Z.][\w.]*:{)|(^|\n)\\(cd?|d|l|p|ts?) )`),
		),
		rule.Or(
			rule.MatchingLanguages("HiveQL"),
			regexp.MustCompile(`(?m)(?i:SELECT\s+[\w*,]+\s+FROM|(CREATE|ALTER|DROP)\s(DATABASE|SCHEMA|TABLE))`),
		),
	},
	".r": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Rebol"),
			regexp.MustCompile(`(?m)(?i:\bRebol\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("R"),
			regexp.MustCompile(`(?m)<-|^\s*#`),
		),
	},
	".rno": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^\.\\" `),
		),
	},
	".rpy": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Python"),
			regexp.MustCompile(`(?m)(?m:^(import|from|class|def)\s)`),
		),
		rule.Always(
			rule.MatchingLanguages("Ren'Py"),
		),
	},
	".rs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Rust"),
			regexp.MustCompile(`(?m)^(use |fn |mod |pub |macro_rules|impl|#!?\[)`),
		),
		rule.Or(
			rule.MatchingLanguages("RenderScript"),
			regexp.MustCompile(`(?m)#include|#pragma\s+(rs|version)|__attribute__`),
		),
	},
	".s": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Motorola 68K Assembly"),
			regexp.MustCompile(`(?m)(?im)\bmoveq(?:\.l)?\s+#(?:\$-?[0-9a-f]{1,3}|%[0-1]{1,8}|-?[0-9]{1,3}),\s*d[0-7]\b|(?im)^\s*move(?:\.[bwl])?\s+(?:sr|usp),\s*[^\s]+|(?im)^\s*move\.[bwl]\s+.*\b[ad]\d|(?im)^\s*movem\.[bwl]\b|(?im)^\s*move[mp](?:\.[wl])?\b|(?im)^\s*btst\b|(?im)^\s*dbra\b`),
		),
	},
	".sc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("SuperCollider"),
			regexp.MustCompile(`(?m)(?i:\^(this|super)\.|^\s*~\w+\s*=\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Scala"),
			regexp.MustCompile(`(?m)(^\s*import (scala|java)\.|^\s*class\b)`),
		),
	},
	".sql": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("PLpgSQL"),
			regexp.MustCompile(`(?m)(?i:^\\i\b|AS \$\$|LANGUAGE '?plpgsql'?|SECURITY (DEFINER|INVOKER)|BEGIN( WORK )?;)`),
		),
		rule.Or(
			rule.MatchingLanguages("SQLPL"),
			regexp.MustCompile(`(?m)(?i:(alter module)|(language sql)|(begin( NOT)+ atomic)|signal SQLSTATE '[0-9]+')`),
		),
		rule.Or(
			rule.MatchingLanguages("PLSQL"),
			regexp.MustCompile(`(?m)(?i:\$\$PLSQL_|XMLTYPE|sysdate|systimestamp|\.nextval|connect by|AUTHID (DEFINER|CURRENT_USER)|constructor\W+function)`),
		),
		rule.And(
			rule.MatchingLanguages("TSQL"),
			rule.Not(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)(?i:IDENTIFIED|NUMBER|VARCHAR2|REPEAT|UNTIL|IMMEDIATE)`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)(?i:(GO)|(@@)|(CREATE PROCEDURE)|BEGIN( TRY| CATCH)|OUTPUT( INSERTED)|IF|ELSE|IIF|CHOOSE|CURSOR|FETCH|DEALLOCATE|DECLARE)`),
			),
		),
		rule.Not(
			rule.MatchingLanguages("SQL"),
			regexp.MustCompile(`(?m)(?i:begin|boolean|package|exception)`),
		),
	},
	".srt": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("SubRip Text"),
			regexp.MustCompile(`(?m)^(\d{2}:\d{2}:\d{2},\d{3})\s*(-->)\s*(\d{2}:\d{2}:\d{2},\d{3})$`),
		),
	},
	".t": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Perl"),
			regexp.MustCompile(`(?m)\buse\s+(?:strict\b|v?5\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Raku"),
			regexp.MustCompile(`(?m)^\s*(?:use\s+v6\b|\bmodule\b|\b(?:my\s+)?class\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("Turing"),
			regexp.MustCompile(`(?m)^\s*%[ \t]+|^\s*var\s+\w+(\s*:\s*\w+)?\s*:=\s*\w+`),
		),
	},
	".toc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("World of Warcraft Addon Data"),
			regexp.MustCompile(`(?m)^## |@no-lib-strip@`),
		),
		rule.Or(
			rule.MatchingLanguages("TeX"),
			regexp.MustCompile(`(?m)^\\(contentsline|defcounter|beamer|boolfalse)`),
		),
	},
	".ts": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)<TS\b`),
		),
		rule.Always(
			rule.MatchingLanguages("TypeScript"),
		),
	},
	".tst": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("GAP"),
			regexp.MustCompile(`(?m)gap> `),
		),
		rule.Always(
			rule.MatchingLanguages("Scilab"),
		),
	},
	".tsx": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("TSX"),
			regexp.MustCompile(`(?m)^\s*(import.+(from\s+|require\()['"]react|\/\/\/\s*<reference\s)`),
		),
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)(?i:^\s*<\?xml\s+version)`),
		),
	},
	".v": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Coq"),
			regexp.MustCompile(`(?m)\(\*.*?\*\)|(?:^|\s)(?:Proof|Qed)\.(?:$|\s)|(?:^|\s)Require[ \t]+Import\s`),
		),
		rule.Or(
			rule.MatchingLanguages("Verilog"),
			regexp.MustCompile(`(?m)^[ \t]*module\s+[^\s()]+\s+\#?\(|^[ \t]*`+"`"+`(?:ifdef|timescale)\s|^[ \t]*always[ \t]+@`),
		),
		rule.Or(
			rule.MatchingLanguages("V"),
			regexp.MustCompile(`(?m)\$(?:if|else)[ \t]|^[ \t]*fn\s+[^\s()]+\(.*?\).*?\{|^[ \t]*for\s*\{`),
		),
	},
	".vba": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Vim script"),
			regexp.MustCompile(`(?m)^UseVimball`),
		),
		rule.Always(
			rule.MatchingLanguages("VBA"),
		),
	},
	".w": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("OpenEdge ABL"),
			regexp.MustCompile(`(?m)&ANALYZE-SUSPEND _UIB-CODE-BLOCK _CUSTOM _DEFINITIONS`),
		),
		rule.Or(
			rule.MatchingLanguages("CWeb"),
			regexp.MustCompile(`(?m)^@(<|\w+\.)`),
		),
	},
	".x": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("DirectX 3D File"),
			regexp.MustCompile(`(?m)^xof 030(2|3)(?:txt|bin|tzip|bzip)\b`),
		),
		rule.Or(
			rule.MatchingLanguages("RPC"),
			regexp.MustCompile(`(?m)\b(program|version)\s+\w+\s*{|\bunion\s+\w+\s+switch\s*\(`),
		),
		rule.Or(
			rule.MatchingLanguages("Logos"),
			regexp.MustCompile(`(?m)^%(end|ctor|hook|group)\b`),
		),
		rule.Or(
			rule.MatchingLanguages("Linker Script"),
			regexp.MustCompile(`(?m)OUTPUT_ARCH\(|OUTPUT_FORMAT\(|SECTIONS`),
		),
	},
	".yy": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("JSON"),
			regexp.MustCompile(`(?m)\"modelName\"\:\s*\"GM`),
		),
		rule.Always(
			rule.MatchingLanguages("Yacc"),
		),
	},
}

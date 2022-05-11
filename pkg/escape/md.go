package escape

import "strings"

// Markdown Escaper
var Markdown = strings.NewReplacer(
	`!`, `\!`,
	`*`, `\*`,
	`.`, `\.`,
	`-`, `\-`,
	`+`, `\+`,
	`=`, `\=`,
	`_`, `\_`,
	`~`, `\~`,
	"`", "\\`",
	`|`, `\|`,
	`[`, `\[`,
	`]`, `\]`,
	`(`, `\(`,
	`)`, `\)`,
	`<`, `\<`,
	`>`, `\>`,
	`{`, `\{`,
	`}`, `\}`,
	`|`, `\|`,
)

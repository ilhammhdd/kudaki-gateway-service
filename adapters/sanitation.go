package adapters

import "strings"

type RedisearchUnsanitary string

func (ru RedisearchUnsanitary) Sanitize() string {
	replacer := strings.NewReplacer(`,`, `\,`, `.`, `\.`, `<`, `\<`, `>`, `\>`, `{`, `\{`, `}`, `\}`, `[`, `\[`, `]`, `\]`, `"`, `\"`, `'`, `\'`, `:`, `\:`, `;`, `\;`, `!`, `\!`, `@`, `\@`, `#`, `\#`, `$`, `\$`, `%`, `\%`, `^`, `\^`, `&`, `\&`, `*`, `\*`, `(`, `\(`, `)`, `\)`, `-`, `\-`, `+`, `\+`, `=`, `\=`, `~`, `\~`)
	return replacer.Replace(string(ru))
}

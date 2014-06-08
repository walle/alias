// Package alias adds alias/redirection to Martini apps using regular expressions to strings. Like mod_rewrite.
package alias

import (
	"net/http"
	"regexp"

	"github.com/go-martini/martini"
)

// Type determines what kind of alias is used
type Type int8

// The different types
const (
	SERVE = iota
	REDIRECT
)

// Alias defines an alias
type Alias struct {
	from      *regexp.Regexp
	to        string
	aliasType Type
}

// A map of all aliases defined
var aliases map[string]*Alias

// Add is used to add a alias. From must be a valid regexp.
func Add(from, to string, aliasType Type) {
	if aliases == nil {
		aliases = make(map[string]*Alias)
	}
	re := regexp.MustCompile(from)
	a := &Alias{from: re, to: to, aliasType: aliasType}
	aliases[from] = a
}

// Remove stops serving a alias
func Remove(from string) {
	if aliases == nil {
		aliases = make(map[string]*Alias)
	}
	delete(aliases, from)
}

// Handler returns the martini handler that can be added to Use()
// Starts the alias functionality
func Handler() martini.Handler {
	if aliases == nil {
		aliases = make(map[string]*Alias)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, alias := range aliases {
			if alias.from.MatchString(r.URL.Path) {
				path := alias.from.ReplaceAllString(r.URL.Path, alias.to)

				if alias.aliasType == SERVE {
					r.URL.Path = path
				} else if alias.aliasType == REDIRECT {
					http.Redirect(w, r, path, 302)
				}

				break
			}
		}
	}
}

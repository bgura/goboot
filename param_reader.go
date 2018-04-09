package goboot

import (
	"strconv"
	"net/http"
)

// Typing for paramType enum
type ParamType int

// Parameter types of which the parameter read can parse
const (
	Unknown		ParamType = iota
	Uint64
	String 
)

type paramReader struct {
	// The request from which to extract parameters
	context *http.Request
	// default parameter values if a value is optional
	defaults map[string]string
	// Mapping of field name to its error
	errors map[string]string
}

// Initialize the ParamReader with a specific http request. This serves
// as the 'context' of our param reader. All subsequent calls will validate
// the params that are present on this assigned http.Request
func (p *paramReader) Context(r *http.Request) {
	p.context  = r
	p.defaults = make(map[string]string)
	p.errors   = make(map[string]string)
}

// Mark a parameter as optional and define its default value
func (p *paramReader) Optional(s string,  d string, t ParamType) {
	// Was a value provided?
	if v, ok := p.context.URL.Query()[s]; ok {
		// A value was provided, make sure it is valid
		if len(v) < 1 {
			p.AddError(s, "Invalid length of parameter value")
			return
		}

		// Check that it is valid
		if !p.isValid(v[0], t) {
			p.AddError(s, "Invalid value for type")
			return
		}
	} else {
		// Is the optional value valid?
		if !p.isValid(d, t) {
			panic("Attempted to add default value which is invalid for type")
		}

		// Add this as our default value
		if _, ok := p.defaults[s]; ok {
			panic("Attempted to add a default value for optional param when one already exists!")
		} else {
			p.defaults[s] = d
		}
	}
}

// Validate that a given param 's' is both present and a valid
// value of type 't'. A value is demeed valid if a conversion from 
// its string representation to 't' is possible
func(p *paramReader) Require(s string, t ParamType) {
	v, ok := p.context.URL.Query()[s]

	// Check that the value was present and that it is not empty
	if !ok || len(v) < 1 {
		p.AddError(s, "Required parameter not found")
		return
	}

	// Check that the value is valid for the type
	if !p.isValid(v[0], t) {
		p.AddError(s, "Invalid value for type")
		return
	}
}

// Get the parameter as a Uint64
func (p *paramReader) ReadUint64(s string) uint64 {
	if !p.HasParameter(s) {
		// Check for default
	
		if v, ok := p.defaults[s]; !ok {
			panic("Parameter not found and no default defined")
		} else if i, err := strconv.ParseUint(v, 10, 64); err != nil {
			panic("Failed to convert optional parameter to Uint64")
		} else {
			return i
		}
	} 
	
	if v, ok := p.context.URL.Query()[s]; ok {
		// Try to convert the parameter
		if i, err := strconv.ParseUint(v[0], 10, 64); err != nil {
			panic("Failed to convert parameter to Uint64")
		} else {
			return i
		}
	} else {
		panic("Parameter not found and no default defined")
	}
}

// Get the parameter as a string
func (p *paramReader) ReadString(s string) string {
	if !p.HasParameter(s) {
		// Check for default
		if v, ok := p.defaults[s]; !ok {
			panic("Parameter not found and no default defined")
		} else {
			return v
		}
	}

	if v, ok := p.context.URL.Query()[s]; ok {
		return v[0]
	} else {
		panic("Parameter not found and no default defined")
	}
}

// Add an error for the given parameter 's'. If an error is already set
// subsequent calls are effectively a no-op
func(p *paramReader) AddError(s string, e string) {
	if _, ok := p.errors[s]; !ok {
		return;
	}
	p.errors[s] = e
}

// Checks if we have had any errors when trying to read our parameters
func(p *paramReader) HasErrors() bool {
	return len(p.errors) != 0
}

// Checks that the specified string is valid using the given type
func(p *paramReader) isValid(s string, t ParamType) bool {
	switch t {
		case Uint64:
			_, err := strconv.ParseUint(s, 10, 64)
			return err == nil
		case String:
			return true
		default:
			return false
	}
}

func(p *paramReader) HasParameter(s string) bool {
	_, ok := p.context.URL.Query()[s]
	return ok 
}
// Code generated by go-vue-ssr: https://github.com/zbysir/go-vue-ssr
// src_hash:79081af18a527e23c10e595f84d26e78

package tplgo

import (
	"strings"
)

type _ strings.Builder

func (r *Render) Component_select(options *Options) string {
	this := extendMap(r.Prototype, options.Props)
	_ = this
	return r.Tag("a", true, &Options{
		Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string {
			return r.Component_select(&Options{
				Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string {
					return "<ff></ff><xx-option> 133 </xx-option><option value=\"1\"></option><option value=\"2\"></option>"
				}},
				P:    options,
				Data: this,
			})
		}},
		P:    options,
		Data: this,
	})
}

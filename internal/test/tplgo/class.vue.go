// Code generated by go-vue-ssr: https://github.com/bysir-zl/go-vue-ssr
// src_hash:40ab2dce8e9cbafa118ca0d019b5e382

package tplgo

import (
	"strings"
)

type _ strings.Builder

func (r *Render) Component_class(options *Options) string {
	this := extendMap(r.Prototype, options.Props)
	_ = this
	return r.Tag("div", true, &Options{
		PropsClass: map[string]interface{}{"a": true},
		Class:      []string{"b"},
		Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string {
			return "<span" + mixinClass(nil, []string{"d"}, map[string]interface{}{"c": true}) + mixinAttr(nil, nil, map[string]interface{}{"a": 1}) + ">\n            test class\n        </span>"
		}},
		P:    options,
		Data: this,
	})
}

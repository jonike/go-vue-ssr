// Code generated by go-vue-ssr: https://github.com/bysir-zl/go-vue-ssr
// src_hash:2ccfdc8e40765c9b234be90c7f8970d1

package tplgo

import (
	"strings"
)

type _ strings.Builder

func (r *Render) Component_bench(options *Options) string {
	this := extendMap(r.Prototype, options.Props)
	_ = this
	return r.Tag("div", true, &Options{
		PropsClass: map[string]interface{}{"a": true},
		Class:      []string{"b"},
		Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string {
			return "<span" + mixinClass(nil, []string{"d"}, map[string]interface{}{"c": true}) + mixinAttr(nil, nil, map[string]interface{}{"a": 1}) + ">\n        " + interfaceToStr(lookInterface(this, "data", "msg"), true) + "\n    </span>" + func() string {
				var b strings.Builder

				for index, item := range interface2Slice(lookInterface(this, "data", "c")) {
					b.WriteString(func(xdata map[string]interface{}) string {
						this := extendMap(xdata, map[string]interface{}{
							"$index": index,
							"item":   item,
						})

						return "<div>" + r.Component_bench(&Options{
							Props: map[string]interface{}{"data": lookInterface(this, "item")},
							Slot:  map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
							P:     options,
							Data:  this,
						}) + "</div>"
					}(this))
				}
				return b.String()
			}()
		}},
		P:    options,
		Data: this,
	})
}

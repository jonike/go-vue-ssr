// Code generated by go-vue-ssr: https://github.com/zbysir/go-vue-ssr
// src_hash:b80af1e00c02e47313bb4a0ed33bb6e4

package bench_string

import "bytes"

func (r *Render) Component_bench(options *Options) string {
	this := extendMap(r.Prototype, options.Props)
	_ = this
	return r.Tag("div", true, &Options{
		PropsClass: map[string]interface{}{"a": true},
		Class:      []string{"b"},
		Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string {
			return "<span" + mixinClass(nil, []string{"d"}, map[string]interface{}{"c": true}) + mixinAttr(nil, nil, map[string]interface{}{"a": "1"}) + ">\n        " + interfaceToStr(lookInterface(this, "data", "msg"), true) + "\n    </span>" + func() string {
				var b bytes.Buffer

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

// Code generated by go-vue-ssr: https://github.com/zbysir/go-vue-ssr
// src_hash:40b8bcafff2ff9a1b3171e190ed85361

package tplgo

import (
	"strings"
)

type _ strings.Builder

func (r *Render) Component_helloworld(options *Options) string {
	this := extendMap(r.Prototype, options.Props)
	_ = this
	return func() string {
		if interfaceToBool(lookInterface(this, "isShow")) {
			return r.Tag("div", true, &Options{
				Class: []string{"VueToNuxtLogo"},
				Slot: map[string]NamedSlotFunc{"abc": func(props map[string]interface{}) string {
					this := extendMap(this, map[string]interface{}{"a": props})
					_ = this
					return "<div>\n        2我是具名slot props msg: " + interfaceToStr(lookInterface(this, "a", "msg"), true) + "\n        2我是具名slot 所属组件属性 age: " + interfaceToStr(lookInterface(this, "age"), true) + "\n      </div>"
				}, "default": func(props map[string]interface{}) string {
					return "<div" + mixinClass(nil, []string{"Triangle", "Triangle--two"}, lookInterface(this, "customClass")) + " style=\"background: #f99; \">\n      我是一个DIV\n    </div>\n\n    name: " + interfaceToStr(interfaceAdd(interfaceAdd(lookInterface(this, "name"), " "), lookInterface(this, "name")), true) + "\n    info:\n\n    " + r.Component_text(&Options{
						Props: map[string]interface{}{"list": lookInterface(this, "list")},
						Slot: map[string]NamedSlotFunc{"abc": func(props map[string]interface{}) string {
							this := extendMap(this, map[string]interface{}{"a": props})
							_ = this
							return "<div>\n        1我是具名slot props msg: " + interfaceToStr(lookInterface(this, "a", "msg"), true) + "\n        1我是具名slot 所属组件属性 age: " + interfaceToStr(lookInterface(this, "age"), true) + "\n      </div>"
						}, "default": func(props map[string]interface{}) string { return "" }},
						P:    options,
						Data: this,
					}) + r.Component_text(&Options{
						Props: map[string]interface{}{"list": lookInterface(this, "list")},
						Slot: map[string]NamedSlotFunc{"abc": func(props map[string]interface{}) string {
							this := extendMap(this, map[string]interface{}{"a": props})
							_ = this
							return "<div>\n        2我是具名slot props msg: " + interfaceToStr(lookInterface(this, "a", "msg"), true) + "\n        2我是具名slot 所属组件属性 age: " + interfaceToStr(lookInterface(this, "age"), true) + "\n      </div>"
						}, "default": func(props map[string]interface{}) string { return "" }},
						P:    options,
						Data: this,
					}) + func() string {
						if interfaceToBool(!interfaceToBool(lookInterface(this, "isShow"))) {
							return "<div>显示隐藏</div>"
						}
						return ""
					}() + r.Component_class(&Options{
						PropsClass: lookInterface(this, "customClass"),
						Class:      []string{"d", "e", "f"},
						Slot:       map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
						P:          options,
						Data:       this,
					}) + r.Component_component(&Options{
						PropsStyle: map[string]interface{}{"padding": "50px"},
						Props:      map[string]interface{}{"is": "xstyle"},
						Style:      map[string]string{"margin": "50px"},
						Slot:       map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
						P:          options,
						Data:       this,
					}) + r.Component_component(&Options{
						Props: map[string]interface{}{"is": "class"},
						Slot:  map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
						P:     options,
						Data:  this,
					}) + r.Component_xStyle(&Options{
						Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
						P:    options,
						Data: this,
					}) + r.Component_xattr(&Options{
						Props: map[string]interface{}{"imgUrl": lookInterface(this, "imgUrl"), "src": lookInterface(this, "name")},
						Attrs: map[string]string{"id": "attr", "type": "input"},
						Slot:  map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
						P:     options,
						Data:  this,
					})
				}},
				P:    options,
				Data: this,
			})
		}
		return ""
	}()
}

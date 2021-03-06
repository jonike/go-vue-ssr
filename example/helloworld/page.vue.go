// Code generated by go-vue-ssr: https://github.com/zbysir/go-vue-ssr
// src_hash:822de0310a2edb01adb3b2dcdde8d1a9

package main

import (
	"strings"
)

type _ strings.Builder

func (r *Render) Component_page(options *Options) string {
	this := extendMap(r.Prototype, options.Props)
	_ = this
	return "<!doctype html><html lang=\"zh\"><head><meta charset=\"UTF-8\"></meta><title>" + interfaceToStr(lookInterface(this, "title"), true) + "</title></head><body><h1 style=\"text-align: center; margin-top: 100px; \">" + interfaceToStr(lookInterface(this, "title")) + "</h1>" +
		r.Component_info(&Options{
			Props: map[string]interface{}{"height": interfaceAdd(lookInterface(this, "height"), 1), "logo": lookInterface(this, "logo"), "name": lookInterface(this, "title"), "slogan": lookInterface(this, "slogan")},
			Style: map[string]string{"padding": "20px"},
			Slot:  map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
			P:     options,
			Data:  this,
		}) + "<div info> author: " + interfaceToStr(lookInterface(this, "info", "author"), true) + " " + interfaceToStr(lookInterface(this, "info", interfaceToStr(lookInterface(this, "slogan"))), true) + " " + interfaceToStr([]interface{}{lookInterface(this, "slogan")}, true) + "</div>" + r.Component_vOn(&Options{
		Props: map[string]interface{}{"msg": "hello event"},
		Slot:  map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
		P:     options,
		Data:  this,
	}) + r.tag("script", false, &Options{
		Slot: map[string]NamedSlotFunc{"default": func(props map[string]interface{}) string { return "" }},
		P:    options,
		Directives: []directive{
			{Name: "v-on-handler", Value: nil, Arg: ""},
		},
		Data: this,
	}) + "<script>\n    // 为dom添加事件\n    for (var i in vOnBinds) {\n        var item = vOnBinds[i];\n        var dom = document.querySelector('[data-von-' + item.DomSelector + ']')\n        dom.addEventListener(item.Event, function (item, dom) {\n            return function (event) {\n                if (window[item.Func]) {\n                    window[item.Func].call(window, event, ...item.Args)\n                } else {\n                    console.error('not found function: ' + item.Func)\n                }\n            }\n        }(item, dom))\n    }\n\n    function buttonClick(msg) {\n        alert(msg)\n    }\n</script></body></html>"
}

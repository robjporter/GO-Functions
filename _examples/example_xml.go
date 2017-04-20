package main

import (
	"fmt"

	xmlx "../xml"
)

func main() {
	data := `<aaaLogin response="yes" outCookie="0123456789abcdefghijklmnopqrstuvwxyz" outRefreshPeriod="600" outPriv="aaa,ext-lan-policy,ext-lan-qos,ext-san-policy,operations,pod-policy,pod-qos,read-only" outDomains="mgmt02-dummy" outChannel="noencssl" outEvtChannel="noencssl"></aaaLogin>`
	doc := xmlx.New()
	if err := doc.LoadString(data, nil); nil != err {
		fmt.Printf("LoadString(): %s\n", err)
	}
	login := doc.SelectNode("", "aaaLogin")
	fmt.Println(login)
	fmt.Println(login.As("", "outCookie"))
	fmt.Println(login.HasAttr("", "outCookie"))
	fmt.Println(login.HasAttr("", "outCookie2"))

	a := xmlx.NewNode(xmlx.NT_ROOT)
	a.Name = xmlx.NewXMLName("", "root")
	b := xmlx.NewNode(xmlx.NT_ELEMENT)
	b.Name = xmlx.NewXMLName("", "A")
	c := xmlx.NewNode(xmlx.NT_ELEMENT)
	c.Name = xmlx.NewXMLName("", "B")
	d := xmlx.NewNode(xmlx.NT_ELEMENT)
	d.Name = xmlx.NewXMLName("", "C")
	e := xmlx.NewNode(xmlx.NT_ELEMENT)
	e.Name = xmlx.NewXMLName("", "D")
	f := xmlx.NewNode(xmlx.NT_TEXT)
	f.Value = "xyzzy"
	g := xmlx.NewNode(xmlx.NT_TEXT)
	g.Value = "xyzzy"
	a.AddChild(b)
	b.AddChild(c)
	b.SetAttr("username", "USERNAME")
	c.AddChild(d)
	c.AddChild(e)
	d.AddChild(f)
	e.AddChild(g)
	fmt.Println(a)
}

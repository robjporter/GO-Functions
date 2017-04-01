package ucs

import (
	//"fmt"
	"github.com/robjporter/go-functions/as"
	"github.com/robjporter/go-functions/etree"
	"github.com/robjporter/go-functions/request"
	"strings"
)

const (
	XML_REPLACEMENT_START = "<"
	XML_REPLACEMENT_END   = ">"
	XML_LOGIN             = `<aaaLogin inName="<USERNAME>" inPassword="<PASSWORD>"/>`
	XML_LOGOUT            = `<aaaLogout inCookie="<COOKIE>"/>`
)

type UCSMDATA struct {
	version string
	cookie  string
	priv    string
}

type UCSMLogin struct {
	ip       string
	username string
	password string
}

type RESPONSE struct {
	Response string
	Body     string
	Errors   []error
}

type UCSM struct {
	handler      *request.SuperAgent
	cookie       string
	status       bool
	replacements map[string]string
	data         UCSMDATA
	login        UCSMLogin
	LastResponse RESPONSE
}

func New() *UCSM {
	u := UCSM{
		handler: request.New(),
		login: UCSMLogin{
			ip:       "",
			username: "",
			password: "",
		},
		cookie:       "",
		status:       false,
		replacements: make(map[string]string),
	}
	u.handler.SetInsecureDefaults()
	u.addReplacementDefaults()
	u.handler.SetRecorder(false)
	return &u
}

//PRIVATE*********************************************************************

func (u *UCSM) internalLogin() {
	var resp2, body2 string
	var err2 []error
	u.addReplacementString("USERNAME", u.login.username)
	u.addReplacementString("PASSWORD", u.login.password)
	xml := u.xmlReplace(XML_LOGIN)
	resp, body, err := u.handler.Post("https://"+u.login.ip+"/nuova").Set("Content-Type", "application/xml").Send(xml).End()
	if err == nil {
		if resp.StatusCode == 200 {
			if u.getCookieVersion(body) {
				u.status = true
				resp2 = as.ToString(resp)
				body2 = as.ToString(body)
			}
		}
	}
	err2 = err
	u.LastResponse.Response = resp2
	u.LastResponse.Body = body2
	u.LastResponse.Errors = err2
}

func (u *UCSM) getCookieVersion(xml string) bool {
	u.data.cookie = getXMLAttributeData(xml, "aaaLogin", "", "outCookie")
	u.data.version = getXMLAttributeData(xml, "aaaLogin", "", "outVersion")
	u.data.priv = getXMLAttributeData(xml, "aaaLogin", "", "outPriv")
	if u.data.cookie != "unknown" && u.data.version != "unknown" {
		u.addReplacementString("COOKIE", u.data.cookie)
		return true
	}
	return false
}

func (u *UCSM) xmlReplace(xml string) string {
	for k, v := range u.replacements {
		k = XML_REPLACEMENT_START + k + XML_REPLACEMENT_END
		xml = strings.Replace(xml, k, v, -1)
	}
	return xml
}

func (u *UCSM) addReplacementDefaults() {

}

func (u *UCSM) addReplacementString(key, value string) {
	u.replacements[key] = value
}

func getXMLAttributeData(xml string, root string, element string, attribute string) string {
	results := ""
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xml); err == nil {
		root := doc.SelectElement(root)
		if element == "" {
			return root.SelectAttrValue(attribute, "unknown")
		} else {
			element := root.SelectElement(element)
			if attribute == "" {
			} else {
				return element.SelectAttrValue(attribute, "unknown")
			}
		}
	}
	return results
}

//PUBLIC***********************************************************************

func (u *UCSM) Login(ip, username, password string) *UCSM {
	u.login.ip = ip
	u.login.username = username
	u.login.password = password
	u.internalLogin()
	return u
}

func (u *UCSM) Logout() []error {
	xml := u.xmlReplace(XML_LOGOUT)
	resp, _, err := u.handler.Post("https://"+u.login.ip+"/nuova").Set("Content-Type", "application/xml").Send(xml).End()
	if err == nil {
		if resp.StatusCode == 200 {
			u.status = false
			return nil
		}
	}
	return err
}

func (u *UCSM) End() (string, []error) {
	u.Logout()
	return "FINSIHED", nil
}

func (u *UCSM) GetPriviledges() string {
	if u.data.priv != "" && u.data.priv != "unknown" {
		return u.data.priv
	}
	return ""
}

func (u *UCSM) GetVersion() string {
	if u.data.version != "" && u.data.version != "unknown" {
		return u.data.version
	}
	return ""
}

func (u *UCSM) IsAdmin() bool {
	if u.data.priv != "" && u.data.priv != "unknown" {
		if strings.Contains(u.data.priv, "admin") {
			return true
		}
	}
	return false
}

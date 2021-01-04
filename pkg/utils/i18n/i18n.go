package i18n

import (
	"e.welights.net/devsecops/findv/configs"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io"
	"strings"
)

var p *message.Printer

func newMatcher(t []language.Tag) *matcher {
	tags := &matcher{make(map[language.Tag]int)}
	for i, tag := range t {
		ct, err := language.All.Canonicalize(tag)
		if err != nil {
			ct = tag
		}
		tags.index[ct] = i
	}
	return tags
}

type matcher struct {
	index map[language.Tag]int
}

type OptionsAgent struct {
	//config *configfile.ConfigFile
}

func NewOptionsAgent() *OptionsAgent {
	return &OptionsAgent{}
}

func (m matcher) Match(want ...language.Tag) (language.Tag, int, language.Confidence) {
	for _, t := range want {
		ct, err := language.All.Canonicalize(t)
		if err != nil {
			ct = t
		}
		conf := language.Exact
		for {
			if index, ok := m.index[ct]; ok {
				return ct, index, conf
			}
			if ct == language.Und {
				break
			}
			ct = ct.Parent()
			conf = language.High
		}
	}
	return language.Und, 0, language.No
}

var supported = newMatcher([]language.Tag{
	language.AmericanEnglish,
	language.English,
	language.SimplifiedChinese,
	language.Chinese,
})

func Init(lang language.Tag) {
	tag, _, _ := supported.Match(lang)
	switch tag {
	case language.AmericanEnglish, language.English:
		initEnUS(lang)
	case language.SimplifiedChinese, language.Chinese:
		initZhCN(lang)
	default:
		initZhCN(lang)
	}
}

func Fprintf(w io.Writer, key message.Reference, a ...interface{}) (n int, err error) {
	return p.Fprintf(w, key, a...)
}

func Printf(format string, a ...interface{}) {
	_, _ = p.Printf(format, a...)
}

func Sprintf(format string, a ...interface{}) string {
	return p.Sprintf(format, a...)
}

func Sprint(a ...interface{}) string {
	return p.Sprint(a...)
}

func init() {
	locale := configs.Language
	if strings.TrimSpace(locale) == "" {
		locale = "zh_CN"
	}
	Init(language.Make(locale + ".UTF-8"))
	p = message.NewPrinter(language.Make(locale + ".UTF-8"))
}

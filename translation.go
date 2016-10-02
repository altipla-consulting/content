package content

import (
	"encoding/json"

	"github.com/juju/errors"
)

type Translation struct {
	Content map[string]string
}

func NewTranslation() Translation {
	return Translation{}
}

func (t *Translation) init() {
	if t.Content == nil {
		t.Content = map[string]string{}
	}
}

func (t *Translation) Set(lang, content string) {
	t.init()
	t.Content[lang] = content
}

func (t *Translation) Get(lang string) string {
	t.init()
	return t.Content[lang]
}

func (t *Translation) LangChain(lang string) string {
	if t.Get(lang) != "" {
		return t.Get(lang)
	}

	if t.Get("en") != "" {
		return t.Get("en")
	}

	return t.Get("es")
}

func (t *Translation) Map() map[string]string {
	return t.Content
}

func (t Translation) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Map())
}

func (t *Translation) UnmarshalJSON(data []byte) error {
	m := map[string]string{}
	if err := json.Unmarshal(data, &m); err != nil {
		return errors.Trace(err)
	}

	*t = TranslationFromMap(m)

	return nil
}

func TranslationFromMap(m map[string]string) Translation {
	return Translation{Content: m}
}

func TestTranslations(suffix string) Translation {
	return TranslationFromMap(map[string]string{
		"es": "foo-" + suffix,
		"en": "bar-" + suffix,
	})
}

func CheckTranslations(suffix string, m map[string]string) bool {
	return m["es"] == "foo-"+suffix && m["en"] == "bar-"+suffix
}

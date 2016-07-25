package content

type Translation struct {
	Content map[string]string
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

func TranslationFromMap(m map[string]string) *Translation {
	return &Translation{Content: m}
}

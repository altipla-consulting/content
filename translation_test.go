package content

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"encoding/json"
	"strings"
)

var _ = Describe("Translation", func() {
	It("Should set a translation", func() {
		t := TranslationFromMap(map[string]string{"es": "foo"})
		t.Set("en", "bar")

		Expect(t.Get("es")).To(Equal("foo"))
		Expect(t.Get("en")).To(Equal("bar"))
	})

	It("Should set a translation from a nil map", func() {
		t := TranslationFromMap(nil)
		t.Set("en", "bar")

		Expect(t.Get("en")).To(Equal("bar"))
	})

	It("Should return the translations as a map", func() {
		t := TranslationFromMap(map[string]string{"es": "foo"})
		t.Set("en", "bar")

		m := t.Map()
		Expect(m["es"]).To(Equal("foo"))
		Expect(m["en"]).To(Equal("bar"))
	})

	It("Should chain the original lang first always", func() {
		t := TranslationFromMap(map[string]string{
			"es": "foo",
			"en": "bar",
			"it": "baz",
		})

		Expect(t.LangChain("es")).To(Equal("foo"))
		Expect(t.LangChain("en")).To(Equal("bar"))
		Expect(t.LangChain("it")).To(Equal("baz"))
	})

	It("Should chain english above spanish", func() {
		t := TranslationFromMap(map[string]string{
			"es": "foo",
			"en": "bar",
			"it": "baz",
		})

		Expect(t.LangChain("de")).To(Equal("bar"))
	})

	It("Should chain spanish as backup", func() {
		t := TranslationFromMap(map[string]string{
			"es": "foo",
			"it": "baz",
		})

		Expect(t.LangChain("de")).To(Equal("foo"))
	})

	// In order to have a consistent behaviour and implementation. It will avoid
	// runtime errors where users set content inadvertently.
	It("Should use english if spanish is not defined", func() {
		t := TranslationFromMap(map[string]string{
			"en": "bar",
			"it": "baz",
		})

		Expect(t.LangChain("en")).To(Equal("bar"))
	})

	It("Should encode itself to JSON as a simple map", func() {
		t := TranslationFromMap(map[string]string{
			"en": "bar",
			"it": "baz",
		})

		buf := bytes.NewBuffer(nil)
		err := json.NewEncoder(buf).Encode(t)
		Expect(err).To(Succeed())

		Expect(buf.String()).To(MatchJSON(`{"en": "bar", "it": "baz"}`))
	})

	It("Should decode itself from JSON as a simple map", func() {
		buf := strings.NewReader(`{"en": "bar", "it": "baz"}`)
		t := NewTranslation()
		err := json.NewDecoder(buf).Decode(&t)
		Expect(err).To(Succeed())

		Expect(t.Get("en")).To(Equal("bar"))
		Expect(t.Get("it")).To(Equal("baz"))
	})
})

package languages

type Language string

const (
	RUSSIAN Language = "ru"
	ENGLISH Language = "en"
)

var languages = map[string]Language{
	"ru": RUSSIAN,
	"en": ENGLISH,
}

func GetLang(lang string) Language {
	v, e := languages[lang]
	if !e {
		return ENGLISH
	}
	return v
}

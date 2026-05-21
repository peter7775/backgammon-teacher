package ports

type TranslationService interface {
	Translate(text, lang string) (string, error)
}

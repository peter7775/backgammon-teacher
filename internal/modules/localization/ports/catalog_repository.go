package ports

type CatalogRepository interface {
	Lookup(key, lang string) (string, error)
}

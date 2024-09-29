package checkout

type Cloner interface {
	Clone(path string, repoURL string, branch string) error
}

type Checkout struct {
	Cloner Cloner
}

func (g Checkout) Checkout(url string, branch string) (string, error) {
	path := DetermineFolderFromURL(url)
	err := g.Cloner.Clone(path, url, branch)
	if err != nil {
		return "", err
	}
	return path, err
}

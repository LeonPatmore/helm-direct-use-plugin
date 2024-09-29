package directuse

import (
	"io"
	"log"
	"path/filepath"
)

type Configuration struct {
	Out             io.Writer
	CheckoutService CheckoutService
	Updater         Updater
}

type Updater interface {
	Update(path string) error
}

type CheckoutService interface {
	Checkout(url string, branch string) (string, error)
}

func InstallChart(url string, subPath string, branch string, c Configuration) error {
	repoPath, err := c.CheckoutService.Checkout(url, branch)
	if err != nil {
		return err
	}
	chartFullPath := filepath.Join(repoPath, subPath)
	log.Printf("Chart path is %s", chartFullPath)
	err = c.Updater.Update(chartFullPath)
	if err != nil {
		return err
	}
	return nil
}

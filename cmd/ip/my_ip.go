package ip

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sourcegraph/conc/pool"
)

func askIp(site string) (string, error) {
	resp, err := http.Get(site)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(bytes), "\n"), err
}

var sites = []string{
	"https://ifconfig.co",
	"https://ifconfig.me",
}

func run() error {

	var chRes chan string = make(chan string, len(sites))

	defer func() {
		close(chRes)
		for res := range chRes {
			fmt.Fprint(os.Stdout, res)
		}
	}()

	p := pool.New().WithMaxGoroutines(len(sites)).WithErrors()
	for _, site := range sites {
		site := site
		p.Go(func() error {
			ip, err := askIp(site)
			if err != nil {
				return err
			}
			chRes <- fmt.Sprintf("%s by %s\n", ip, site)
			return nil
		})
	}
	return p.Wait()
}

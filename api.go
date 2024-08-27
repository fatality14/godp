package godp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// PackageInfo представляет информацию о пакете из API
type PackageInfo struct {
	Name      string `json:"name"`
	Epoch     int    `json:"epoch"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Arch      string `json:"arch"`
	Disttag   string `json:"disttag"`
	Buildtime int64  `json:"buildtime"`
	Source    string `json:"source"`
}

// APIResponse представляет полный ответ от API
type APIResponse struct {
	RequestArgs map[string]interface{} `json:"request_args"`
	Length      int                    `json:"length"`
	Packages    []PackageInfo          `json:"packages"`
}

// FetchPackagesFromAPI получает пакеты из API для указанной ветки
func FetchPackagesFromAPI[T any](route, branch string) (T, error) {
	var empty T
	url := fmt.Sprintf("%s/%s", route, branch)
	resp, err := http.Get(url)
	if err != nil {
		return empty, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return empty, err
	}

	var res T
	err = json.Unmarshal(body, &res)
	if err != nil {
		return empty, err
	}

	return res, nil
}

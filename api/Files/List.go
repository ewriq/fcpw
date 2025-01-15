package Files

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-json"
)


func ListItemToJSON(path string) (string, error) {
    var fileInfos []FileInfo

    err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
		
        fileInfo := FileInfo{
            Name:        info.Name(),
            IsDirectory: info.IsDir(),
            Size:        info.Size(),
            Path:        strings.ReplaceAll(p, "\\", "/"), 
            ModTime:     info.ModTime().Format(time.RFC3339), 
            IsHidden:    strings.HasPrefix(info.Name(), "."), 
            Extension:   filepath.Ext(info.Name()),
        }

        fileInfos = append(fileInfos, fileInfo)

        return nil
    })

    if err != nil {
        return "", err
    }

    jsonData, err := json.MarshalIndent(fileInfos, "", "  ")
    if err != nil {
        return "", err
    }

    return string(jsonData), nil
}

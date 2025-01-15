package Files

type FileInfo struct {
    Name           string `json:"name"`
    IsDirectory    bool   `json:"isDirectory"`
    Size           int64  `json:"size"`
    Path           string `json:"path"`
    ModTime        string `json:"modTime"` 
    IsHidden       bool   `json:"isHidden"` 
    Extension      string `json:"extension"` 
}
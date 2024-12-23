package beetlejuice

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "os/user"
  "path"
)

const Version = "1.0.0"

type BeetleJuice struct {
  FilePath string
}

const StoreFileName = ".bj"

func NewBeetleJuice() *BeetleJuice {
  return &BeetleJuice{}
}

func (b *BeetleJuice) WithFilePath(path string) *BeetleJuice {
  b.FilePath = path
  return b
}

func homePath() (string, error) {
  u, err := user.Current()
  if err != nil {
    return "", err
  }
  p := path.Join(u.HomeDir, StoreFileName)
  return p, nil
}

func (b *BeetleJuice) Lookup(name string) (string, error) {
  if len(b.FilePath) == 0 {
    path, err := homePath()
    if err != nil {
      return "", err
    }
    b.FilePath = path
  }

  dict, err := readYAMLFile(b.FilePath)
  if err != nil {
    return "", err
  }
  pass, ok := dict[name]
  if !ok {
    return "", fmt.Errorf("No entry found for %s", name)
  }

  return pass, nil
}

func readYAMLFile(path string) (map[string]string, error) {
  data := make(map[string]string)

  raw, err := ioutil.ReadFile(path)
  if err != nil {
    return data, err
  }

  err = yaml.Unmarshal(raw, &data)
  if err != nil {
    return data, err
  }

  return data, nil
}

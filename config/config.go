package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

// Conf is global config variable
var Conf Config

// Config is a struct of config
type Config struct {
	General GeneralConfig `toml:"General"`
	Gist    GistConfig    `toml:"Gist"`
	GitLab  GitLabConfig  `toml:"GitLab"`
	GitHub  GithubConfig  `toml:"Github"`
}

// GeneralConfig is a struct of general config
type GeneralConfig struct {
	SnippetFile string   `toml:"snippetfile"`
	Editor      string   `toml:"editor"`
	Column      int      `toml:"column"`
	SelectCmd   string   `toml:"selectcmd"`
	Backend     string   `toml:"backend"`
	SortBy      string   `toml:"sortby"`
	Cmd         []string `toml:"cmd"`
}

// GistConfig is a struct of config for Gist
type GistConfig struct {
	FileName    string `toml:"file_name"`
	AccessToken string `toml:"access_token"`
	GistID      string `toml:"gist_id"`
	Public      bool   `toml:"public"`
	AutoSync    bool   `toml:"auto_sync"`
}

// GitLabConfig is a struct of config for GitLabSnippet
type GitLabConfig struct {
	FileName    string `toml:"file_name"`
	AccessToken string `toml:"access_token"`
	Url         string `toml:"url"`
	ID          string `toml:"id"`
	Visibility  string `toml:"visibility"`
	AutoSync    bool   `toml:"auto_sync"`
	Insecure    bool   `toml:"skip_ssl"`
}

type GithubConfig struct {
	FileName    string `toml:"file_name"`
	AccessToken string `toml:"access_token"`
	RepoOwner   string `toml:"repo_owner"`
	RepoName    string `toml:"repo_name"`
	AutoSync    bool   `toml:"auto_sync"`
}

// Flag is global flag variable
var Flag FlagConfig

// FlagConfig is a struct of flag
type FlagConfig struct {
	Debug     bool
	Query     string
	FilterTag string
	Command   bool
	Delimiter string
	OneLine   bool
	Color     bool
	Tag       bool
	Output    bool
}

// Load loads a config toml
func (cfg *Config) Load(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		_, err := toml.DecodeFile(file, cfg)
		if err != nil {
			return err
		}
		cfg.General.SnippetFile = expandPath(cfg.General.SnippetFile)
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	cfg.General.SnippetFile = GetGithubSnippetsFilePath()

	cfg.General.Editor = os.Getenv("EDITOR")
	if cfg.General.Editor == "" && runtime.GOOS != "windows" {
		if isCommandAvailable("sensible-editor") {
			cfg.General.Editor = "sensible-editor"
		} else {
			cfg.General.Editor = "vim"
		}
	}

	// setup default config for hosted grafana - this shouldn't need to be changed by the user besides the access token
	cfg.General.Column = 40
	cfg.General.SelectCmd = "fzf --reverse --ansi"
	cfg.General.Backend = "github"

	cfg.GitHub.FileName = "snippet.toml"
	cfg.GitHub.AccessToken = "YOUR-ACCESS-TOKEN"
	cfg.GitHub.RepoOwner = "grafana"
	cfg.GitHub.RepoName = "hg-snippets-config"

	return toml.NewEncoder(f).Encode(cfg)
}

func GetGithubSnippetsFilePath() string {
	// assuming this command was run in the hg-snippets dir. This should only happen once - on setup
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(wd, "hg-snippets-config", "snippet.toml")
}

// GetDefaultConfigDir returns the default config directory
func GetDefaultConfigDir() (dir string, err error) {
	if env, ok := os.LookupEnv("PET_CONFIG_DIR"); ok {
		dir = env
	} else if runtime.GOOS == "windows" {
		dir = os.Getenv("APPDATA")
		if dir == "" {
			dir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data", "pet")
		}
		dir = filepath.Join(dir, "pet")
	} else {
		dir = filepath.Join(os.Getenv("HOME"), ".config", "pet")
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", fmt.Errorf("cannot create directory: %v", err)
	}
	return dir, nil
}

func expandPath(s string) string {
	if len(s) >= 2 && s[0] == '~' && os.IsPathSeparator(s[1]) {
		if runtime.GOOS == "windows" {
			s = filepath.Join(os.Getenv("USERPROFILE"), s[2:])
		} else {
			s = filepath.Join(os.Getenv("HOME"), s[2:])
		}
	}
	return os.Expand(s, os.Getenv)
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

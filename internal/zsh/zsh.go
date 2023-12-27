package zsh

import (
	"ectl/pkg/util"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// 安装oh-my-zsh
func InstallOhMyZsh() error {
	err := InstallZsh()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	zshrc := homeDir + "/.zshrc"

	_, err = os.Stat(zshrc)
	if err == nil {
		logrus.Info("oh-my-zsh is already installed.")
	} else if os.IsNotExist(err) {
		logrus.Info("开始安装oh-my-zsh")
		// https://github.com/ohmyzsh/ohmyzsh/wiki
		_, err = util.GetOutput("sh", "-c", "$(curl -fsSL https://install.ohmyz.sh)")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%w", err)
	}
	
	err = InstallZshAutoSuggestions()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// 安装zsh
func InstallZsh() error {
	oSReleaseInfo, err := util.GetOSRelease()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	isExists, err := util.IsCommandExists("zsh")
	if isExists {
		logrus.Info("zsh is already installed.")
		return nil
	} else {
		logrus.Info(err, "开始安装zsh")

		osName := strings.ToLower(oSReleaseInfo.NAME)

		switch osName {
		case "centos":
			_, err := util.GetOutput("yum", "install", "-y", "zsh")
			if err != nil {
				return err
			}
		case "ubuntu":
			_, err := util.GetOutput("apt", "install", "-y", "zsh")
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// 安装zsh-autosuggestions
func InstallZshAutoSuggestions() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	zshrc := homeDir + "/.zshrc"
	content, err := os.ReadFile(zshrc)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fileContent := string(content)
	if strings.Contains(fileContent, "zsh-autosuggestions") {
		logrus.Info("plugin zsh-autosuggestions is already installed.")
		return nil
	} else {
		logrus.Info("开始安装zsh-autosuggestions插件")
		_, err := util.GetOutput("git", "clone", "https://github.com/zsh-users/zsh-autosuggestions", "${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions")
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		_, err = util.GetOutput("sed", "-i", "'s/plugins=(git)/plugins=(\n    git\n    zsh-autosuggestions\n)/g'", "~/.zshrc")
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		_, err = util.GetOutput("echo", "bindkey ',' autosuggest-accept\nexport TERM=xterm-256color\nZSH_AUTOSUGGEST_HIGHLIGHT_STYLE=\"fg=yellow,bg=bold\"", ">>", zshrc)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}
}

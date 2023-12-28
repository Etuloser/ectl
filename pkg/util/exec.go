package util

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 已兼容CentOS, Ubuntu
type OSReleaseInfo struct {
	NAME                   string `json:"NAME"`
	VERSION                string `json:"VERSION"`
	ID                     string `json:"ID"`
	ID_LIKE                string `json:"ID_LIKE"`
	VERSION_ID             string `json:"VERSION_ID"`
	VERSION_CODENAME       string `json:"VERSION_CODENAME"`
	PLATFORM_ID            string `json:"PLATFORM_ID"`
	PRETTY_NAME            string `json:"PRETTY_NAME"`
	ANSI_COLOR             string `json:"ANSI_COLOR"`
	LOGO                   string `json:"LOGO"`
	CPE_NAME               string `json:"CPE_NAME"`
	HOME_URL               string `json:"HOME_URL"`
	SUPPORT_URL            string `json:"SUPPORT_URL"`
	BUG_REPORT_URL         string `json:"BUG_REPORT_URL"`
	PRIVACY_POLICY_URL     string `json:"PRIVACY_POLICY_URL"`
	REDHAT_SUPPORT_PRODUCT string `json:"REDHAT_SUPPORT_PRODUCT"`
	UBUNTU_CODENAME        string `json:"UBUNTU_CODENAME"`
}

// 判断命令是否存在
func IsCommandExists(cmd string) (bool, error) {
	_, err := GetOutput("command", "-v", cmd)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// 获取系统版本
func GetOSRelease() (*OSReleaseInfo, error) {
	output, err := GetOutput("cat", "/etc/os-release")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	lines := strings.Split(output, "\n")
	info := make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.Trim(parts[1], "\"")
			info[key] = value
		}
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var oSReleaseInfo OSReleaseInfo
	json.Unmarshal(jsonData, &oSReleaseInfo)
	return &oSReleaseInfo, nil
}

// 获取系统命令输出, 返回标准输出和ExitCode
func GetOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return string(output), nil
}

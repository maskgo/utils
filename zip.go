package utils

import (
	"os"
	"os/exec"
)

// 打包文件夹 eg:path=/tmp tgz=/data/log/tmp.tar.gz
func PackPath(path, tgz string) error {
	hasPigz := true
	if _, err := exec.LookPath("pigz"); err != nil {
		// fmt.Println("建议:启用pigz压缩(利用多线程压缩,需要安装 yum -y install pigz)")
		hasPigz = false
	}

	if hasPigz {
		f, err := os.Create(tgz)
		if err != nil {
			return err
		}

		c1 := exec.Command("tar", "cf", "-", path, "--remove-files")
		c2 := exec.Command("pigz")

		c2.Stdin, _ = c1.StdoutPipe()
		c2.Stdout = f

		if err := c2.Start(); err != nil {
			return nil
		}

		if err := c1.Run(); err != nil {
			return err
		}

		return c2.Wait()
	}

	_, err := exec.Command("tar", "czvf", tgz, path).Output()

	return err
}

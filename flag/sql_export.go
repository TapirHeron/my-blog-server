package flag

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"server/global"
	"time"
)

func SQLExport() error {
	mysqlConfig := global.Config.Mysql
	timer := time.Now().Format("20060102")
	sqlPath := fmt.Sprintf("mysql_%s.sql", timer)
	cmd := exec.Command("docker",
		"exec", "mysql", "mysqldump",
		"-u "+mysqlConfig.Username, "-p "+mysqlConfig.Password, mysqlConfig.DBName)
	outFile, err := os.Create(sqlPath)
	if err != nil {
		return err
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			global.Log.Error("Failed to close file:", zap.Error(err))
		}
	}(outFile)

	cmd.Stdout = outFile
	return cmd.Run()
}

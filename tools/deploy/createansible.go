package deploy

import (
	"bytes"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
	"text/template"
)

var (
	CreateAnsibleCommand = cli.Command{
		Name:   "createansible",
		Action: createAnsible,
		Flags: []cli.Flag{
			OutputFlag,
			AnsibleDirFlag,
			RemoteAnsibleDirFlag,
			ForceFlag,
			LocalFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func createAnsible(ctx *cli.Context) error {
	output := ctx.String(OutputFlag.Name)
	name := ctx.String(AnsibleDirFlag.Name)
	ansibleDir := filepath.Join(output, name)
	if ctx.Bool(ForceFlag.Name) {
		os.RemoveAll(ansibleDir)
	}
	if err := os.Mkdir(ansibleDir, 0755); err != nil {
		log.Error("Create dir failed", "path", ansibleDir, "err", err)
		return err
	}
	inventoriesDir := filepath.Join(output, name, "inventories")
	if err := os.Mkdir(inventoriesDir, 0755); err != nil {
		log.Error("Create dir failed", "path", inventoriesDir, "err", err)
		return err
	}
	playbooksDir := filepath.Join(output, name, "playbooks")
	if err := os.Mkdir(playbooksDir, 0755); err != nil {
		log.Error("Create dir failed", "path", playbooksDir, "err", err)
		return err
	}
	file := filepath.Join(ansibleDir, "ansible.cfg")
	if err := os.WriteFile(file, []byte(fmt.Sprintf(ansibleCfg)), 0666); err != nil {
		log.Error("Write file failed", "path", file, "err", err)
	}
	buffer := new(bytes.Buffer)
	script := commandScript
	if ctx.Bool(LocalFlag.Name) {
		script = localCommandScript
	}
	tmpl := template.Must(template.New("").Parse(script))

	if err := tmpl.Execute(buffer, &ScriptParams{
		Dir: ctx.String(RemoteAnsibleDirFlag.Name),
	}); err != nil {
		log.Error("Execute template failed", "err", err)
		return err
	}
	file = filepath.Join(playbooksDir, "command.yml")
	if err := os.WriteFile(file, buffer.Bytes(), 0666); err != nil {
		log.Error("Write file failed", "path", file, "err", err)
	}
	buffer.Reset()
	script = deployScript
	if ctx.Bool(LocalFlag.Name) {
		script = localDeployScript
	}
	tmpl = template.Must(template.New("").Parse(script))

	if err := tmpl.Execute(buffer, &ScriptParams{
		Dir: ctx.String(RemoteAnsibleDirFlag.Name),
	}); err != nil {
		log.Error("Execute template failed", "err", err)
		return err
	}
	file = filepath.Join(playbooksDir, "deploy.yml")

	if err := os.WriteFile(file, buffer.Bytes(), 0666); err != nil {
		log.Error("Write file failed", "path", file, "err", err)
	}

	return nil
}

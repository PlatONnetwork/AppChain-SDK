package deploy

type HostConfig struct {
	All All `yaml:"all"`
}

type All struct {
	Hosts map[string]Host `yaml:"hosts"`
	Vars  Vars            `yaml:"vars"`
}
type Vars struct {
	AnsibleConnection    string `yaml:"ansible_connection"`
	AnsibleSSHCommonArgs string `yaml:"ansible_ssh_common_args"`
}
type Host struct {
	AnsibleHost     string `yaml:"ansible_host"`
	AnsibleUser     string `yaml:"ansible_user"`
	AnsiblePassword string `yaml:"ansible_password"`
}

type NodeFolder struct {
	Local   string `yaml:"local"`
	Archive string `yaml:"archive"`
}

type EnvConfig struct {
	NodeFolders []NodeFolder `yaml:"node_folders"`
}

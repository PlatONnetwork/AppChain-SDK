package deploy

type ScriptParams struct {
	Dir string
}

var commandScript = `---
- name: Execute script on remote servers
  hosts: all
  gather_facts: no
  vars_files: "vars/env.yml"

  tasks:
    - name: Ensure script exists on remote
      stat:
        path: {{.Dir}}/node/{{ "{{cmd}}" }}.sh
      register: script_stat

    - name: Execute command remote server
      when: inventory_hostname == item.local
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      command: "/bin/bash ./{{ "{{cmd}}" }}.sh"
      args:
        chdir: "{{.Dir}}/node"
      register: script_output

    - name: print result
      debug:
        msg: "{{ "{{item.stdout}}" }}"
      loop: "{{ "{{script_output.results}}" }}"
      when: item.stdout is defined`

var deployScript = `---
- name: Deploy node folders to remote servers
  hosts: all
  gather_facts: no
  vars_files: "vars/env.yml"

  tasks:
    - name: Create local archives (run locally)
      delegate_to: localhost
      run_once: yes
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      command: "tar -czf {{ "{{playbook_dir}}" }}/files/{{ "{{current_item.archive}}" }} -C {{ "{{playbook_dir}}" }}/files {{ "{{current_item.local}}" }}"
      changed_when: false

    - name: Ensure {{.Dir}} directory exists on remote
      file:
        path: {{.Dir}}
        state: directory
        mode: '0755'

    - name: Copy archive to remote server
      when: inventory_hostname == item.local
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      copy:
        src: "{{ "{{playbook_dir}}" }}/files/{{ "{{current_item.archive}}" }}"
        dest: "{{.Dir}}/{{ "{{current_item.archive}}" }}"
        mode: '0644'

    - name: Clean up remote dir
      when: inventory_hostname == item.local
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      file:
        path: "{{.Dir}}/node"
        state: absent

    - name: Extract archive on remote server
      when: inventory_hostname == item.local
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      command: "tar -xzf {{.Dir}}/{{ "{{current_item.archive}}" }} -C {{.Dir}}"


    - name: Rename dir on remote server
      when: inventory_hostname == item.local
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      command: "mv {{.Dir}}/{{ "{{current_item.local}}" }} {{.Dir}}/node"

    - name: Clean up remote archives
      when: inventory_hostname == item.local
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      file:
        path: "{{.Dir}}/{{ "{{current_item.archive}}" }}"
        state: absent

    - name: Clean up local archives (run locally)
      delegate_to: localhost
      run_once: yes
      loop: "{{ "{{node_folders}}" }}"
      vars:
        current_item: "{{ "{{item}}" }}"
      file:
        path: "{{ "{{current_item.archive}}" }}"
        state: absent
`
var ansibleCfg = `[defaults]
inventory      = ./inventories/hosts
playbook_dir   = ./playbooks
host_key_checking = false
[connection]
pipelining = False
[privilege_escalation]
become = False
become_ask_pass = False`

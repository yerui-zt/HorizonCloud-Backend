package pvectl

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

type ACLPermission struct {
	Path  string `json:"path,omitempty"`
	Roles string `json:"roles,omitempty"`
	Users string `json:"users,omitempty"`

	Delete    int    `json:"delete,omitempty" default:"0"`
	Groups    string `json:"groups,omitempty"`
	Propagate int    `json:"propagate,omitempty" default:"1"`
	Tokens    string `json:"tokens,omitempty"`
}

// CreateUser 创建用户,请传入 UserId = xxxx@pve
func (l *PVECtl) CreateUser(userId, password, email string) error {
	return l.Client.Post("/access/users", map[string]string{
		"userid":   userId,
		"password": password,
		"email":    email,
	}, nil)
}

func (l *PVECtl) DeleteUser(userId string) error {
	return l.Client.Delete(fmt.Sprintf("/access/users/%s", userId), nil)
}

func (l *PVECtl) SetUserPermission(userId string, v *proxmox.VirtualMachine) error {
	acl := &ACLPermission{
		Path:      fmt.Sprintf("/vms/%d", v.VMID),
		Roles:     "PVEVMUser",
		Users:     userId,
		Propagate: 1,
		Delete:    0,
	}
	return l.Client.Put("/access/acl", acl, nil)
}

func (l *PVECtl) DeleteUserPermission(userId string, v *proxmox.VirtualMachine) error {
	acl := &ACLPermission{
		Path:      fmt.Sprintf("/vms/%d", v.VMID),
		Roles:     "PVEVMUser",
		Users:     userId,
		Propagate: 0,
		Delete:    1,
	}
	return l.Client.Put("/access/acl", acl, nil)
}

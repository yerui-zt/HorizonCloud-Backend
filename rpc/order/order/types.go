package order

type OrderItemContentVMCreateContent struct {
	Plan          string `json:"plan"`
	VMGroup       string `json:"vm_group"`
	ServicePeriod string `json:"service_period"`
	OSImage       string `json:"os_image"`
}

type OrderItemActionVmInstanceCreateAction struct {
	HypervisorGroupId int64 `json:"hypervisor_group_id"`
	PlanID            int64 `json:"plan_id"`
	OSImageID         int64 `json:"os_image_id"`
}

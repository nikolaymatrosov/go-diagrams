package yc

import "github.com/blushft/go-diagrams/diagram"

type infrastructureContainer struct {
	path string
	opts []diagram.NodeOption
}

var Infrastructure = &infrastructureContainer{
	opts: diagram.OptionSet{diagram.Provider("yc"), diagram.NodeShape("none")},
	path: "assets/yc/infrastructure",
}

func (c *infrastructureContainer) DdosProtection(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/ddos-protection.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) ManagementConsole(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/management-console.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Monitoring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/monitoring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) ResourceManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/resource-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) CertificateManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/certificate-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Compute(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/compute.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) ContainerRegistry(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/container-registry.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Interconnect(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/interconnect.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Kms(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/kms.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) ObjectStorage(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/object-storage.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Iam(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/iam.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) LoadBalancer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/load-balancer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) Vpc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/vpc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) InstanceGroup(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/instance-group.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *infrastructureContainer) ManagedKubernetes(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/infrastructure/managed-kubernetes.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

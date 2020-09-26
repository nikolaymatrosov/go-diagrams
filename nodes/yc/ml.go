package yc

import "github.com/blushft/go-diagrams/diagram"

type mlContainer struct {
	path string
	opts []diagram.NodeOption
}

var Ml = &mlContainer{
	opts: diagram.OptionSet{diagram.Provider("yc"), diagram.NodeShape("none")},
	path: "assets/yc/ml",
}

func (c *mlContainer) Translate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/ml/translate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Vision(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/ml/vision.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Api(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/ml/api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Datasphere(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/ml/datasphere.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Locator(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/ml/locator.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *mlContainer) Speechkit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/ml/speechkit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

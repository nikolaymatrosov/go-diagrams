package yc

import "github.com/blushft/go-diagrams/diagram"

type serverlessContainer struct {
	path string
	opts []diagram.NodeOption
}

var Serverless = &serverlessContainer{
	opts: diagram.OptionSet{diagram.Provider("yc"), diagram.NodeShape("none")},
	path: "assets/yc/serverless",
}

func (c *serverlessContainer) ApiGateway(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/serverless/api-gateway.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *serverlessContainer) Functions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/serverless/functions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *serverlessContainer) IotCore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/serverless/iot-core.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *serverlessContainer) MessageQueue(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/serverless/message-queue.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *serverlessContainer) ServerlessFunctions(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/serverless/serverless-functions.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *serverlessContainer) ServerlessTriggers(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/serverless/serverless-triggers.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *serverlessContainer) Ydb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/serverless/ydb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

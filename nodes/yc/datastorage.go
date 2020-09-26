package yc

import "github.com/blushft/go-diagrams/diagram"

type datastorageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Datastorage = &datastorageContainer{
	opts: diagram.OptionSet{diagram.Provider("yc"), diagram.NodeShape("none")},
	path: "assets/yc/datastorage",
}

func (c *datastorageContainer) ManagedClickhouse(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-clickhouse.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) ManagedElasticsearch(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-elasticsearch.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) ManagedKafka(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-kafka.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) ManagedMysql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-mysql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) ManagedPostgresql(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-postgresql.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) ManagedSqlserver(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-sqlserver.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) Mdb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/mdb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) Datalens(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/datalens.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) DataTransfer(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/data-transfer.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) ManagedMongodb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-mongodb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) ManagedRedis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/managed-redis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *datastorageContainer) DataProc(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/yc/datastorage/data-proc.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

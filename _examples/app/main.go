package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/yc"
)

func main() {
	d, err := diagram.New(diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	user := apps.Client.User(diagram.NodeLabel("User"))
	lb := yc.Infrastructure.LoadBalancer(diagram.NodeLabel("LB"))
	cache := yc.Datastorage.ManagedRedis(diagram.NodeLabel("Cache"))
	db := yc.Datastorage.ManagedPostgresql(diagram.NodeLabel("Database"))

	dc := diagram.NewGroup("YandexCloud").Label("YandexCloud")
	dc.NewGroup("services").
		Label("Service Layer").
		Add(
			yc.Infrastructure.Compute(diagram.NodeLabel("Server 1")),
			yc.Infrastructure.Compute(diagram.NodeLabel("Server 2")),
			yc.Infrastructure.Compute(diagram.NodeLabel("Server 3")),
		).
		ConnectAllFrom(lb.ID(), diagram.Forward()).
		ConnectAllTo(cache.ID(), diagram.Forward())

	dc.NewGroup("data").Label("Data Layer").Add(cache, db).Connect(cache, db)

	d.Connect(user, lb, diagram.Forward()).Group(dc)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}

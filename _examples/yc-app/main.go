package main

import (
	"log"
	"strconv"

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

	dc := diagram.NewGroup("YandexCloud")

	vmNum := 0
	var servers []*diagram.Node

	services := dc.NewGroup("services").
		Label("Service Layer")

	for _, az := range []string{"A", "B", "C"} {
		azNode := services.NewGroup("AZ" + az).
			Label("AZ " + az)

		for i := 0; i < 1; i++ {
			vmNum += 1
			server := yc.Infrastructure.Compute(diagram.NodeLabel("Server " + strconv.Itoa(vmNum)))
			azNode.Add(server).
				ConnectAllFrom(lb.ID(), diagram.Forward()).
				ConnectAllTo(cache.ID(), diagram.Forward()).
				ConnectAllTo(db.ID(), diagram.Forward())

			servers = append(servers, server)
		}
	}

	dc.NewGroup("data").Label("Data Layer").Add(cache, db)

	d.Connect(user, lb, diagram.Forward()).Group(dc)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}

package store

import (
	"fmt"
	"github.com/4whomtbts/timefabric/model"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)


func saveNode(db *sqlx.DB, node *model.TimeFabricNode) error {
	stat := "INSERT INTO NODE " +
		"(storage_group_id, host_name, node_name, host_ip, port, excluded, exclusive, registered_at) " +
		"VALUES (:storage_group_id, :host_name, :node_name, :host_ip, :port, :excluded, :exclusive, :registered_at)"
	rst, err := db.NamedExec(stat, node)
	if err != nil {
		log.Errorf("failed to insert NODE %s", err.Error())
		return err
	}
	fmt.Println(rst.LastInsertId())
	return nil
}
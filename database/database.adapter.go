package database

import (
	"fmt"

	"github.com/Hectorcaste/test-gocluster/database/models"
	"gorm.io/gorm"
)

type LealDBAdapter struct {
	conn *gorm.DB
}

func NewLealDBAdapter(conn *gorm.DB) *LealDBAdapter {
	return &LealDBAdapter{
		conn: conn,
	}
}

func (adapter *LealDBAdapter) GetSucursales() *[]models.BranchModel {
	model := &[]models.BranchModel{}

	queryDB := adapter.conn.Raw(`select cs.nombre ,cs.latitud ,cs.longitud  from com_sucursales  cs
	left join com_comercios cc on cs.id_comercio = cc.id_comercio 
	where cc.cod_pais = 'CO' and cs.id_ciudad = 1 limit 300
	`).Scan(&model)

	fmt.Printf("hola %s", queryDB)

	return model
}

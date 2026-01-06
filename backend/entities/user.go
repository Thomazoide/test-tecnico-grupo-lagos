package entities

import "gorm.io/gorm"

// Como no se pidio un manejo de sesiones ni autenticacion y autorizacion solo usare un id y un UUID para guardar en BBDD para no hacer un formulario de registro ni de inicio de sesion
type User struct {
	gorm.Model
	ExternalUUID string `gorm:"size:64;index"`
	Cart         Cart   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

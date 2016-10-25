package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}

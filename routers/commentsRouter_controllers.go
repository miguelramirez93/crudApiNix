package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DesagregacionIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DestinoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:DisponibilidadRubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EntidadHomologacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:FuenteFinanciacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RegistroPresupuestalDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/crudApiNix/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}

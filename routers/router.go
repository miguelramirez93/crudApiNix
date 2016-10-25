// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/miguelramirez93/crudApiNix/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/anulacion_reserva",
			beego.NSInclude(
				&controllers.AnulacionReservaController{},
			),
		),

		beego.NSNamespace("/estado_apropiacion",
			beego.NSInclude(
				&controllers.EstadoApropiacionController{},
			),
		),

		beego.NSNamespace("/desagregacion_ingreso",
			beego.NSInclude(
				&controllers.DesagregacionIngresoController{},
			),
		),

		beego.NSNamespace("/destino_disponibilidad",
			beego.NSInclude(
				&controllers.DestinoDisponibilidadController{},
			),
		),

		beego.NSNamespace("/estado_disponibilidad",
			beego.NSInclude(
				&controllers.EstadoDisponibilidadController{},
			),
		),

		beego.NSNamespace("/fuente_financiacion",
			beego.NSInclude(
				&controllers.FuenteFinanciacionController{},
			),
		),

		beego.NSNamespace("/ingreso",
			beego.NSInclude(
				&controllers.IngresoController{},
			),
		),

		beego.NSNamespace("/apropiacion",
			beego.NSInclude(
				&controllers.ApropiacionController{},
			),
		),

		beego.NSNamespace("/modificacion_presupuestal",
			beego.NSInclude(
				&controllers.ModificacionPresupuestalController{},
			),
		),

		beego.NSNamespace("/estado_registro_presupuestal",
			beego.NSInclude(
				&controllers.EstadoRegistroPresupuestalController{},
			),
		),

		beego.NSNamespace("/disponibilidad",
			beego.NSInclude(
				&controllers.DisponibilidadController{},
			),
		),

		beego.NSNamespace("/disponibilidad_rubro",
			beego.NSInclude(
				&controllers.DisponibilidadRubroController{},
			),
		),

		beego.NSNamespace("/registro_presupuestal_disponibilidad",
			beego.NSInclude(
				&controllers.RegistroPresupuestalDisponibilidadController{},
			),
		),

		beego.NSNamespace("/estado_reserva_presupuestal",
			beego.NSInclude(
				&controllers.EstadoReservaPresupuestalController{},
			),
		),

		beego.NSNamespace("/reserva_presupuestal",
			beego.NSInclude(
				&controllers.ReservaPresupuestalController{},
			),
		),

		beego.NSNamespace("/entidad",
			beego.NSInclude(
				&controllers.EntidadController{},
			),
		),

		beego.NSNamespace("/entidad_homologacion",
			beego.NSInclude(
				&controllers.EntidadHomologacionController{},
			),
		),

		beego.NSNamespace("/rubro_homologado",
			beego.NSInclude(
				&controllers.RubroHomologadoController{},
			),
		),

		beego.NSNamespace("/rubro",
			beego.NSInclude(
				&controllers.RubroController{},
			),
		),

		beego.NSNamespace("/rubro_rubro",
			beego.NSInclude(
				&controllers.RubroRubroController{},
			),
		),

		beego.NSNamespace("/concepto",
			beego.NSInclude(
				&controllers.ConceptoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

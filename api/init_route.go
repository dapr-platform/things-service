package api

import "github.com/go-chi/chi/v5"

func InitRoute(r chi.Router) {
	InitProductRoute(r)
	InitDevice_modelRoute(r)
	InitDeviceRoute(r)
	InitCustomDeviceDataRoute(r)
	InitTestRoute(r)
	InitCustomTagRoute(r)
	InitAlarm_ruleRoute(r)
	InitModel_metaRoute(r)
	InitCustomDeviceRoute(r)
	InitCustomProductRoute(r)
	InitDevice_mirrorRoute(r)
	InitDevice_infoRoute(r)
	InitSim_deviceRoute(r)
	InitTagRoute(r)
	InitCustomSim_deviceRoute(r)
	InitCustomDevice_modelRoute(r)
	InitUser_deviceRoute(r)
	InitUser_device_infoRoute(r)
	InitKpi_infoRoute(r)
	initDbExcelRouter(r)
	InitDeviceExcelImportRoute(r)
	InitCustomDeviceMirrorRoute(r)
	InitMetricsQueryHandler(r)
	InitMonitorHandler(r)
}

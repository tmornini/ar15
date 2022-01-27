package sensors

import (
	"github.com/tmornini/ar15"
)

func Start(ar ar15.Model) {
	ar.Info("Starting Sensors")

	StartBoltReleaseSensor(ar)
	StartChargingHandleSensor(ar)
	StartEjectionPortCoverSensor(ar)
	StartForwardAssistSensor(ar)
	StartMagazineReleaseSensor(ar)
	StartMagazineWellSensor(ar)
	StartSafetySensor(ar)
	StartTriggerSensor(ar)

	ar.Info("Sensors started")
}

func StartBoltReleaseSensor(ar ar15.Model) {
	ar.Info("Starting BoltReleaseSensor")

	b := boltReleaseSensor{ar}

	go b.cycle()
}

type boltReleaseSensor struct {
	ar15.Model
}

func (b boltReleaseSensor) cycle() {
	// bInfo("Cycling BoltReleaseSensor")

	select {}
}

func StartChargingHandleSensor(ar ar15.Model) {
	// arInfo("Starting ChargingHandleSensor")

	b := chargingHandleSensor{ar}

	go b.cycle()
}

type chargingHandleSensor struct {
	ar15.Model
}

func (b chargingHandleSensor) cycle() {
	// bInfo("Cycling ChargingHandleSensor")

	select {}
}

func StartEjectionPortCoverSensor(ar ar15.Model) {
	// arInfo("Starting EjectionPortCoverSensor")

	b := ejectionPortCoverSensor{ar}

	go b.cycle()
}

type ejectionPortCoverSensor struct {
	ar15.Model
}

func (b ejectionPortCoverSensor) cycle() {
	// bInfo("Cycling ChargingHandleSensor")

	select {}
}

func StartForwardAssistSensor(ar ar15.Model) {
	ar.Info("Starting ForwardAssistSensor")

	b := forwardAssistSensor{ar}

	go b.cycle()
}

type forwardAssistSensor struct {
	ar15.Model
}

func (b forwardAssistSensor) cycle() {
	// bInfo("Cycling ChargingHandleSensor")

	select {}
}

func StartMagazineReleaseSensor(ar ar15.Model) {
	ar.Info("Starting MagazineReleaseSensor")

	b := magazineReleaseSensor{ar}

	go b.cycle()
}

type magazineReleaseSensor struct {
	ar15.Model
}

func (b magazineReleaseSensor) cycle() {
	// bInfo("Cycling MagazineReleaseSensor")

	select {}
}

func StartMagazineWellSensor(ar ar15.Model) {
	ar.Info("Starting MagazineWellSensor")

	b := magazineWellSensor{ar}

	go b.cycle()
}

type magazineWellSensor struct {
	ar15.Model
}

func (b magazineWellSensor) cycle() {
	// bInfo("Cycling MagazineWellSensor")

	select {}
}

func StartSafetySensor(ar ar15.Model) {
	ar.Info("Starting SafteySensor")

	b := safetySensor{ar}

	go b.cycle()
}

type safetySensor struct {
	ar15.Model
}

func (b safetySensor) cycle() {
	// bInfo("Cycling SafteySensor")

	select {}
}

func StartTriggerSensor(ar ar15.Model) {
	ar.Info("Starting TriggerSensor")

	b := triggerSensor{ar}

	go b.cycle()
}

type triggerSensor struct {
	ar15.Model
}

func (b triggerSensor) cycle() {
	// bInfo("Cycling TriggerSensor")

	select {}
}

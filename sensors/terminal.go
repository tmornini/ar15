// Package sensors models sensors
package sensors

import "github.com/tmornini/ar15/weapons"

// Start starts the sensors
func Start(ar weapons.AR15) {
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

// StartBoltReleaseSensor starts the bolt release sensor
func StartBoltReleaseSensor(ar weapons.AR15) {
	ar.Info("Starting BoltReleaseSensor")

	b := boltReleaseSensor{ar}

	go b.cycle()
}

type boltReleaseSensor struct {
	weapons.AR15
}

func (b boltReleaseSensor) cycle() {
	// bInfo("Cycling BoltReleaseSensor")

	select {}
}

// StartChargingHandleSensor starts the charging handle sensor
func StartChargingHandleSensor(ar weapons.AR15) {
	// arInfo("Starting ChargingHandleSensor")

	b := chargingHandleSensor{ar}

	go b.cycle()
}

type chargingHandleSensor struct {
	weapons.AR15
}

func (b chargingHandleSensor) cycle() {
	// bInfo("Cycling ChargingHandleSensor")

	select {}
}

// StartEjectionPortCoverSensor starts the ejection port cover sensor
func StartEjectionPortCoverSensor(ar weapons.AR15) {
	// arInfo("Starting EjectionPortCoverSensor")

	b := ejectionPortCoverSensor{ar}

	go b.cycle()
}

type ejectionPortCoverSensor struct {
	weapons.AR15
}

func (b ejectionPortCoverSensor) cycle() {
	// bInfo("Cycling ChargingHandleSensor")

	select {}
}

// StartForwardAssistSensor starts the forward assist sensor
func StartForwardAssistSensor(ar weapons.AR15) {
	ar.Info("Starting ForwardAssistSensor")

	b := forwardAssistSensor{ar}

	go b.cycle()
}

type forwardAssistSensor struct {
	weapons.AR15
}

func (b forwardAssistSensor) cycle() {
	// bInfo("Cycling ChargingHandleSensor")

	select {}
}

// StartMagazineReleaseSensor starts the magazine release sensor
func StartMagazineReleaseSensor(ar weapons.AR15) {
	ar.Info("Starting MagazineReleaseSensor")

	b := magazineReleaseSensor{ar}

	go b.cycle()
}

type magazineReleaseSensor struct {
	weapons.AR15
}

func (b magazineReleaseSensor) cycle() {
	// bInfo("Cycling MagazineReleaseSensor")

	select {}
}

// StartMagazineWellSensor starts the magazine well sensor
func StartMagazineWellSensor(ar weapons.AR15) {
	ar.Info("Starting MagazineWellSensor")

	b := magazineWellSensor{ar}

	go b.cycle()
}

type magazineWellSensor struct {
	weapons.AR15
}

func (b magazineWellSensor) cycle() {
	// bInfo("Cycling MagazineWellSensor")

	select {}
}

// StartSafetySensor starts the safety sensor
func StartSafetySensor(ar weapons.AR15) {
	ar.Info("Starting SafteySensor")

	b := safetySensor{ar}

	go b.cycle()
}

type safetySensor struct {
	weapons.AR15
}

func (b safetySensor) cycle() {
	// bInfo("Cycling SafteySensor")

	select {}
}

// StartTriggerSensor starts the trigger sensor
func StartTriggerSensor(ar weapons.AR15) {
	ar.Info("Starting TriggerSensor")

	b := triggerSensor{ar}

	go b.cycle()
}

type triggerSensor struct {
	weapons.AR15
}

func (b triggerSensor) cycle() {
	// bInfo("Cycling TriggerSensor")

	select {}
}

package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"silburyslot/rng"
	"silburyslot/routes"
)

const (
	CHERRY = iota
	ORANGE
	PLUM
	BELL
	SEVEN
	CHOCOLATE
	BAR
)

func main() {
	loadEnvFile(".env")
	symbols := []rng.Symbol{
		{SymbolId: CHERRY, WeightsPerReel: []int{1, 0, 0}, BetMultiplier: 2},
		{SymbolId: ORANGE, WeightsPerReel: []int{3, 0, 0}, BetMultiplier: 4},
		{SymbolId: PLUM, WeightsPerReel: []int{0, 0, 0}, BetMultiplier: 6},
		{SymbolId: BELL, WeightsPerReel: []int{0, 0, 0}, BetMultiplier: 8},
		{SymbolId: SEVEN, WeightsPerReel: []int{0, 0, 0}, BetMultiplier: 10},
		{SymbolId: CHOCOLATE, WeightsPerReel: []int{0, 1, 0}, BetMultiplier: 12},
		{SymbolId: BAR, WeightsPerReel: []int{0, 0, 1}, BetMultiplier: 20},
	}

	slotMachine := rng.SlotMachine{}
	for i := range len(symbols[0].WeightsPerReel) {
		newReel := rng.BuildReel(symbols, i)
		slotMachine = append(slotMachine, newReel)
	}

	http.HandleFunc("/", ShowIndex(slotMachine))
	http.HandleFunc("/sampleEndpoint", routes.Endpoint)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func SpinTheWheels(sl rng.SlotMachine) (slots []int) {
	for i := range sl {
		slots = append(slots, sl[i].PickRandomSymbol().SymbolId)
	}
	return slots
}

func ShowIndex(sl rng.SlotMachine) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		templateFile, _ := os.Open("index.gohtml")
		tplString, _ := io.ReadAll(templateFile)
		defer templateFile.Close()

		tpl, _ := template.New("index").Parse(string(tplString))
		tpl.Execute(w, SpinTheWheels(sl))
	}
}

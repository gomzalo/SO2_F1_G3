package models

type MemsimProc struct {
	Ciclo    int      `json:"ciclo"`
	Procesos []string `json:"procesos"`
}
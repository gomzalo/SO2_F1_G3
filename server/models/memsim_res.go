package models

type MemsimRes struct {
	Memsim   []MemsimProc `json:"memsim"`
	Duracion int64        `json:"duracion"`
}
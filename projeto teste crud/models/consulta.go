package models

import "time"

type Consulta struct {
	ID            int    `json:"id"`
	NomePaciente  string `json:"nome_paciente"`
	CPFPaciente   string `json:"cpf_paciente"`
	Telefone      string `json:"telefone,omitempty"`
	Especialidade string `json:"especialidade"`
	NomeMedico    string `json:"nome_medico"`
	DataConsulta  string `json:"data_consulta"` // formato esperado: YYYY-MM-DD
	Horario       string `json:"horario"`       // formato esperado: HH:MM:SS
	Status        string `json:"status,omitempty"`
	Observacoes   string `json:"observacoes,omitempty"`
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func ParseTime(value string) (time.Time, error) {
	return time.Parse("15:04:05", value)
}

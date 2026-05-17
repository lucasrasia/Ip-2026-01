package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"projeto-teste-crud/models"
)

// ConsultaHandler concentra as dependências dos endpoints de consultas.
type ConsultaHandler struct {
	DB *sql.DB
}

func NewConsultaHandler(db *sql.DB) *ConsultaHandler {
	return &ConsultaHandler{DB: db}
}

func (h *ConsultaHandler) ListarConsultas(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.DB.Query(`SELECT id, nome_paciente, cpf_paciente, telefone, especialidade, nome_medico, data_consulta, horario, status, observacoes FROM consultas ORDER BY id`)
	if err != nil {
		Error(w, http.StatusInternalServerError, "erro ao listar consultas")
		return
	}
	defer rows.Close()

	consultas := []models.Consulta{}
	for rows.Next() {
		var c models.Consulta
		var data time.Time
		var horario time.Time
		if err := rows.Scan(&c.ID, &c.NomePaciente, &c.CPFPaciente, &c.Telefone, &c.Especialidade, &c.NomeMedico, &data, &horario, &c.Status, &c.Observacoes); err != nil {
			Error(w, http.StatusInternalServerError, "erro ao ler resultado")
			return
		}
		c.DataConsulta = data.Format("2006-01-02")
		c.Horario = horario.Format("15:04:05")
		consultas = append(consultas, c)
	}

	JSON(w, http.StatusOK, consultas)
}

func (h *ConsultaHandler) BuscarConsulta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDFromPath(r.URL.Path)
	if !ok {
		Error(w, http.StatusBadRequest, "id inválido")
		return
	}

	var c models.Consulta
	var data time.Time
	var horario time.Time
	err := h.DB.QueryRow(`SELECT id, nome_paciente, cpf_paciente, telefone, especialidade, nome_medico, data_consulta, horario, status, observacoes FROM consultas WHERE id = $1`, id).
		Scan(&c.ID, &c.NomePaciente, &c.CPFPaciente, &c.Telefone, &c.Especialidade, &c.NomeMedico, &data, &horario, &c.Status, &c.Observacoes)
	if err == sql.ErrNoRows {
		Error(w, http.StatusNotFound, "consulta não encontrada")
		return
	}
	if err != nil {
		Error(w, http.StatusInternalServerError, "erro ao buscar consulta")
		return
	}
	c.DataConsulta = data.Format("2006-01-02")
	c.Horario = horario.Format("15:04:05")

	JSON(w, http.StatusOK, c)
}

func (h *ConsultaHandler) CriarConsulta(w http.ResponseWriter, r *http.Request) {
	var c models.Consulta
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		Error(w, http.StatusBadRequest, "json inválido")
		return
	}

	if msg := validarConsulta(c); msg != "" {
		Error(w, http.StatusBadRequest, msg)
		return
	}

	var id int
	err := h.DB.QueryRow(`INSERT INTO consultas (nome_paciente, cpf_paciente, telefone, especialidade, nome_medico, data_consulta, horario, status, observacoes)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`,
		c.NomePaciente, c.CPFPaciente, c.Telefone, c.Especialidade, c.NomeMedico, c.DataConsulta, c.Horario, c.Status, c.Observacoes).Scan(&id)
	if err != nil {
		Error(w, http.StatusInternalServerError, "erro ao criar consulta")
		return
	}

	c.ID = id
	JSON(w, http.StatusCreated, c)
}

func (h *ConsultaHandler) AtualizarConsulta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDFromPath(r.URL.Path)
	if !ok {
		Error(w, http.StatusBadRequest, "id inválido")
		return
	}

	var c models.Consulta
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		Error(w, http.StatusBadRequest, "json inválido")
		return
	}

	if msg := validarConsulta(c); msg != "" {
		Error(w, http.StatusBadRequest, msg)
		return
	}

	res, err := h.DB.Exec(`UPDATE consultas SET nome_paciente=$1, cpf_paciente=$2, telefone=$3, especialidade=$4, nome_medico=$5, data_consulta=$6, horario=$7, status=$8, observacoes=$9 WHERE id=$10`,
		c.NomePaciente, c.CPFPaciente, c.Telefone, c.Especialidade, c.NomeMedico, c.DataConsulta, c.Horario, c.Status, c.Observacoes, id)
	if err != nil {
		Error(w, http.StatusInternalServerError, "erro ao atualizar consulta")
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		Error(w, http.StatusNotFound, "consulta não encontrada")
		return
	}

	c.ID = id
	JSON(w, http.StatusOK, c)
}

func (h *ConsultaHandler) DeletarConsulta(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDFromPath(r.URL.Path)
	if !ok {
		Error(w, http.StatusBadRequest, "id inválido")
		return
	}

	res, err := h.DB.Exec(`DELETE FROM consultas WHERE id = $1`, id)
	if err != nil {
		Error(w, http.StatusInternalServerError, "erro ao deletar consulta")
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		Error(w, http.StatusNotFound, "consulta não encontrada")
		return
	}

	JSON(w, http.StatusOK, map[string]string{"mensagem": "consulta removida com sucesso"})
}

func parseIDFromPath(path string) (int, bool) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		return 0, false
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, false
	}
	return id, true
}

func validarConsulta(c models.Consulta) string {
	if c.NomePaciente == "" || c.CPFPaciente == "" || c.Especialidade == "" || c.NomeMedico == "" || c.DataConsulta == "" || c.Horario == "" {
		return "campos obrigatórios não preenchidos"
	}
	if _, err := models.ParseDate(c.DataConsulta); err != nil {
		return "data_consulta deve estar no formato YYYY-MM-DD"
	}
	if _, err := models.ParseTime(c.Horario); err != nil {
		return "horario deve estar no formato HH:MM:SS"
	}
	return ""
}

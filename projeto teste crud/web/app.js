const API_URL = "/consultas";

const form = document.querySelector("#consultaForm");
const consultasList = document.querySelector("#consultasList");
const emptyState = document.querySelector("#emptyState");
const message = document.querySelector("#message");
const counter = document.querySelector("#counter");
const refreshButton = document.querySelector("#refreshButton");
const cancelEditButton = document.querySelector("#cancelEditButton");
const submitButton = document.querySelector("#submitButton");
const formTitle = document.querySelector("#formTitle");

const fields = {
  id: document.querySelector("#consultaId"),
  nomePaciente: document.querySelector("#nomePaciente"),
  cpfPaciente: document.querySelector("#cpfPaciente"),
  telefone: document.querySelector("#telefone"),
  especialidade: document.querySelector("#especialidade"),
  nomeMedico: document.querySelector("#nomeMedico"),
  dataConsulta: document.querySelector("#dataConsulta"),
  horario: document.querySelector("#horario"),
  status: document.querySelector("#status"),
  observacoes: document.querySelector("#observacoes"),
};

let consultas = [];

function setLoading(isLoading) {
  submitButton.disabled = isLoading;
  refreshButton.disabled = isLoading;
  submitButton.textContent = isLoading ? "Salvando..." : fields.id.value ? "Atualizar consulta" : "Salvar consulta";
}

function showMessage(text, type = "success") {
  message.textContent = text;
  message.className = `message show ${type}`;
  window.clearTimeout(showMessage.timeout);
  showMessage.timeout = window.setTimeout(() => {
    message.className = "message";
    message.textContent = "";
  }, 4200);
}

function consultaFromForm() {
  return {
    nome_paciente: fields.nomePaciente.value.trim(),
    cpf_paciente: fields.cpfPaciente.value.trim(),
    telefone: fields.telefone.value.trim(),
    especialidade: fields.especialidade.value.trim(),
    nome_medico: fields.nomeMedico.value.trim(),
    data_consulta: fields.dataConsulta.value,
    horario: normalizeHorario(fields.horario.value),
    status: fields.status.value,
    observacoes: fields.observacoes.value.trim(),
  };
}

function normalizeHorario(value) {
  if (!value) {
    return "";
  }
  return value.length === 5 ? `${value}:00` : value;
}

function fillForm(consulta) {
  fields.id.value = consulta.id;
  fields.nomePaciente.value = consulta.nome_paciente;
  fields.cpfPaciente.value = consulta.cpf_paciente;
  fields.telefone.value = consulta.telefone || "";
  fields.especialidade.value = consulta.especialidade;
  fields.nomeMedico.value = consulta.nome_medico;
  fields.dataConsulta.value = consulta.data_consulta;
  fields.horario.value = consulta.horario?.slice(0, 8) || "";
  fields.status.value = consulta.status || "agendada";
  fields.observacoes.value = consulta.observacoes || "";

  formTitle.textContent = `Editando consulta #${consulta.id}`;
  submitButton.textContent = "Atualizar consulta";
  cancelEditButton.hidden = false;
  fields.nomePaciente.focus();
}

function resetForm() {
  form.reset();
  fields.id.value = "";
  fields.status.value = "agendada";
  formTitle.textContent = "Nova consulta";
  submitButton.textContent = "Salvar consulta";
  cancelEditButton.hidden = true;
}

async function requestJSON(url, options = {}) {
  const response = await fetch(url, {
    headers: { "Content-Type": "application/json", ...options.headers },
    ...options,
  });

  const payload = await response.json().catch(() => null);
  if (!response.ok) {
    throw new Error(payload?.erro || "Não foi possível concluir a operação.");
  }
  return payload;
}

async function loadConsultas() {
  try {
    consultas = await requestJSON(API_URL);
    renderConsultas();
  } catch (error) {
    consultas = [];
    renderConsultas();
    showMessage(error.message, "error");
  }
}

async function saveConsulta(event) {
  event.preventDefault();
  setLoading(true);

  const id = fields.id.value;
  const method = id ? "PUT" : "POST";
  const url = id ? `${API_URL}/${id}` : API_URL;

  try {
    await requestJSON(url, {
      method,
      body: JSON.stringify(consultaFromForm()),
    });
    showMessage(id ? "Consulta atualizada com sucesso." : "Consulta criada com sucesso.");
    resetForm();
    await loadConsultas();
  } catch (error) {
    showMessage(error.message, "error");
  } finally {
    setLoading(false);
  }
}

async function deleteConsulta(id) {
  const consulta = consultas.find((item) => item.id === id);
  const nome = consulta?.nome_paciente || `#${id}`;

  if (!window.confirm(`Remover a consulta de ${nome}?`)) {
    return;
  }

  try {
    await requestJSON(`${API_URL}/${id}`, { method: "DELETE" });
    showMessage("Consulta removida com sucesso.");
    if (fields.id.value === String(id)) {
      resetForm();
    }
    await loadConsultas();
  } catch (error) {
    showMessage(error.message, "error");
  }
}

function renderConsultas() {
  consultasList.innerHTML = "";
  counter.textContent = `${consultas.length} ${consultas.length === 1 ? "consulta" : "consultas"}`;
  emptyState.hidden = consultas.length > 0;

  consultas.forEach((consulta) => {
    const card = document.createElement("article");
    card.className = "card";
    card.innerHTML = `
      <div class="card-header">
        <div>
          <h3>${escapeHTML(consulta.nome_paciente)}</h3>
          <p>${escapeHTML(consulta.especialidade)} · ${escapeHTML(consulta.nome_medico)}</p>
        </div>
        <span class="badge">${escapeHTML(consulta.status || "sem status")}</span>
      </div>
      <div class="card-grid">
        <p><span>CPF</span>${escapeHTML(consulta.cpf_paciente)}</p>
        <p><span>Telefone</span>${escapeHTML(consulta.telefone || "Não informado")}</p>
        <p><span>Data</span>${formatDate(consulta.data_consulta)}</p>
        <p><span>Horário</span>${escapeHTML(consulta.horario)}</p>
      </div>
      ${consulta.observacoes ? `<p class="observacoes">${escapeHTML(consulta.observacoes)}</p>` : ""}
      <div class="card-actions">
        <button class="card-button" type="button" data-action="edit" data-id="${consulta.id}">Editar</button>
        <button class="card-button danger" type="button" data-action="delete" data-id="${consulta.id}">Remover</button>
      </div>
    `;
    consultasList.appendChild(card);
  });
}

function formatDate(value) {
  if (!value) {
    return "Não informada";
  }

  const [year, month, day] = value.split("-");
  return `${day}/${month}/${year}`;
}

function escapeHTML(value) {
  return String(value ?? "")
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#039;");
}

form.addEventListener("submit", saveConsulta);
refreshButton.addEventListener("click", loadConsultas);
cancelEditButton.addEventListener("click", resetForm);

consultasList.addEventListener("click", (event) => {
  const button = event.target.closest("button[data-action]");
  if (!button) {
    return;
  }

  const id = Number(button.dataset.id);
  const consulta = consultas.find((item) => item.id === id);

  if (button.dataset.action === "edit" && consulta) {
    fillForm(consulta);
  }

  if (button.dataset.action === "delete") {
    deleteConsulta(id);
  }
});

loadConsultas();

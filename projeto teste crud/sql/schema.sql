-- Criação da tabela principal do sistema de agendamento de consultas
CREATE TABLE IF NOT EXISTS consultas (
    id SERIAL PRIMARY KEY,
    nome_paciente TEXT NOT NULL,
    cpf_paciente TEXT NOT NULL,
    telefone TEXT,
    especialidade TEXT NOT NULL,
    nome_medico TEXT NOT NULL,
    data_consulta DATE NOT NULL,
    horario TIME NOT NULL,
    status TEXT,
    observacoes TEXT
);

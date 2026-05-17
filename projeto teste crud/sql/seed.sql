-- Dados iniciais para testes da API
INSERT INTO consultas (
    nome_paciente,
    cpf_paciente,
    telefone,
    especialidade,
    nome_medico,
    data_consulta,
    horario,
    status,
    observacoes
) VALUES
('Maria Silva', '123.456.789-00', '(11) 99999-0001', 'Cardiologia', 'Dr. João Pereira', '2026-05-20', '09:00:00', 'agendada', 'Paciente com histórico familiar de hipertensão'),
('Carlos Souza', '987.654.321-00', '(11) 98888-0002', 'Dermatologia', 'Dra. Ana Costa', '2026-05-21', '14:30:00', 'confirmada', 'Retorno para avaliação de tratamento');

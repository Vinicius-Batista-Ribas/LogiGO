# 💰 PayTrack – Gerenciador de Pagamentos (Backend em Go)

Este projeto é um **sistema backend para gerenciamento de pagamentos e recebimentos**, desenvolvido em **Golang**, com foco em aprendizado progressivo, boas práticas de backend e arquitetura de software.

O sistema permite registrar entradas e saídas financeiras, classificando pagamentos por data, valor, descrição e método de pagamento (PIX, crédito, débito, etc.).

---

## 🎯 Objetivo do Projeto

O objetivo deste projeto é evoluir gradualmente um **gerenciador financeiro**, começando com uma API REST simples e avançando para conceitos utilizados em sistemas reais de mercado, como cache, mensageria, microsserviços, CI/CD, deploy em nuvem e monitoramento.

Cada etapa do projeto fecha um ciclo funcional antes de avançar para a próxima.

---

## 🧠 Visão Geral do Sistema

O sistema será responsável por:

* Registrar pagamentos e recebimentos
* Controlar entradas e saídas financeiras
* Calcular saldo
* Manter histórico financeiro
* (Futuro) Processar eventos de forma assíncrona
* (Futuro) Operar em arquitetura distribuída

---

## 🧩 Tecnologias

* **Linguagem:** Go
* **API:** HTTP REST
* **Banco de Dados:** PostgreSQL
* **Cache:** Redis (planejado)
* **Mensageria:** RabbitMQ / Kafka (planejado)
* **Containerização:** Docker (planejado)
* **CI/CD:** GitHub Actions (planejado)
* **Cloud:** AWS / GCP / DigitalOcean (planejado)
* **Monitoramento:** Prometheus, Grafana (planejado)

---

## 🗂️ Estrutura Inicial do Projeto

```bash
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── model/
├── go.mod
├── go.sum
└── README.md
```

---

## 📦 Modelo Inicial: Payment

O sistema começa com a entidade **Payment**, que representa uma movimentação financeira.

Campos principais:

* Data
* Valor
* Descrição
* Método de pagamento (PIX, crédito, débito, etc.)
* Tipo (entrada ou saída)

Esse modelo será expandido conforme o projeto evolui.

---

## 🛣️ Roadmap de Evolução

* [x] Fase 0 – API básica em Go
* [x] Fase 1 – Persistência com PostgreSQL
* [ ] Fase 2 – Organização em camadas
* [ ] Fase 3 – Cache com Redis
* [ ] Fase 4 – Mensageria
* [ ] Fase 5 – Microsserviços
* [ ] Fase 6 – Docker
* [ ] Fase 7 – CI/CD
* [ ] Fase 8 – Deploy em nuvem
* [ ] Fase 9 – Monitoramento

---

## ▶️ Como executar o projeto

```bash
go run cmd/api/main.go
```

---

## 📌 Observações

Este projeto serve como laboratório prático para aprendizado em backend, arquitetura de software e boas práticas de desenvolvimento.

---

## 👨‍💻 Autor

Desenvolvido por **Vinícius Batista Ribas**

* Backend Developer (Golang)
* Foco em APIs, microsserviços e arquitetura de software

---

> 🚀 Este projeto não é apenas um CRUD, é um laboratório de backend profissional.

# ERPAA (Enterprise Resource Planning & Analyst)

![ERPAA Logo](./assets/logo.png)

---

## 📖 Overview
**ERPAA** is an enterprise‑grade **Resource Planning and Analyst** platform that seamlessly integrates **AI‑powered business analytics** into traditional ERP workflows. Designed for midsize to large organizations, ERPAA provides:

- **Comprehensive resource planning** (inventory, procurement, HR, finance).
- **AI‑driven insights** for forecasting, demand‑supply optimization, and strategic decision making.
- **Modular architecture** allowing you to enable/disable components as needed.
- **Extensible plugin system** for custom business logic and third‑party integrations.

> "From raw data to actionable intelligence – all in one unified system."

---

## ✨ Key Features
| Category | Feature | Description |
|---|---|---|
| **Core ERP** | **Inventory Management** | Real‑time stock tracking, batch/lot handling, multi‑warehouse support. |
| | **Procurement** | Automated purchase orders, vendor management, contract compliance. |
| | **Human Resources** | Employee lifecycle, payroll integration, attendance & leave. |
| | **Finance** | General ledger, accounts payable/receivable, financial reporting. |
| **AI Analyst** | **Predictive Forecasting** | Time‑series models (ARIMA, Prophet) and LLM‑based scenario analysis. |
| | **Anomaly Detection** | Auto‑detect out‑liers in spend, inventory swing, production KPIs. |
| | **Natural‑Language Queries** | Ask questions like *"What is the projected cash flow for Q3?"* and receive instant visual reports. |
| **Integrations** | **REST & GraphQL APIs** | Full‑featured API layer for external systems. |
| | **Connector Marketplace** | Pre‑built connectors (Shopify, SAP, Oracle, QuickBooks). |
| **Security & Compliance** | **RBAC** | Granular role‑based access control.
| | **Audit Trail** | Immutable logs for regulatory compliance (GDPR, SOX). |

---

## 🏗️ Architecture Overview
```mermaid
flowchart LR
    subgraph Frontend[Web UI (React + Vite)]
        UI[User Interface]
    end
    subgraph Backend[Backend (Node.js/Express)]
        API[REST & GraphQL API]
        AI[AI Analyst Service]
        Core[Core ERP Services]
    end
    subgraph DB[Database Layer]
        Postgres[(PostgreSQL)]
        Redis[[Redis Cache]]
    end
    subgraph AI[AI Engine]
        LLM{{LLM Model}}
        ML[ML Models]
    end
    UI --> API
    API --> Core
    API --> AI
    Core --> Postgres
    AI --> LLM
    AI --> ML
    ML --> Postgres
    AI --> Redis
    Core --> Redis
```

---

## 🚀 Getting Started
### Prerequisites
- **Node.js** >= 18.x
- **npm** >= 9.x
- **PostgreSQL** 14+ (recommended) with a database created for ERPAA
- **Python** 3.10+ (required for AI models)
- **Docker** (optional, for containerized deployment)

### Installation (Development)
```bash
# Clone the repository
git clone https://github.com/your-org/erpaa.git
cd erpaa

# Install backend dependencies
cd backend && npm install && cd ..

# Install frontend dependencies
cd frontend && npm install && cd ..

# Set up Python environment for AI services
python -m venv .venv
.\.venv\Scripts\activate
pip install -r ai/requirements.txt
```

### Environment Configuration
Create a `.env` file in the `backend/` folder:
```dotenv
# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=erpaa
DB_USER=erp_user
DB_PASSWORD=secure_password

# Redis
REDIS_URL=redis://localhost:6379

# AI Service
LLM_API_KEY=your-openai-api-key
ML_MODEL_PATH=./ai/models/forecasting.pt
```

### Database Migration
```bash
# From the backend directory
npx prisma migrate dev --name init
```

### Running the Application
```bash
# Start backend
cd backend && npm run dev
# In a separate terminal, start frontend
cd frontend && npm run dev
# Start AI service (Python)
cd ai && python serve.py
```

Open your browser at `http://localhost:3000` (frontend) and `http://localhost:4000/api` (backend API).

---

## 📦 Production Deployment
We provide **Docker Compose** for a one‑click production setup.
```bash
docker compose -f docker-compose.prod.yml up -d
```
The compose file includes:
- `erpaa-backend`
- `erpaa-frontend`
- `erpaa-db`
- `erpaa-redis`
- `erpaa-ai`

### Scaling
- Use **Kubernetes** manifests located in `k8s/` for cloud‑native scaling.
- Enable **Horizontal Pod Autoscaler** for AI service based on GPU utilization.

---

## 🤖 AI Analyst Usage
### Natural‑Language Query Example
```text
User: "Show me a sales forecast for the next 6 months, grouped by region."
```
The AI service:
1. Parses intent.
2. Retrieves historical sales data.
3. Runs the forecasting model.
4. Returns a JSON payload and a rendered chart.

### API Endpoint
`POST /api/ai/query`
```json
{
  "question": "<your natural language query>"
}
```
Response:
```json
{
  "answer": "<textual explanation>",
  "chart": "<base64‑encoded PNG>",
  "data": { /* raw data used for the chart */ }
}
```

---

## 🛠️ Extending ERPAA
### Plugin System
Create a new folder under `backend/plugins/` with an `index.js` exporting a function:
```javascript
module.exports = function(register) {
  // Register a custom route
  register('GET', '/custom-report', async (req, res) => {
    // Your custom logic here
  });
};
```
Plugins are auto‑loaded on startup.

### Adding New AI Models
1. Place model files under `ai/models/`.
2. Update `ai/config.yaml` with the model name and path.
3. Restart the AI service.

---

## 📚 Documentation
- **API Reference** – `docs/api.md`
- **Developer Guide** – `docs/developer_guide.md`
- **AI Model Training** – `docs/ai_training.md`

---

## 🤝 Contributing
We welcome contributions! Please follow these steps:
1. Fork the repository.
2. Create a feature branch (`git checkout -b feat/awesome-feature`).
3. Follow the **code style** defined in `.eslintrc.js` and `black` for Python.
4. Write unit tests for new functionality.
5. Submit a Pull Request with a clear description.

All contributions are covered under the **Contributor Covenant Code of Conduct**.

---

## 📄 License
ERPAA is released under the **Apache License 2.0**. See `LICENSE` for details.

---

## 📬 Support & Community
- **Slack:** https://erpaa-community.slack.com
- **GitHub Issues:** https://github.com/your-org/erpaa/issues
- **Email:** support@erpaa.io

---

*Documentation generated by Antigravity – your AI‑assisted development partner.*

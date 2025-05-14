# ArbitraX

## 🧠 About ArbitraX

**ArbitraX** is a high-performance algorithmic trading platform designed to execute real-time currency arbitrage and strategy-driven trades across FX and crypto markets. It combines the speed of **Golang** with the intelligence of **Python-based AI agents**, and gives users the ability to deploy their own smart trading bots — all running on our underlying infrastructure.

## 🚀 Features

- ⚡ **Real-time arbitrage execution** across multiple FX/Crypto platforms
- 🧮 **Concurrent trade processing** using Go routines and WaitGroups
- 🧠 **User-deployable AI agents** powered by real ML strategies
- 🗺️ **Multi-exchange support** with modular integration
- 🔧 **Tested and extensible architecture**
- 📊 **API-first** approach for managing trades, strategies, and logs
- 💬 **Web dashboard** for configuring agents and monitoring performance (in progress)

## 🛠️ Tech Stack

### 🧩 API

- **Go (Golang)** – for high-performance concurrent web backend and trade execution
- **pgx** – PostgreSQL driver and query builder
- **Goose** – for DB schema migrations
- **Gorilla/mux** – for routing
- **Socket.IO** – for live updates/monitoring

### 🗃️ Database

- **PostgreSQL** – reliable, high-performance relational database (SQL)
- **JSONB columns** – for flexible storage of semi-structured strategy configs or agent metadata
- **Indexes & constraints** – for efficient querying and data integrity (e.g. composite keys on trades)
- **Time-series data handling** – optimized schema for storing and querying high-frequency price or trade data
- **Dockerized setup** – for reproducible local development with pgAdmin support

### 🤖 AI Agent Engine

- **Python** – for powering intelligent agent behavior
- **pandas & NumPy** – for processing market data streams
- **scikit-learn / XGBoost / TensorFlow** – for predictive models and reinforcement learning agents
- **threading / multiprocessing** – for concurrent agent strategy execution
- **WebSocket clients** – for live market feeds and trigger-based trading
- **Custom orchestration** – to manage AI agents per user with sandboxed execution

### 🖥️ Frontend (Web)

- **React v18** – for building the user dashboard
- **Tailwind CSS** – for styling
- **TypeScript** – for type safety
- **React Query** – for API state and data fetching
- **MobX** – for UI state management

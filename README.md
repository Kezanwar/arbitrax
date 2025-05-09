# ArbitraX

## 🧠 About ArbitraX

**ArbitraX** is a high-performance algorithmic trading platform designed to execute real-time currency arbitrage strategies. It's a developer playground and technical challenge space for exploring **Golang**, **concurrent systems**, and **financial APIs**. It’s also a production-grade system designed with scalability and extensibility in mind.

## 🚀 Features

- ⚡ **Real-time arbitrage execution** across multiple FX/Crypto platforms
- 🧮 **Concurrent trade processing** using Go routines and WaitGroups
- 🗺️ **Multi-exchange support** with modular integration
- 🔧 **Tested and extensible architecture**
- 📊 **API-first** approach for managing trades, strategies, and logs
- 💬 **Web dashboard** for monitoring activity (in progress)

## 🛠️ Tech Stack

### API

- **Go (Golang)** – for high-performance backend and trade execution
- **pgx** – PostgreSQL driver and query builder
- **Goose** – for DB schema migrations
- **Gorilla/mux** – for routing
- **Socket.IO** – for live updates/monitoring
- **Docker** – for containerized development

### AI Trading Engine

### AI Trading Engine

- **Python** – for research, prototyping, and production-grade trading strategies
- **pandas & NumPy** – for time-series data manipulation and feature engineering
- **scikit-learn / XGBoost / TensorFlow** – for training predictive models on historical market data
- **joblib / multiprocessing / threading / concurrent.futures** – for executing and evaluating multiple strategies in parallel
- **ta / backtrader / vectorbt** – for technical analysis and backtesting pipelines
- **WebSocket clients (e.g. `websockets`, `aiohttp`)** – for live price feeds and event-driven signal generation
- **Jupyter Notebooks** – for exploratory development, model tuning, and visualization

### Frontend (Web)

- **React v19** – for building an admin interface
- **Tailwind CSS** – for styling
- **TypeScript** – for type safety
- **React Query** – for API data management
- **MobX** – for API data management

# One Step GPS Take Home

## Overview

This project visualizes devices on a Google Map, allowing users to interact with device markers, view detailed information, and manage device preferences such as visibility and custom photos. The backend handles data fetching and provides REST APIs, while the frontend is built with Vue.js and integrates with Google Maps API.

---

## Features

- **Interactive Map:**
  - Displays device locations with custom markers.
  - Display device information by clicking or hovering on markers
  - Dynamic marker updates based on device states and user preferences (e.g. device photos, device states, devices visibility).
- **Device Management:**
  - Toggle device visibility.
  - Upload custom photos for devices.
  - Reorder device list preferences by drag and drop.
  - Device management via context menu.
  - User authorization.
  - Persistent user preferences stored in the database and secured through JWT authorization.
- **Data Polling:**
  - Regularly fetches the latest device data.

---

## Tech Stack

### Frontend
- **Framework:** Vue.js
- **State Management:** Pinia
- **Styling:** Tailwind CSS
- **Google Maps Integration:** Google Maps JavaScript API

### Backend
- **Language:** Go (Golang)
- **Database:** MySQL
- **Environment Management:** Docker Compose

---

## Prerequisites

- **Node.js:** v18+
- **Golang:** v1.20+
- **Docker:** Installed and running.
- **Google Maps API Key:** Required for map functionality.

---

## Installation and Setup

### 1. Clone the Repository
```bash
git clone https://github.com/your-repository/OSGPSTH.git
cd OSGPSTH
```

### 2. Environment Variables

**Backend (backend/.env.docker):**

OSGPS_API_KEY=your_api_key
DB_USER=root
DB_PASSWORD=password
DB_HOST=db
DB_PORT=3306
DB_NAME=OSGPSTH
JWT_SECRET=your_secret

**Frontend (backend/.env):**

VITE_API_URL=http://localhost:8080/api
VITE_GOOGLE_MAPS_API_KEY=your_google_maps_api_key

### 3. Docker Setup and Deployment

**Build and start the project using Docker Compose.**

```bash
docker-compose up --build
```

**The application will be available at:**

- **Frontend:** http://localhost:5173
- **Backend API:** http://localhost:8080

**If local deployment is prefered:**

**A separate .env for local backend (backend/.env):**

OSGPS_API_KEY=your_api_key
DB_USER=root
DB_PASSWORD=password
DB_HOST=db
DB_PORT=3306
DB_NAME=OSGPSTH
JWT_SECRET=your_secret

**Frontend:**

```bash
cd frontend
npm install
npm run dev
```

**Backend:**

```bash
cd backend
go run cmd/main.go
```

## Contact

**For questions or suggestions, feel free to reach out**:

Justin Andersen
hire.justin.andersen@gmail.com
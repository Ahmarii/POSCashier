# POSCashier

POSCashier is a **desktop Point-of-Sale (POS) system** written in Go, designed for seamless integration with both **hardware devices** and **web services**. The project includes a custom UI built with the Corgent Core GUI library, integration with an **RFID reader** using TinyGo, and a **WebSocket-based barcode scanner** interface.

---

## 🔧 Features

- 🖥️ **Desktop App** with modular design (Logic, Data, UI)
- 📡 **WebSocket Integration** for external device communication (e.g., barcode scanner)
- 📶 **RFID Support** using Raspberry Pi Pico and MFRC522 module
- 🔄 Real-time cart management and UI updates
- 🧾 Order processing and PDF invoice generation
- 📁 Clean code structure for easy development and scaling

---

## 🧠 Architecture Overview

- `Logic/` – Core application logic and state management  
- `Data/` – Handles API calls and data fetching from web server  
- `UI/` – Frontend components using Corgent Core UI  

The app separates concerns into **Logic**, **Data**, and **UI**, allowing clean interaction between systems and maintainability.

---

## 📦 Hardware Integration

### RFID Reader
- Uses **Raspberry Pi Pico**
- Programmed with **TinyGo**
- Communicates with MFRC522 RFID module
- Sends UID to the desktop app via serial

Custom RFID library:  
🔗 [RFID-TinyGo on GitHub](https://github.com/Ahmarii/RFID-TinyGo)

---

### External WebSocket Barcode Scanner
- Barcode data is sent to the desktop app using the WebSocket protocol
- Handled by an internal WebSocket service in the app

---

## 🚀 Getting Started

### Requirements

- Go 1.20+
- TinyGo (for compiling firmware to Raspberry Pi Pico)
- Corgent Core GUI library
- A web server backend with API endpoints for:
  - Fetching products
  - Fetching customer by UID
  - Submitting orders

### Run Desktop App

```bash
go run main.go

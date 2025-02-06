# **Receipt Processor API** 🚀  
A lightweight web service that processes store receipts, assigns points based on predefined rules, and allows users to retrieve points for a given receipt ID.  

## **Table of Contents**  
- [Overview](#overview)  
- [Features](#features)  
- [Technology Stack](#technology-stack)  
- [Installation](#installation)  
- [Running the Application](#running-the-application)  
- [API Endpoints](#api-endpoints)  
  - [1. Process Receipt](#1-process-receipt)  
  - [2. Get Points](#2-get-points)  
- [Point Calculation Rules](#point-calculation-rules)  
- [Example Requests](#example-requests)  
- [Docker Support (Optional)](#docker-support-optional)  
- [License](#license)  
- [Contact](#contact)  

---

## **Overview**  
The **Receipt Processor API** is a RESTful web service that processes JSON-based store receipts and calculates points based on specific business rules. The points can later be retrieved using a unique receipt ID.  

This project was developed to showcase backend development skills, including **REST API design, JSON handling, in-memory data storage, and Go programming best practices**.  

---

## **Features** ✅  
✔ **Submit a receipt** and receive a unique ID.  
✔ **Retrieve points** awarded for a given receipt.  
✔ **In-memory storage** (no external database required).  
✔ **Follows RESTful principles** and uses JSON-based communication.  
✔ **Built with Go** for high performance and efficiency.  

---

## **Technology Stack** 🛠  
- **Language:** Go (Golang)  
- **Framework:** Gorilla Mux (for routing)  
- **UUID Generator:** Google UUID Package  
- **Data Storage:** In-memory (using Go Maps)  

---

## **Installation** 💻  

### **1. Prerequisites**  
- Install **Go** (if not already installed)  
  📥 Download Go from: [https://go.dev/dl/](https://go.dev/dl/)  
  Verify installation:  
  ```sh
  go version
  ```

### **2. Clone the Repository**  
```sh
git clone https://github.com/your-username/receipt-processor.git
cd receipt-processor
```

### **3. Install Dependencies**  
```sh
go mod tidy
```

---

## **Running the Application** ▶️  
To start the server, run:  
```sh
go run main.go
```
The server will start at:  
🔗 `http://localhost:8080`

---

## **API Endpoints**  

### **1. Process Receipt**  
📌 **Endpoint:**  
```
POST /receipts/process
```
📌 **Description:**  
Submits a receipt for processing and returns a unique ID.

📌 **Request Body (JSON Example):**  
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-02",
  "purchaseTime": "13:13",
  "total": "1.25",
  "items": [
    {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
  ]
}
```

📌 **Response Example:**  
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

---

### **2. Get Points**  
📌 **Endpoint:**  
```
GET /receipts/{id}/points
```
📌 **Description:**  
Retrieves the points awarded for a given receipt ID.

📌 **Response Example:**  
```json
{ "points": 10 }
```

---

## **Point Calculation Rules** 🎯  
The points awarded for each receipt follow these rules:  

1️⃣ **Retailer Name:** One point per alphanumeric character.  
2️⃣ **Total Amount:**  
   - 50 points if the total is a whole number.  
   - 25 points if the total is a multiple of 0.25.  
3️⃣ **Number of Items:** 5 points for every two items.  
4️⃣ **Item Description:** If the description length is a multiple of 3, `ceil(item_price * 0.2)` points are awarded.  
5️⃣ **Purchase Date:** 6 points if the purchase day is **odd**.  
6️⃣ **Purchase Time:** 10 points if the purchase time is **between 2:00 PM and 4:00 PM**.  

---

## **Example Requests** 📝  

### **Process Receipt (cURL)**
```sh
curl -X POST "http://localhost:8080/receipts/process" \
     -H "Content-Type: application/json" \
     -d '{
         "retailer": "Target",
         "purchaseDate": "2022-01-02",
         "purchaseTime": "13:13",
         "total": "1.25",
         "items": [
             {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
         ]
     }'
```

### **Get Points (cURL)**
```sh
curl -X GET "http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points"
```

---

## **Contact** ✉️  
If you have any questions, feel free to reach out:  
📧 **Email:** [kruthik@jobhuntmails.com](mailto:kruthik@jobhuntmails.com)  
🔗 **LinkedIn:** [Kruthik Kumar](https://www.linkedin.com/in/kruthik-kumar-vandana/)  

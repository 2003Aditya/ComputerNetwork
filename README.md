# Network Communication Simulator 🌐

This project is a simple simulator for a computer network, written in **Go**. It demonstrates fundamental concepts of data transmission, focusing on how information is structured, sent, and received over a simulated medium. It’s perfect for understanding the basics of **data framing** and **error detection**.

---

## 🌟 Features

- **Framing:** Encapsulates data into frames, complete with start and end flags, and a parity bit for error checking.
- **Packet Structure:** Simulates a basic packet header with fields like source and destination addresses, TTL (Time-to-Live), sequence and acknowledgment numbers, and flags (e.g., SYN, ACK, FIN) to manage communication flow.
- **Bit-by-Bit Transfer:** Transmission is simulated by writing and reading data **one bit at a time** to and from a file named `wire.txt`, mimicking a physical connection.
- **Parity Check:** The receiver program uses a parity bit to detect single-bit errors in the transmitted data, providing a basic form of error detection.

---

## 🚀 How to Run

### Prerequisites

You need to have **Go** installed on your system.
The code was tested with **Go version 1.18**.

### Execution

1. **Clone the repository:**

```bash
git clone https://github.com/2003Aditya/ComputerNetwork.git
cd ComputerNetwork
```
Run the sender:

The sender program creates the wire.txt file and writes the simulated bitstream to it.

```bash

go run sender.go
```
Run the receiver:

The receiver program reads the bitstream from wire.txt, reconstructs the frames, and processes the packets.

```bash

go run receiver.go
```
You will see output in your terminal as the bits are "transmitted" and "received," along with the decoded packet information and parity check results.
```

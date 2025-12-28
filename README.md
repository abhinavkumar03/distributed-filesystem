<p align="center">
  <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" width="120" alt="Go Gopher">
</p>

<h1 align="center">Distributed File System (Go-DFS)</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/Network-TCP%20%2F%20P2P-blueviolet?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Architecture-Distributed-blue?style=for-the-badge" />
</p>

---

## 🌐 Overview
A high-performance, **Peer-to-Peer (P2P)** distributed storage engine designed in Go. This system is engineered to store and serve massive files by fragmenting them into chunks and distributing them across a fault-tolerant cluster of nodes.

This project demonstrates core systems engineering principles: **Custom TCP protocols, Content-Addressable Storage, and Concurrency.**

---

## 🛠 Technical Highlights

| Feature | Technical Implementation |
| :--- | :--- |
| <img src="https://cdn-icons-png.flaticon.com/512/3665/3665923.png" width="20"> **Transport** | **Custom TCP Protocol:** Built from scratch to handle low-latency peer-to-peer communication without the overhead of HTTP. |
| <img src="https://cdn-icons-png.flaticon.com/512/2082/2082805.png" width="20"> **Concurrency** | **Goroutines & Channels:** Leverages Go’s primitives for non-blocking I/O and handling thousands of concurrent data streams. |
| <img src="https://cdn-icons-png.flaticon.com/512/2784/2784065.png" width="20"> **Storage** | **Content Addressable Storage (CAS):** Files are indexed by their hash, ensuring data integrity and automatic deduplication. |
| <img src="https://cdn-icons-png.flaticon.com/512/2592/2592317.png" width="20"> **Network** | **P2P Architecture:** Decentralized design where nodes can act as both clients and servers to eliminate bottlenecks. |

---

## 🏗 System Architecture & Scalability



### How it Scales
1. **Horizontal Scaling:** New storage nodes can join the cluster dynamically. The system uses a handshake protocol over TCP to register new capacity without downtime.
2. **Chunking Strategy:** Instead of storing a 10GB file on one machine, it is split into smaller segments (chunks). These are distributed across the network, allowing for **parallel uploads/downloads** which maximize bandwidth.
3. **Replication Logic:** To ensure 99.99% availability, each chunk is replicated across multiple physical nodes. If Node A fails, the system automatically redirects the request to Node B.

---

## 🧠 Engineering Challenges Solved

* **Memory Efficiency:** Implemented `io.Reader` and `io.Writer` interfaces to stream data directly from the network to the disk. This prevents "Out of Memory" errors when handling files larger than the system's RAM.
* **Protocol Design:** Designed a custom binary protocol for peer handshaking and file transmission, significantly reducing header overhead compared to REST or JSON-based APIs.
* **State Consistency:** Managed the mapping between file hashes and physical locations across a distributed network.

---

## 🚀 Future Roadmap
- [x] TCP Transport Layer & Peer Handshaking
- [x] File Encryption at Rest
- [ ] **Next:** Implementing Raft or Paxos for distributed consensus.
- [ ] **Next:** Adding a FUSE interface to mount the DFS as a local drive.

---

## 📄 License
Distributed under the MIT License. See `LICENSE` for more information.

---
<p align="center">
  Built with ❤️ for Distributed Systems Research
</p>

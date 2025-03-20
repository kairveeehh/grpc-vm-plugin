
## 📌 Project Description
This project demonstrates a simple **gRPC-based plugin architecture** in Go, where a **controller** binary starts a **plugin** that:
- **Starts** a mock VM
- **Stops** the mock VM
- **Fetches VM Info**

The plugin simulates VM operations and communicates over gRPC.

---

## 📂 Project Structure
```
grpc-vm-plugin/
├── go.mod
├── proto/
│   └── vm.proto        # gRPC service definitions
├── plugin/
│   └── main.go         # gRPC server (mock VM operations)
├── main.go             # gRPC client (controller)
```

---

## ✅ Local Setup Instructions

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/kairveeehh/grpc-vm-plugin.git
cd grpc-vm-plugin
```

### 2️⃣ Install Required Go Modules
```bash
go mod tidy
```

### 3️⃣ Install `protoc` and Go plugins (if not already installed)
```bash
sudo apt update
sudo apt install -y protobuf-compiler

# Install Go protobuf plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Add Go bin to PATH (if not already)
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 4️⃣ Generate Go gRPC Code from Proto
```bash
protoc --go_out=. --go-grpc_out=. proto/vm.proto
```

### 5️⃣ Build the Plugin (gRPC Server)
```bash
cd plugin
go build -o vmplugin
cd ..
```

### 6️⃣ Build the Controller (gRPC Client)
```bash
go build -o controller
```

### 7️⃣ Run the Controller (It will start the plugin and interact via gRPC)
```bash
./controller
```

---

## 🔧 Example Output (Screenshot Example)
```
2025/03/21 10:15:20 Start Response: VM Started Successfully
2025/03/21 10:15:20 Info Response: VM is Running
2025/03/21 10:15:20 Stop Response: VM Stopped Successfully
```

---

## 🔹 Further Use Case Extension - Real QEMU VM Control
This mock plugin structure can be extended to control **real VMs using QEMU**:

- Replace the mock `Start`, `Info`, `Stop` logic in the plugin with **QMP (QEMU Monitor Protocol)** commands.
- Launch actual QEMU instances.
- Interact with running VMs (check status, shut down, query stats).

📅 **Example:**
```go
cmd := exec.Command("qemu-system-x86_64", "-qmp", "tcp:localhost:4444,server,nowait", "-hda", "alpine.qcow2")
```

## 👉 Author
GitHub: [kairveeehh](https://github.com/kairveeehh)

---


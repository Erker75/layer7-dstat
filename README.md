# ⚡️ High-Performance Layer 7 Dstat

A lightweight, ultra-fast web dashboard to monitor your HTTP/HTTPS Requests Per Second (RPS) in real-time. Built from the ground up in Go, this dstat uses atomic counters and efficient memory management to effortlessly monitor 100k+ req/s without crashing your server!

![Layer7 Dstat](https://i.ibb.co/WWDf1tPk/3097a422754cb2b2394520dc8119.png)

## 🔥 Features
* **Ultra-Fast Backend:** Written purely in Go. Zero-lock atomic counters handle extreme load (tested at 300k+ req/s). 
* **Live Monitoring:** Real-time animated Chart.js graphical analytics with digital glitches under high load.
* **1-Click Copy:** Click the target URL to copy your host endpoint smoothly.
* **Zero Dependencies:** No databases or clunky UI frameworks required—just one executable file!

## 🚀 Simple VPS Setup Guide (Ubuntu/Debian)

**Step 1:** Connect to your VPS via SSH.
**Step 2:** Install Golang on your server:
```bash
sudo apt update && sudo apt install golang -y
```
**Step 3:** Upload and unzip the source files to your server (or `git clone` the repository).
**Step 4:** Navigate to the folder and start the monitor in the background:
```bash
cd layer7-dstat
go mod init layer7-dstat
nohup go run main.go &
```
**Step 5:** You're done! Visit `http://SERVER_IP:8080/` in your browser to view the live Dstat panel.

*Target endpoint for stress testing is: `http://SERVER_IP:8080/target`*

## ☕ Support / Donations

If you found this tool helpful and want to support further development, donations are heavily appreciated:

* **BTC:** `bc1qtkm7dzjp76gx8t9c02pshfd8rzarj6gj9yzglu`
* **ETH:** `0x88Aa0E09a5A62919321f38Fb4782A17f4dc91A9B`
* **XMR:** `0x6730c52B3369fD22E3ACc6090a3Ee7d5C617aBE0`

---
*Developed by [@Triotion](https://triotion.t.me)*

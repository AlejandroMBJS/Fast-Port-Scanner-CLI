package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"strings"
	"io/ioutil"
)

func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func checkPort(ip string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func getServiceName(port int) string {
	services := map[int]string{
		21:   "ftp",
		22:   "ssh",
		23:   "telnet",
		25:   "smtp",
		53:   "dns",
		80:   "http",
		110:  "pop3",
		143:  "imap",
		443:  "https",
		3306: "mysql",
		5432: "postgresql",
		6379: "redis",
		8080: "http-alt",
	}
	if service, exists := services[port]; exists {
		return service
	}
	return "unknown"
}

func customUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("\nDescription:")
	fmt.Println("This program scans ports on a given IP address and shows open ports with the associated service.")
	fmt.Println("Required options:")
	fmt.Println("  -ip <IP address> : Specify the IP address to scan.")
	fmt.Println("\nOptional options:")
	fmt.Println("  -timeout <milliseconds> : Timeout in milliseconds for each connection attempt. Default value: 500 ms.")
	fmt.Println("  -o <filename> : Output filename to export results.")
	fmt.Println("\nExample usage:")
	fmt.Println("  sudo goscan -ip 127.0.0.1 -timeout 1000 -o results")
}

func loadingAnimation(done chan bool) {
	loadingChars := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\rScanning ports... %s", loadingChars[i])
			i = (i + 1) % len(loadingChars)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func printProgressBar(current, total int) {
	percent := float64(current) / float64(total) * 100
	bars := int(percent / 5)
	fmt.Printf("\r[%-20s] %3.0f%% completed", strings.Repeat("=", bars), percent)
}

func main() {
	var ip string
	var timeout int
	var outputFile string
	flag.StringVar(&ip, "ip", "", "IP address to scan (required)")
	flag.IntVar(&timeout, "timeout", 500, "Timeout in milliseconds for connection")
	flag.StringVar(&outputFile, "o", "", "Output filename to export results")

	flag.Usage = customUsage
	flag.Parse()

	if ip == "" || !isValidIP(ip) {
		color.Red("Error: a valid IP address to scan must be provided.")
		flag.Usage()
		os.Exit(1)
	}

	color.Cyan("Scanning ports on %s...\n", ip)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	var mu sync.Mutex

	startTime := time.Now()

	go func() {
		<-stopChan
		color.Yellow("\nScan interrupted by the user.")
		os.Exit(0)
	}()

	done := make(chan bool)
	go loadingAnimation(done)

	const maxGoroutines = 5000
	guard := make(chan struct{}, maxGoroutines)
	totalPorts := 65535

	var openPorts []string
	var openPortNumbers []string

	for port := 1; port <= totalPorts; port++ {
		wg.Add(1)
		guard <- struct{}{}

		go func(port int) {
			defer wg.Done()
			defer func() { <-guard }()

			if checkPort(ip, port, time.Duration(timeout)*time.Millisecond) {
				service := getServiceName(port)
				mu.Lock()
				openPorts = append(openPorts, fmt.Sprintf("Port %d is open (%s)", port, service))
				openPortNumbers = append(openPortNumbers, fmt.Sprintf("%d", port))
				mu.Unlock()
			}

			mu.Lock()
			printProgressBar(port, totalPorts)
			mu.Unlock()
		}(port)
	}

	wg.Wait()

	done <- true
	fmt.Println()

	elapsedTime := time.Since(startTime)

	if len(openPorts) > 0 {
		color.Cyan("\nOpen ports found:")
		for _, portInfo := range openPorts {
			fmt.Println(portInfo)
		}
		if outputFile != "" {
			// Preparar la línea nmap
			nmapLine := fmt.Sprintf("nmap -sCV -p%s %s", strings.Join(openPortNumbers, ","), ip)
			openPorts = append(openPorts, "\n"+nmapLine)

			// Escribir los resultados y la línea de nmap en el archivo
			err := ioutil.WriteFile(outputFile, []byte(strings.Join(openPorts, "\n")), 0644)
			if err != nil {
				color.Red("Error writing to file:", err)
			} else {
				color.Cyan("Results exported to %s\n", outputFile)
			}
		}
	} else {
		color.Yellow("No open ports found.")
	}

	color.Cyan("\nScan completed in %s\n", elapsedTime)
}

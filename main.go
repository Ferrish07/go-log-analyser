package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	// CLI flag for log file
	logFile := flag.String("file", "auth.log", "Path to the log file to analyze")
	flag.Parse()

	// Open the file
	file, err := os.Open(*logFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Maps for counting attempts
	ipCount := make(map[string]int)
	userCount := make(map[string]int)

	// Scan the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if containsFailedSSH(line) {
			user, ip := extractUserAndIP(line)
			if user != "" && ip != "" {
				ipCount[ip]++
				userCount[user]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Print results
	fmt.Println("\n=== SSH Brute Force Summary ===")

	fmt.Println("\nTop IPs:")
	for ip, count := range ipCount {
		fmt.Printf("  %s: %d attempts\n", ip, count)
	}

	fmt.Println("\nTop Usernames:")
	for user, count := range userCount {
		fmt.Printf("  %s: %d attempts\n", user, count)
	}
}

// Helper: check if line has failed ssh login
func containsFailedSSH(line string) bool {
	return strings.Contains(line, "Failed password") && strings.Contains(line, "sshd")
}

// Extract user + IP using regex
func extractUserAndIP(line string) (string, string) {
	re := regexp.MustCompile(`Failed password for (invalid user )?(\w+) from ([\d.]+)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) >= 4 {
		user := matches[2]
		ip := matches[3]
		return user, ip
	}
	return "", ""
}

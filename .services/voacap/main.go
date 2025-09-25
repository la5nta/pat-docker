package main

import (
	"bytes"
	"cmp"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	server := &VoacapServer{
		dataDir: os.ExpandEnv(cmp.Or(os.Getenv("VOACAP_DATA_DIR"), "$HOME/itshfbc")),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/predict", server.handlePredict)
	mux.HandleFunc("/version", server.handleVersion)
	httpServer := &http.Server{
		Addr:    cmp.Or(os.Getenv("VOACAP_ADDR"), ":8080"),
		Handler: mux,
	}

	log.Printf("Listening %s...", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type VoacapServer struct{ dataDir string }

func (s *VoacapServer) handleVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cmd := exec.Command("voacapl", "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get version: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(bytes.TrimSpace(output))
}

func (s *VoacapServer) handlePredict(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create temporary run directory
	runDir, err := os.MkdirTemp("", "voacap-api-")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create temp directory: %v", err), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(runDir)

	// Write request body as input file
	in, err := os.OpenFile(filepath.Join(runDir, "input.dat"), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create input file: %v", err), http.StatusInternalServerError)
		return
	}
	defer in.Close()
	switch n, err := io.Copy(in, r.Body); {
	case err != nil:
		http.Error(w, fmt.Sprintf("Failed to read request body: %v", err), http.StatusBadRequest)
		return
	case n == 0:
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}
	in.Close()

	// Execute VOACAP
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()
	args := []string{"--run-dir=" + runDir, s.dataDir, "input.dat", "output.out"}
	cmd := exec.CommandContext(ctx, "voacapl", args...)
	cmd.Dir = runDir
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Printf("VOACAP execution failed: %v\nOutput: %s", err, string(output))
		http.Error(w, fmt.Sprintf("VOACAP execution failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Write output as response
	out, err := os.Open(filepath.Join(runDir, "output.out"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read output file: %v", err), http.StatusInternalServerError)
		return
	}
	defer out.Close()
	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, out)
}

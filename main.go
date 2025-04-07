package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/vin-rmdn/imagesort-go/cmd"
)

func main() {
	setupLogger()

	command := &cobra.Command{Short: "imagesort-go"}
	
	imageSortCommand, err := cmd.ImageSortCommand()
	if err != nil {
		log.Fatalf("failed to init image sort command: %v", err)
	}

	command.AddCommand(imageSortCommand)
	command.AddCommand(cmd.RenameCommand())

	if err := command.ExecuteContext(context.Background()); err != nil {
		log.Fatalf("command execution failed: %v", err)
	}
}

func setupLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
	}))

	slog.SetDefault(logger)
}
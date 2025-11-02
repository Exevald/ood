package main

import (
	"os"

	"editor/pkg/model"
	"editor/pkg/service"
)

func main() {
	workDir, _ := os.Getwd()
	doc := model.NewDocument(workDir)
	svc := service.NewDocumentService(doc)
	handler := NewCommandsHandler(svc)
	handler.Run()
}

package main

import (
	"context"
	"log"
	"os"

	"github.com/GhandyP/adk-go-enterprise-cookbook/internal/adkhelper"
)

func main() {
	if err := adkhelper.Run(
		context.Background(),
		os.Args[1:],
		"case_45_ab_copy_variations",
		"A text-only starting point for a/b copy variations.",
		"Create distinct copy variations from the campaign brief supplied by the user. Give each a hypothesis, avoid duplicates, and do not invent product claims.",
	); err != nil {
		log.Fatal(err)
	}
}

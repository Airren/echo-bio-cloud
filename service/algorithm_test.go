package service

import (
	"context"
	"github.com/airren/echo-bio-backend/test_suit"
	"os"
	"testing"
)

func init() {
	test_suit.TestInit()
}

func TestCreateAlgorithm(t *testing.T) {
	file, err := os.Open("../conf/venn.yaml")
	if err != nil {
		t.Fatal("invalid file path")
	}
	if err := CreateAlgorithm(context.TODO(), file); err != nil {
		t.Fatal("create algorithm failed:", err)
	}
	t.Logf("create algorithm success")
}

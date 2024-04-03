package ifile

import (
	"context"
	"fmt"
	"testing"
)

func TestFileAttract_Attract(t *testing.T) {
	ctx := context.Background()
	fmt.Println(New(ctx, `/e/www/Go/idea-space/idea-space-admin/resource/uploads/d09fmfjxtei850bb3b.jpg`).MustHasherString())
}

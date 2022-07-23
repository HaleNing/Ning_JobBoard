package delegate

import (
	"context"
	"fmt"
	"github.com/HaleNing/bustrack/src/Model/ent"
	"log"
)

func CreateBook(ctx context.Context, client *ent.Client) (*ent.Book, error) {
	res, err := client.Book.Create().SetAuthor("lining02xxx").Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating book: %w", err)
	}
	log.Println("book was created: ", res)
	return res, nil
}

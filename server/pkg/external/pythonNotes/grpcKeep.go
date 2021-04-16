package pythonNotes

import (
	"context"
	"fmt"
	"time"

	keepServiceApi "github.com/luckyshmo/api-example/api/keep"
	"github.com/luckyshmo/api-example/models/keep"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GrpcKeepData struct {
	Client keepServiceApi.KeepClient
}

func NewKeepClient(address string) (*GrpcKeepData, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't connect to %s", address))
	}
	client := keepServiceApi.NewKeepClient(conn)
	return &GrpcKeepData{client}, err
}

func (gkd *GrpcKeepData) GetAll() (keep.Note, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := gkd.Client.GetWords(ctx, &keepServiceApi.SearchRequest{
		Email: "mishka2017@gmail.com",
		Token: "iopnilguhgigbbht",
		Name:  "English words",
	})
	if err != nil {
		return keep.Note{}, errors.Wrap(err, "Can't get words")
	}
	wordsList := r.GetWord()

	return keep.Note{Data: wordsList}, nil
}

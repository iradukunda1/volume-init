package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/containerd/continuity/fs"
	"github.com/iradukunda1/volume-init/volume"
)

func main() {
	err := realMain()
	if err != nil {
		out := volume.GuestVolumeImageOutput{Error: err.Error()}
		b, err := json.Marshal(out)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to unmarshal %+v: %s", out, err)
			os.Exit(2)
		}
		fmt.Printf("%s", string(b))
		os.Exit(1)
	}
}

func copy(in volume.GuestVolumeImageInput) error {

	for _, v := range in.Volumes {
		to, err := fs.RootPath(in.To, v)
		if err != nil {
			return fmt.Errorf("failed to join %q and %q: %w", in.To, v, err)
		}

		from, err := fs.RootPath(in.From, v)
		if err != nil {
			return fmt.Errorf("failed to join %q and %q: %w", in.From, v, err)
		}

		err = os.MkdirAll(to, 0700)
		if err != nil {
			return err
		}

		err = fs.CopyDir(to, from)
		if err != nil {
			return err
		}
	}
	return nil
}

func realMain() error {

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	var in volume.GuestVolumeImageInput
	err = json.Unmarshal(b, &in)
	if err != nil {
		return err
	}

	return copy(in)
}

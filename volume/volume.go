package volume

// Volume is a special directory that could be shared between containers.
type Volume struct {
	name          string
	hostPath      string
	vmPath        string
	containerPath string
}

// New returns an empty volume.
func New(name string) *Volume {
	return &Volume{name: name}
}

// FromHost returns a volume which has the files from the given path.
func FromHost(name, path string) *Volume {
	return &Volume{name: name, hostPath: path}
}

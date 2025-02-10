package id

type EpochOptions struct {
	MachineId     int
	MachineIdBits int
	SeqBits       int
}

func NewSingleEpoch(machineId int) *EpochOptions {
	return &EpochOptions{
		MachineId:     machineId,
		MachineIdBits: 10,
		SeqBits:       10,
	}
}

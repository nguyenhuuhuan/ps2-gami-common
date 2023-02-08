package cache

type BitFieldModel struct {
	Command string
	Byte    string
	Offset  int
}

func (m *BitFieldModel) convertCommand() []interface{} {
	var command = make([]interface{}, 3)
	command[0] = m.Command
	command[1] = m.Byte
	command[2] = m.Offset
	return command
}

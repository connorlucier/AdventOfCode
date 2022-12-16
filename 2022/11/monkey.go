package main

type monkey struct {
	ID           string
	Items        []int
	Operation    string
	Operand      int
	Divisor      int
	Inspections  int
	FirstTarget  *monkey
	SecondTarget *monkey
}

func (m *monkey) Act(lcm int, part1 bool) {
	for len(m.Items) > 0 {
		m.Inspections++

		if m.Operation == "**" {
			m.Items[0] *= m.Items[0]
		} else if m.Operation == "*" {
			m.Items[0] *= m.Operand
		} else if m.Operation == "+" {
			m.Items[0] += m.Operand
		}

		m.Items[0] %= lcm
		if part1 {
			m.Items[0] /= 3
		}

		if m.Items[0]%m.Divisor == 0 {
			m.FirstTarget.Items = append(m.FirstTarget.Items, m.Items[0])
		} else {
			m.SecondTarget.Items = append(m.SecondTarget.Items, m.Items[0])
		}
		m.Items = append(m.Items[:0], m.Items[1:]...)
	}
}

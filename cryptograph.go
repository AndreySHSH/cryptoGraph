package cryptoGraph

type Accesses struct {
	ID  int64
	Key string

	PathFileMap string // distillation file bd
}

type Cryptos struct {
}

func Init(accesses Accesses) (*Cryptos, error) {
	c := &Cryptos{}
	_, err := c.decipher(accesses)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cryptos) decipher(accesses Accesses) (bool, error) {
	return true, nil
}

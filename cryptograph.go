package cryptoGraph

type Access struct {
	ID  int64
	Key string

	PathFileMap string // distillation file map db
}

type Cryptos struct {
}

func Init(access Access) (*Cryptos, error) {
	c := &Cryptos{}
	_, err := c.decipher(access)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cryptos) decipher(access Access) (bool, error) {
	return true, nil
}

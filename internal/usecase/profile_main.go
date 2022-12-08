package usecase

type PrflUseCase struct {
	ProfileKafka
	ProfilePostgres
}

func New(k ProfileKafka, p ProfilePostgres) *PrflUseCase {
	return &PrflUseCase{
		ProfileKafka:    k,
		ProfilePostgres: p,
	}
}

func (p *PrflUseCase) TransportData() error {

	datas, err := p.ProfilePostgres.GetProfiles()
	if err != nil {
		return err
	}

	for _, val := range datas {
		p.ProfileKafka.Write(val)
	}
	return nil
}

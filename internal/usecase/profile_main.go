package usecase

import "ptok/internal/entity"

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

func (p *PrflUseCase) TransportDataV2() error {

	prflChan := make(chan entity.Profile)
	defer close(prflChan)

	if err := p.ProfilePostgres.GetProfilesV2(prflChan); err != nil {
		return err
	}

	p.ProfileKafka.WriteV2(prflChan)

	return nil
}

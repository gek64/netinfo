package main

import (
	"geolite2"
	"log"
)

func asnSearch(ip string, database string) (asn geolite2.AsnResult, err error) {
	asnDB, err := geolite2.OpenAsnDB(database)
	if err != nil {
		return geolite2.AsnResult{}, err
	}
	defer func(asnDB geolite2.AsnDB) {
		err := asnDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(asnDB)
	return asnDB.Search(ip)
}

func cityBlockSearch(ip string, database string) (cbResult geolite2.CityBlockResult, err error) {
	cbDB, err := geolite2.OpenCityBlockDB(database)
	if err != nil {
		return geolite2.CityBlockResult{}, err
	}
	defer func(cbDB geolite2.CityBlockDB) {
		err := cbDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(cbDB)
	return cbDB.Search(ip)
}

func cityLocationSearch(geoNameID int, database string) (clResult geolite2.CityLocationResult, err error) {
	clDB, err := geolite2.OpenCityLocationDB(database)
	if err != nil {
		return geolite2.CityLocationResult{}, err
	}
	defer func(clDB geolite2.CityLocationDB) {
		err := clDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(clDB)
	return clDB.Search(geoNameID)
}

func countryBlockSearch(ip string, database string) (cbResult geolite2.CountryBlockResult, err error) {
	cbDB, err := geolite2.OpenCountryBlockDB(database)
	if err != nil {
		return geolite2.CountryBlockResult{}, err
	}
	defer func(cbDB geolite2.CountryBlockDB) {
		err := cbDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(cbDB)
	return cbDB.Search(ip)
}

func countryLocationSearch(geoNameID int, database string) (clResult geolite2.CountryLocationResult, err error) {
	clDB, err := geolite2.OpenCountryLocationDB(database)
	if err != nil {
		return geolite2.CountryLocationResult{}, err
	}
	defer func(clDB geolite2.CountryLocationDB) {
		err := clDB.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(clDB)
	return clDB.Search(geoNameID)
}

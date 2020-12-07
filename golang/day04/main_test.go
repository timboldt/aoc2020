package main

import "testing"

func TestInvalidPart2(t *testing.T) {
	data := []*passport{
		&passport{
			kv: map[string]string{
				"eyr": "1972",
				"cid": "100",
				"hcl": "#18171d",
				"ecl": "amb",
				"hgt": "170",
				"pid": "186cm",
				"iyr": "2018",
				"byr": "1926",
			},
		},
		&passport{
			kv: map[string]string{
				"iyr": "2019",
				"hcl": "#602927",
				"eyr": "1967",
				"hgt": "170cm",
				"ecl": "grn",
				"pid": "012533040",
				"byr": "1946",
			},
		},
		&passport{
			kv: map[string]string{
				"hcl": "dab227",
				"iyr": "2012",
				"ecl": "brn",
				"hgt": "182cm",
				"pid": "021572410",
				"eyr": "2020",
				"byr": "1992",
				"cid": "277",
			},
		},
		&passport{
			kv: map[string]string{
				"hgt": "59cm",
				"ecl": "zzz",
				"eyr": "2038",
				"hcl": "74454a",
				"iyr": "2023",
				"pid": "3556412378",
				"byr": "2007",
			},
		},
		&passport{
			kv: map[string]string{
				"ecl": "hzl",
				"byr": "1971",
				"pid": "030850749",
				"hgt": "170in",
				"hcl": "#ceb3a1",
				"eyr": "2023",
				"iyr": "2018",
			},
		},
	}
	for _, p := range data {
		if p.isValid2() {
			t.Errorf("should not be valid: %v", p.kv)
		}
	}
}

func TestValidPart2(t *testing.T) {
	data := []*passport{
		&passport{
			kv: map[string]string{
				"pid": "087499704",
				"hgt": "74in",
				"ecl": "grn",
				"iyr": "2012",
				"eyr": "2030",
				"byr": "1980",
				"hcl": "#623a2f",
			},
		},
		&passport{
			kv: map[string]string{
				"eyr": "2029",
				"ecl": "blu",
				"cid": "129",
				"byr": "1989",
				"iyr": "2014",
				"pid": "896056539",
				"hcl": "#a97842",
				"hgt": "165cm",
			},
		},
		&passport{
			kv: map[string]string{
				"hcl": "#888785",
				"hgt": "164cm",
				"byr": "2001",
				"iyr": "2015",
				"cid": "88",
				"pid": "545766238",
				"ecl": "hzl",
				"eyr": "2022",
			},
		},
		&passport{
			kv: map[string]string{
				"iyr": "2010",
				"hgt": "158cm",
				"hcl": "#b6652a",
				"ecl": "blu",
				"byr": "1944",
				"eyr": "2021",
				"pid": "093154719",
			},
		},
	}
	for _, p := range data {
		if !p.isValid2() {
			t.Errorf("should be valid: %v", p.kv)
		}
	}
}

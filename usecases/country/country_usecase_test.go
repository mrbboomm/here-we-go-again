package usecases_test

// func TestGetCountries(t *testing.T){
//   countryData := []entities.CountryEntity{{Id: primitive.NewObjectID(), Name: "Thailand", Continent:"Asia"} , {Id: primitive.NewObjectID(), Name: "India", Continent:"Asia" }}

// 	countryRepo := mock.NewCountryRepoMock()
// 	countryRepo.On("FindAll").Return(countryData)

// 	countryUseCase := usecases.NewCountryUseCase(countryRepo)

// 	results := countryUseCase.GetCountries()

// 	expected := countryData

// 	assert.Equal(t, expected, results)
// }
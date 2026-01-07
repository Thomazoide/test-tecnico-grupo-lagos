package payloads

type ResponsePayload struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      bool   `json:"error"`
	Data       any    `json:"data,omitempty"`
}

type ProductsPayload struct {
	Count      int       `json:"count"`
	Page       int       `json:"page"`
	Page_count int       `json:"page_count"`
	Page_size  int       `json:"page_size"`
	Products   []Product `json:"products"`
	Skip       int       `json:"skip"`
}

type ProductByCodePayload struct {
	Status         int     `json:"status"`
	Status_verbose string  `json:"status_verbose"`
	Code           string  `json:"code"`
	Product        Product `json:"product"`
}

type Product struct {
	Brands           string    `json:"brands"`
	Code             string    `json:"code"`
	Image_url        string    `json:"image_url"`
	Nutriscore_grade string    `json:"nutriscore_grade"`
	Product_name     string    `json:"product_name"`
	Nutriments       Nutrition `json:"nutriments"`
}

type Nutrition struct {
	Carbohydrates100g                                     float64 `json:"carbohydrates_100g"`
	CarbohydratesServing                                  float64 `json:"carbohydrates_serving"`
	CarbohydratesUnit                                     string  `json:"carbohydrates_unit"`
	CarbohydratesValue                                    float64 `json:"carbohydrates_value"`
	CarbonFootprintFromKnownIngredients100g               float64 `json:"carbon-footprint-from-known-ingredients_100g"`
	CarbonFootprintFromKnownIngredientsProduct            float64 `json:"carbon-footprint-from-known-ingredients_product"`
	CarbonFootprintFromKnownIngredientsServing            float64 `json:"carbon-footprint-from-known-ingredients_serving"`
	Energy                                                float64 `json:"energy"`
	EnergyKcal                                            float64 `json:"energy-kcal"`
	EnergyKcal100g                                        float64 `json:"energy-kcal_100g"`
	EnergyKcalServing                                     float64 `json:"energy-kcal_serving"`
	EnergyKcalUnit                                        string  `json:"energy-kcal_unit"`
	EnergyKcalValue                                       float64 `json:"energy-kcal_value"`
	EnergyKcalValueComputed                               float64 `json:"energy-kcal_value_computed"`
	EnergyKj                                              float64 `json:"energy-kj"`
	EnergyKj100g                                          float64 `json:"energy-kj_100g"`
	EnergyKjServing                                       float64 `json:"energy-kj_serving"`
	EnergyKjUnit                                          string  `json:"energy-kj_unit"`
	EnergyKjValue                                         float64 `json:"energy-kj_value"`
	EnergyKjValueComputed                                 float64 `json:"energy-kj_value_computed"`
	Energy100g                                            float64 `json:"energy_100g"`
	EnergyServing                                         float64 `json:"energy_serving"`
	EnergyUnit                                            string  `json:"energy_unit"`
	EnergyValue                                           float64 `json:"energy_value"`
	Fat                                                   float64 `json:"fat"`
	Fat100g                                               float64 `json:"fat_100g"`
	FatServing                                            float64 `json:"fat_serving"`
	FatUnit                                               string  `json:"fat_unit"`
	FatValue                                              float64 `json:"fat_value"`
	Fiber                                                 float64 `json:"fiber"`
	Fiber100g                                             float64 `json:"fiber_100g"`
	FiberServing                                          float64 `json:"fiber_serving"`
	FiberUnit                                             string  `json:"fiber_unit"`
	FiberValue                                            float64 `json:"fiber_value"`
	FruitsVegetablesLegumesEstimateFromIngredients100g    float64 `json:"fruits-vegetables-legumes-estimate-from-ingredients_100g"`
	FruitsVegetablesLegumesEstimateFromIngredientsServing float64 `json:"fruits-vegetables-legumes-estimate-from-ingredients_serving"`
	FruitsVegetablesNuts                                  float64 `json:"fruits-vegetables-nuts"`
	FruitsVegetablesNutsEstimateFromIngredients100g       float64 `json:"fruits-vegetables-nuts-estimate-from-ingredients_100g"`
	FruitsVegetablesNutsEstimateFromIngredientsServing    float64 `json:"fruits-vegetables-nuts-estimate-from-ingredients_serving"`
	FruitsVegetablesNuts100g                              float64 `json:"fruits-vegetables-nuts_100g"`
	FruitsVegetablesNutsServing                           float64 `json:"fruits-vegetables-nuts_serving"`
	FruitsVegetablesNutsUnit                              string  `json:"fruits-vegetables-nuts_unit"`
	FruitsVegetablesNutsValue                             float64 `json:"fruits-vegetables-nuts_value"`
	NovaGroup                                             int     `json:"nova-group"`
	NovaGroup100g                                         int     `json:"nova-group_100g"`
	NovaGroupServing                                      int     `json:"nova-group_serving"`
	NutritionScoreFR                                      int     `json:"nutrition-score-fr"`
	NutritionScoreFR100g                                  int     `json:"nutrition-score-fr_100g"`
	Proteins                                              float64 `json:"proteins"`
	Proteins100g                                          float64 `json:"proteins_100g"`
	ProteinsServing                                       float64 `json:"proteins_serving"`
	ProteinsUnit                                          string  `json:"proteins_unit"`
	ProteinsValue                                         float64 `json:"proteins_value"`
	Salt                                                  float64 `json:"salt"`
	Salt100g                                              float64 `json:"salt_100g"`
	SaltServing                                           float64 `json:"salt_serving"`
	SaltUnit                                              string  `json:"salt_unit"`
	SaltValue                                             float64 `json:"salt_value"`
	SaturatedFat                                          float64 `json:"saturated-fat"`
	SaturatedFat100g                                      float64 `json:"saturated-fat_100g"`
	SaturatedFatServing                                   float64 `json:"saturated-fat_serving"`
	SaturatedFatUnit                                      string  `json:"saturated-fat_unit"`
	SaturatedFatValue                                     float64 `json:"saturated-fat_value"`
	Sodium                                                float64 `json:"sodium"`
	Sodium100g                                            float64 `json:"sodium_100g"`
	SodiumServing                                         float64 `json:"sodium_serving"`
	SodiumUnit                                            string  `json:"sodium_unit"`
	SodiumValue                                           float64 `json:"sodium_value"`
	Sugars                                                float64 `json:"sugars"`
	Sugars100g                                            float64 `json:"sugars_100g"`
	SugarsServing                                         float64 `json:"sugars_serving"`
	SugarsUnit                                            string  `json:"sugars_unit"`
	SugarsValue                                           float64 `json:"sugars_value"`
}

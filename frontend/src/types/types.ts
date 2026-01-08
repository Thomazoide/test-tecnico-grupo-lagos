export interface User {
    ID: number;
    ExternalUUID: string;
    Cart: Cart;
}

export interface Cart {
    UserID: number;
    Barcodes: string[];
}

export interface ResponsePayload<T> {
    statusCode: number;
    message: string;
    error: boolean;
    data?: T;
}

export interface ProductByCodePayload {
    status: number;
    status_verbose: string;
    code: string;
    product: Product;
}

export interface ProductsPayload {
    count: number;
    page: number;
    page_count: number;
    page_size: number;
    products: Product[];
    skip: number;
}

export interface Product {
    brands: string;
    code: string;
    image_url: string;
    nutriscore_grade: string;
    product_name: string;
    nutriments: Nutrition;
}

export interface Nutrition {
  carbohydrates_100g: number;
  carbohydrates_serving: number;
  carbohydrates_unit: string;
  carbohydrates_value: number;
  'carbon-footprint-from-known-ingredients_100g': number;
  'carbon-footprint-from-known-ingredients_product': number;
  'carbon-footprint-from-known-ingredients_serving': number;
  energy: number;
  'energy-kcal': number;
  'energy-kcal_100g': number;
  'energy-kcal_serving': number;
  'energy-kcal_unit': string;
  'energy-kcal_value': number;
  'energy-kcal_value_computed': number;
  'energy-kj': number;
  'energy-kj_100g': number;
  'energy-kj_serving': number;
  'energy-kj_unit': string;
  'energy-kj_value': number;
  'energy-kj_value_computed': number;
  energy_100g: number;
  energy_serving: number;
  energy_unit: string;
  energy_value: number;
  fat: number;
  fat_100g: number;
  fat_serving: number;
  fat_unit: string;
  fat_value: number;
  fiber: number;
  fiber_100g: number;
  fiber_serving: number;
  fiber_unit: string;
  fiber_value: number;
  'fruits-vegetables-legumes-estimate-from-ingredients_100g': number;
  'fruits-vegetables-legumes-estimate-from-ingredients_serving': number;
  'fruits-vegetables-nuts': number;
  'fruits-vegetables-nuts-estimate-from-ingredients_100g': number;
  'fruits-vegetables-nuts-estimate-from-ingredients_serving': number;
  'fruits-vegetables-nuts_100g': number;
  'fruits-vegetables-nuts_serving': number;
  'fruits-vegetables-nuts_unit': string;
  'fruits-vegetables-nuts_value': number;
  'nova-group': number;
  'nova-group_100g': number;
  'nova-group_serving': number;
  'nutrition-score-fr': number;
  'nutrition-score-fr_100g': number;
  proteins: number;
  proteins_100g: number;
  proteins_serving: number;
  proteins_unit: string;
  proteins_value: number;
  salt: number;
  salt_100g: number;
  salt_serving: number;
  salt_unit: string;
  salt_value: number;
  'saturated-fat': number;
  'saturated-fat_100g': number;
  'saturated-fat_serving': number;
  'saturated-fat_unit': string;
  'saturated-fat_value': number;
  sodium: number;
  sodium_100g: number;
  sodium_serving: number;
  sodium_unit: string;
  sodium_value: number;
  sugars: number;
  sugars_100g: number;
  sugars_serving: number;
  sugars_unit: string;
  sugars_value: number;
}
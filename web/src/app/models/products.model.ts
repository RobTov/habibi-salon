export interface ProductsModel {
  id: number;
  name: string;
  description: string;
  quantity: number;
}

export interface ICreateProductsModel {
  name: string;
  description: string;
  quantity: number;
}

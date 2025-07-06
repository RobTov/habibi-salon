export interface ServicesModel {
  id: number;
  name: string;
  description: string;
  price: number;
  is_active: boolean;
}

export interface ICreateServicesModel {
  name: string;
  description: string;
  price: number;
  is_active: boolean;
}

export interface ClientsModel {
  id: number;
  name: string;
  email: string;
  phone: string;
  address: string;
  created_at: string;
}

export interface ICreateClientsModel {
  name: string;
  email: string;
  phone: string;
  address: string;
  created_at: string;
}

import { computed, Injectable, signal } from '@angular/core';

export type FakeServicesModel = {
  id: number;
  image: string;
  name: string;
  price: string;
};

@Injectable({
  providedIn: 'root',
})
export class FakeServicesService {
  private services = signal<FakeServicesModel[]>([]);

  setInitialServices(): void {
    this.services.set([
      {
        id: 1,
        image: '/img/8.jpeg',
        name: 'Uñas nuevas',
        price: '$1500',
      },
      {
        id: 2,
        image: '/img/8.jpeg',
        name: 'Rellenos',
        price: '$100',
      },
      {
        id: 3,
        image: '/img/8.jpeg',
        name: 'Decoraciones',
        price: 'Relieve $50c/u - Perlas $10c/u',
      },
      {
        id: 4,
        image: '/img/9.jpeg',
        name: 'Keratina',
        price: '$1200 onz',
      },
      {
        id: 5,
        image: '/img/9.jpeg',
        name: 'Alisado',
        price: '$1000 onz',
      },

      {
        id: 6,
        image: '/img/9.jpeg',
        name: 'Tintes',
        price: '$600 onz',
      },

      {
        id: 7,
        image: '/img/11.jpeg',
        name: 'Podología',
        price: '$1500',
      },
      {
        id: 8,
        image: '/img/2.jpeg',
        name: 'Masaje Facial',
        price: '$1000',
      },
    ]);

    window.localStorage.setItem('services', JSON.stringify(this.services()));
  }

  public getServices(): FakeServicesModel[] {
    if (!window.localStorage.getItem('services')) {
      this.setInitialServices()!;
    }
    
    this.services.set(
      JSON.parse(
        window.localStorage.getItem('services')!
      ) as FakeServicesModel[]
    );

    return this.services();
  }
}

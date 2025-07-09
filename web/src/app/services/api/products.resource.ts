import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, firstValueFrom } from 'rxjs';
import { RestService } from '../rest.service';
import {
  ProductsModel,
  ICreateProductsModel,
} from '../../models/products.model';

@Injectable({
  providedIn: 'root',
})
export class ProductsResource {
  private readonly products$: BehaviorSubject<ProductsModel[]> =
    new BehaviorSubject<ProductsModel[]>([]);
  public products = this.products$.asObservable();

  private readonly productDetails$: BehaviorSubject<ProductsModel | null> =
    new BehaviorSubject<ProductsModel | null>(null);
  public productDetails = this.productDetails$.asObservable();

  private readonly restService: RestService = inject(RestService);

  async get() {
    const url = `products/`;
    const products = await firstValueFrom(
      this.restService.get<ProductsModel[]>(url)
    );
    this.products$.next(products);
  }

  async getDetails(productID: string) {
    const service = await firstValueFrom(
      this.restService.get<ProductsModel>(`products/${productID}/`)
    );
    this.productDetails$.next(service);
  }

  async create(productData: ICreateProductsModel): Promise<void> {
    console.log('Creating product', productData);

    const s = (await firstValueFrom(
      this.restService.post('products/', {
        name: productData.name,
        description: productData.description,
        quantity: productData.quantity,
      })
    )) as ProductsModel;
    const products = this.products$.value;
    this.products$.next([...products, s]);
  }

  async edit(request: ICreateProductsModel, productID: number): Promise<void> {
    const response = await firstValueFrom(
      this.restService.put(`products/${productID}/`, {
        name: request.name,
        description: request.description,
        quantity: request.quantity,
      })
    );

    const currentProducts = this.products$.getValue();
    const updatedProduct = currentProducts.map((s) => {
      if (s.id === productID) {
        console.log(request.name);
      }
      return s.id === productID ? { ...s, request } : s;
    });

    console.log(updatedProduct);
    this.products$.next(updatedProduct);
  }

  async delete(request: ProductsModel): Promise<void> {
    const response = await firstValueFrom(
      this.restService.delete(`products/${request.id}/`)
    );

    const currentProducts = this.products$.getValue();
    const updatedProduct = currentProducts.filter((service) => {
      return service.id !== request.id;
    });

    console.log(updatedProduct);
    this.products$.next(updatedProduct);
  }
}

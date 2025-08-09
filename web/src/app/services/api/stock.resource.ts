import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, firstValueFrom } from 'rxjs';
import { RestService } from '../rest.service';
import { ICreateStockModel, StockModel } from '../../models/stock.model';

@Injectable({
  providedIn: 'root',
})
export class StockResource {
  private readonly stocks$: BehaviorSubject<StockModel[]> = new BehaviorSubject<
    StockModel[]
  >([]);
  public stocks = this.stocks$.asObservable();

  private readonly stockDetails$: BehaviorSubject<StockModel | null> =
    new BehaviorSubject<StockModel | null>(null);
  public stockDetails = this.stockDetails$.asObservable();

  private readonly restService: RestService = inject(RestService);

  async get() {
    const url = `stock/`;
    const stocks = await firstValueFrom(
      this.restService.get<StockModel[]>(url)
    );
    this.stocks$.next(stocks);
  }

  async getDetails(stockID: string) {
    const service = await firstValueFrom(
      this.restService.get<StockModel>(`stock/${stockID}/`)
    );
    this.stockDetails$.next(service);
  }

  async create(stockData: ICreateStockModel): Promise<void> {
    console.log('Creating stock', stockData);

    const s = (await firstValueFrom(
      this.restService.post('stock/', {
        product_id: stockData.product_id,
        service_id: stockData.service_id,
        quantity: stockData.quantity,
      })
    )) as StockModel;
    const stocks = this.stocks$.value;
    this.stocks$.next([...stocks, s]);
  }

  async edit(request: ICreateStockModel, stockID: number): Promise<void> {
    const response = await firstValueFrom(
      this.restService.put(`stock/${stockID}/`, {
        product_id: request.product_id,
        service_id: request.service_id,
        quantity: request.quantity,
      })
    );

    const currentstocks = this.stocks$.getValue();
    const updatedStock = currentstocks.map((s) => {
      if (s.id === stockID) {
        console.log(
          `stock of product ${request.product_id} and service ${request.service_id}`
        );
      }
      return s.id === stockID ? { ...s, request } : s;
    });

    console.log(updatedStock);
    this.stocks$.next(updatedStock);
  }

  async delete(request: StockModel): Promise<void> {
    const response = await firstValueFrom(
      this.restService.delete(`stock/${request.id}/`)
    );

    const currentstocks = this.stocks$.getValue();
    const updatedStock = currentstocks.filter((service) => {
      return service.id !== request.id;
    });

    console.log(updatedStock);
    this.stocks$.next(updatedStock);
  }
}

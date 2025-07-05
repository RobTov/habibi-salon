import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class RestService {
  private readonly apiUrl: string = environment.apiURL;
  constructor(private readonly http: HttpClient) {}

  get<T>(path: string): Observable<T> {
    return this.http.get<T>(`${this.apiUrl}/${path}`);
  }

  post<T, C>(path: string, body: C): Observable<T> {
    console.log('path', `${this.apiUrl}/${path}`);
    return this.http.post<T>(`${this.apiUrl}/${path}`, body);
  }

  put<T, C>(path: string, body: C): Observable<T> {
    console.log(`path: ${this.apiUrl}/${path}/`);
    return this.http.patch<T>(`${this.apiUrl}/${path}`, body);
  }

  delete<T>(path: string): Observable<T> {
    return this.http.delete<T>(`${this.apiUrl}/${path}`);
  }
}

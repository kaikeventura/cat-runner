import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, map, throwError } from 'rxjs';
import { HttpRunner, Strategy, StrategyBase } from '../model/strategy.model';

@Injectable({
  providedIn: 'root'
})
export class StrategyService {

  apiHostV1 = "http://localhost:8080/v1"
  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }

  constructor(
    private httpClient: HttpClient
  ) { }

  public getAllStrategies(): Observable<StrategyBase[]> {
    return this.httpClient.get<StrategyBase[]>(`${this.apiHostV1}/strategy`)
  }

  public getStrategyByName(strategyName: string): Observable<Strategy> {
    return this.httpClient.get<Strategy>(`${this.apiHostV1}/strategy/${strategyName}`)
  }

  public createHttpRunner(strategyName: string, httpRunner: HttpRunner): Observable<Strategy> {
    return this.httpClient.post<HttpResponse<any>>(`${this.apiHostV1}/strategy/${strategyName}/http`, httpRunner, this.httpOptions).pipe(
      map((response: HttpResponse<any>) => {
        return response.body as Strategy;
      }),
      catchError((error) => {
        console.error('Error while trying to create an http request: ', error);
        return throwError(error);
      })
    );
  }
}
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Strategy, StrategyBase } from '../model/strategy.model';

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
}

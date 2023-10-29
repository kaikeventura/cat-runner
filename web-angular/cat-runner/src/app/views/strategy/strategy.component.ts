import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { delay } from 'rxjs';
import { Strategy } from 'src/app/shared/model/strategy.model';
import { StrategyService } from 'src/app/shared/service/strategy.service';

@Component({
  selector: 'app-strategy',
  templateUrl: './strategy.component.html',
  styleUrls: ['./strategy.component.css']
})
export class StrategyComponent {

  strategy!: Strategy

  constructor(
    private route: ActivatedRoute, 
    private strategyService: StrategyService
  ) {}
  ngOnInit() {
    this.route.params.subscribe(params => {
      const strategyName = params['strategyName'];
      this.getStrategyByName(strategyName)
    });
  }

  getStrategyByName(strategyName: string) {
    this.strategyService.getStrategyByName(strategyName).subscribe(data => {
      this.strategy = data
    })
  }
}

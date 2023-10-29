import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Strategy } from 'src/app/shared/model/strategy.model';
import { GlobalService } from 'src/app/shared/service/global.service';
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
    private strategyService: StrategyService,
    private globalService: GlobalService
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
      this.globalService.setGlobalVariable(data)
    })
  }
}

import { AfterViewInit, Component, OnInit, ViewChild } from "@angular/core";
import { StrategyComponent } from "../strategy.component";
import { HttpRunner, Strategy } from "src/app/shared/model/strategy.model";
import { GlobalService } from "src/app/shared/service/global.service";

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.css']
})
export class SidenavComponent implements OnInit {
  strategy: Strategy | null = null
  httpRunners: HttpRunner[] = []

  constructor(private globalService: GlobalService) {}

  ngOnInit() {
    this.globalService.getStrategyGlobal().subscribe((strategy) => {
      if (strategy) {
        this.strategy = strategy;
        this.httpRunners = this.strategy.HttpRequestRunners
      }
    })
  }
}
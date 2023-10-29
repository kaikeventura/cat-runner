import { AfterViewInit, Component, ViewChild } from "@angular/core";
import { StrategyComponent } from "../strategy.component";
import { HttpRunner, Strategy } from "src/app/shared/model/strategy.model";
import { GlobalService } from "src/app/shared/service/global.service";

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.css']
})
export class SidenavComponent {
  strategy!: Strategy
  httpRunners!: HttpRunner[]
  
  constructor(private globalService: GlobalService) {
    this.strategy = this.globalService.getGlobalVariable()
    this.httpRunners = this.strategy.HttpRequestRunners
  }
}
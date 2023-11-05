import { Component, QueryList, ViewChildren } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { KeyValueDialogComponent } from './key-value-dialog/key-value-dialog.component';
import { MatTable } from '@angular/material/table';
import { StrategyService } from 'src/app/shared/service/strategy.service';
import { HttpRunner } from 'src/app/shared/model/strategy.model';
import { ActivatedRoute } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-http-form',
  templateUrl: './http-form.component.html',
  styleUrls: ['./http-form.component.css']
})
export class HttpFormComponent {
  constructor(
    private dialog: MatDialog, 
    private strategyService: StrategyService, 
    private route: ActivatedRoute,
    private _snackBar: MatSnackBar
  ) {}

  private currentStrategyName!: string

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.currentStrategyName = params['strategyName'];
    });
  }
  
  tiles: Tile[] = [
    {text: 'Hostname', cols: 5, rows: 1, placeholder: "hostname.api | 192.0.0.1", inputType: "text"},
    {text: 'Port', cols: 1, rows: 1, placeholder: "8090", inputType: "number"},
    {text: 'Method', cols: 1, rows: 1, select: {value: ["GET", "POST", "PUT", "PATCH", "DELETE"]}},
    {text: 'Path', cols: 3, rows: 1, placeholder: "/foo", inputType: "text"},
    {text: 'Timeout', cols: 1, rows: 1, placeholder: "5000", inputType: "number"},
    {text: 'Protocol', cols: 1, rows: 1, select: {value: ["HTTP", "HTTPS"]}}
  ];

  queryParamsToggle = false;
  headerParamsToggle = false;
  requestBodyToggle = false;

  @ViewChildren(MatTable) private tables!: QueryList<MatTable<KeyValue>>;

  private queryParams: KeyValue[] = [];
  private headerParams: KeyValue[] = [];

  queryParamDisplayedColumns = ['key', 'value'];
  queryParamsDataSource = [...this.queryParams];

  headerParamDisplayedColumns = ['key', 'value'];
  headerParamsDataSource = [...this.headerParams];
  
  openQueryParamDialog() {
    const dialogRef = this.dialog.open(KeyValueDialogComponent, {
      data: {
        key: "Add Query Param"
      }
    });

    dialogRef.componentInstance?.addKeyValue.subscribe((values: any) => {
      this.addQueryParam(values.key, values.value);
    });
  }

  openHeaderParamDialog() {
    const dialogRef = this.dialog.open(KeyValueDialogComponent, {
      data: {
        key: "Add Header Param"
      }
    });

    dialogRef.componentInstance?.addKeyValue.subscribe((values: any) => {
      this.addHeaderParam(values.key, values.value);
    });
  }

  private addQueryParam(key: string, value: string) {
    this.queryParamsDataSource.push({key: key, value: value})
    this.renderTables();
  }

  private addHeaderParam(key: string, value: string) {
    this.headerParamsDataSource.push({key: key, value: value})
    this.renderTables();
  }

  private renderTables() {
    this.tables.forEach((table) => {
      table.renderRows();
    });
  }

  formData: any = {};
  
  onSubmitHttpRequest() {
    const httpRunner: HttpRunner = {
      RequestName: "Http Request 90",
      Http: {
        Protocol: this.formData.Protocol,
        Host: this.formData.Hostname,
        Port: this.formData.Port != "" ? parseInt(this.formData.Port) : 0,
        Path: this.formData.Path,
        HttpMethod: this.formData.Method,
        Timeout: this.formData.Timeout != "" ? parseInt(this.formData.Timeout) : 0,
        Headers: this.buildKeyValue(this.headerParamsDataSource),
        Parameters: this.buildKeyValue(this.queryParamsDataSource)
      },
      VirtualUser: {
        UsersAmount: 1,
        InteractionsAmount: 1,
        InteractionDelay: 0
      }
    }

    if (this.formData.requestBody != "") {
      httpRunner.Http.Body = {
        BodyFormat: "JSON",
        ContentText: this.formData.requestBody
      }
    }

    this.strategyService.createHttpRunner(this.currentStrategyName, httpRunner).subscribe(
      success => {
        window.location.reload();
      },
      error => {
        this.openSnackBar(`Error: ${error.error.error}`, "Close")
      }
    );

    this.openSnackBar("Success", "Close")
  }

  private buildKeyValue(source: KeyValue[]) {
    return source.length == 0 ? [] : source.map((keyValue) => ({
      Key: keyValue.key,
      Value: keyValue.value
    }))
  }

  private openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action);
  }
}

export interface KeyValue {
  key: string;
  value: string;
}

export interface Tile {
  cols: number;
  rows: number;
  text: string;
  placeholder?: string;
  inputType?: string,
  select?: Select
}

export interface Select {
  value: string[]
}
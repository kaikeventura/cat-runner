import { Component, QueryList, ViewChildren } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { KeyValueDialogComponent } from './key-value-dialog/key-value-dialog.component';
import { MatTable } from '@angular/material/table';

@Component({
  selector: 'app-http-form',
  templateUrl: './http-form.component.html',
  styleUrls: ['./http-form.component.css']
})
export class HttpFormComponent {
  constructor(private dialog: MatDialog) {}
  
  tiles: Tile[] = [
    {text: 'Hostname', cols: 5, rows: 1, placeholder: "hostname.api | 192.0.0.1", inputType: "text"},
    {text: 'Port', cols: 1, rows: 1, placeholder: "8090", inputType: "number"},
    {text: 'Http Method', cols: 1, rows: 1, select: {value: ["GET", "POST", "PUT", "PATCH", "DELETE"]}},
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

  formData: any = {
  };
  
  onSubmit() {
    console.log('Formulário enviado com sucesso', this.formData);
    console.log('Query Params', this.queryParamsDataSource);
    console.log('Header Params', this.headerParamsDataSource);
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
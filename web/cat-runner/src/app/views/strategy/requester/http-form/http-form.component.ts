import { Component, ViewChild } from '@angular/core';

export interface Tile {
  cols: number;
  rows: number;
  text: string;
  inputType?: string
}

@Component({
  selector: 'app-http-form',
  templateUrl: './http-form.component.html',
  styleUrls: ['./http-form.component.css']
})
export class HttpFormComponent {
  tiles: Tile[] = [
    {text: 'Hostname', cols: 5, rows: 1, inputType: "text"},
    {text: 'Port', cols: 1, rows: 1, inputType: "number"},
    {text: 'Protocol', cols: 1, rows: 1, inputType: "text"},
    {text: 'Http Method', cols: 1, rows: 1, inputType: "text"},
    {text: 'Path', cols: 3, rows: 1, inputType: "text"},
    {text: 'Timeout', cols: 1, rows: 1, inputType: "number"}
  ];

  queryParamsToggle = false;
  headerParamsToggle = false;
  requestBodyToggle = false;

  queryParamDisplayedColumns = ['key', 'value'];
  queryParamsDataSource = QUERY_PARAMS;

  headerParamDisplayedColumns = ['key', 'value'];
  headerParamsDataSource = HEADER_PARAMS;
}

export interface KeyValue {
  key: string;
  value: string;
}

const QUERY_PARAMS: KeyValue[] = [
  {key: 'Hydrogen', value: '1.0079'},
  {key: 'Helium', value: '4.0026'},
  {key: 'Lithium', value: '6.941'},
  {key: 'Beryllium', value: '9.0122'},
  {key: 'Boron', value: '10.811'},
  {key: 'Carbon', value: '12.0107'},
  {key: 'Nitrogen', value: '14.0067'},
  {key: 'Oxygen', value: '15.9994'},
  {key: 'Fluorine', value: '18.9984'},
  {key: 'Neon', value: '20.1797'},
];

const HEADER_PARAMS: KeyValue[] = [
  {key: 'Hydrogen', value: '1.0079'},
  {key: 'Helium', value: '4.0026'},
  {key: 'Lithium', value: '6.941'},
  {key: 'Beryllium', value: '9.0122'},
  {key: 'Boron', value: '10.811'},
  {key: 'Carbon', value: '12.0107'},
  {key: 'Nitrogen', value: '14.0067'},
  {key: 'Oxygen', value: '15.9994'},
  {key: 'Fluorine', value: '18.9984'},
  {key: 'Neon', value: '20.1797'},
];

@Component({
  selector: 'app-input-text',
  template: '<input type="text" placeholder="Enter text">'
})
export class InputTextComponent {}
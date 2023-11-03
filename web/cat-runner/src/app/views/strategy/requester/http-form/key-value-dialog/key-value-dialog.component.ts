import { Component, EventEmitter, Inject, Input, Output } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-key-value-dialog',
  templateUrl: './key-value-dialog.component.html',
  styleUrls: ['./key-value-dialog.component.css']
})
export class KeyValueDialogComponent {
  constructor(@Inject(MAT_DIALOG_DATA) public data: any) {}
  
  @Output() addKeyValue = new EventEmitter<any>();

  key!: string
  value!: string

  addNewKeyValue(key: string, value: string) {
    const params = { key: key, value: value }
    this.addKeyValue.emit(params);
  }
}

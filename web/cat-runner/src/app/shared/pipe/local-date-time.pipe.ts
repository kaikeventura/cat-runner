import { Pipe, PipeTransform } from '@angular/core';
import * as moment from 'moment';

@Pipe({
  name: 'localDateTime'
})
export class LocalDateTimePipe implements PipeTransform {

  transform(date: string): string {
    if (date == "") return ""
    let dateOut: moment.Moment = moment(date, "YYYY-MM-DD HH:mm:ss")
    return dateOut.format("YYYY-MM-DD HH:mm:ss");
  }
}

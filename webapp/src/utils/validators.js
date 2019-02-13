import moment  from 'moment';

export function DateValidator(date) {
  return moment(date, 'L').isValid();
}

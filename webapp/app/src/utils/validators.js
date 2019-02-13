import moment  from 'moment';

export function DateValidator(date) {
  return moment(date, 'L').isValid();
}

export function MaxDateValidator(date) {
  if (!DateValidator(date)) {
    return false;
  }

  date = moment(date, 'L').hours(0).minutes(0).seconds(0);
  const today = moment().hours(23).minutes(59).seconds(59);
  return today.isAfter(date);
}

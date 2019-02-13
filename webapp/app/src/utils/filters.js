import moment from 'moment';

const DEFAULT_FROM_FORMAT = 'YYYY-MM-DD';
const DEFAULT_TO_FORMAT = 'L';

export function DateFilter(from = DEFAULT_FROM_FORMAT, to = DEFAULT_TO_FORMAT) {
  return value => moment(value, from).format(to);
}